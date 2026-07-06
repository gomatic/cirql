# cirql — Design Spec

**Status:** Draft for review **Date:** 2026-06-27 **Scope:** The cirql language as a standalone pure-Go library + tool in the `gomatic` org, and its consumption by `gloo-foo/cmd-json`.

## Goal

Give `cmd-json` a real query language so it rivals jq, by building **cirql** — a composable JSON pipeline language — as a reusable pure-Go module (`gomatic/cirql`) and a standalone CLI/REPL (`cq`). `cmd-json` becomes a thin [gloo](https://github.com/gloo-foo/framework) `Command` adapter over cirql.

The authoritative language description is the existing draft at [`backplane-arcane/www.cirql.io`](https://github.com/backplane-arcane/www.cirql.io) (spec §5 grammar, §5.5 builtins, §9 examples). This document is the _implementation_ design and **overrides the spec where the two conflict** — specifically the parser tooling (see [Grammar](#grammar-antlr-mandatory)).

## Why

- `cmd-json` today is a Go-function toolkit (`Pluck`, `Select`, `from*` converters) with no query _language_. Each operation is a separately constructed Go `Command`; there is no string-driven pipeline a user can type.
- cirql already has a complete language spec, a designated repo, and a website. Implementing it satisfies both goals at once: it is the JSON query language `cmd-json` needs, and a reusable asset for the wider ecosystem.
- `cmd-jq` (the sibling subprocess wrapper around the real `jq` binary) covers users who want literal jq. cirql/`cmd-json` is the pure-Go alternative with a distinct, GraphQL-influenced surface.

## Non-Goals

- The hosted `api.cirql.io` service (spec §6.2) and the optimizer (spec §7) are out of scope for this design; deferred to later sub-projects.
- jq syntax compatibility. cirql is its own language; `cmd-jq` covers jq.

## Architecture

cirql is decomposed into four sub-projects, each with its own spec → plan → build cycle. This document specifies the architecture for all four and the detailed design of **#1 (cirql-core)**, the foundation everything else depends on.

| # | Sub-project | Module | Responsibility | Depends on |
| --- | --- | --- | --- | --- |
| 1 | cirql-core | `gomatic/cirql` | Grammar, AST, value model, expression evaluator, builtins, transform stages, pipeline runtime — **pure, no IO**. | — |
| 2 | cirql-sources | `gomatic/cirql` | Source stages `stdin`/`file`/`http`/`query` (GraphQL); variable propagation + fan-out. Network behind injected interfaces. | 1 |
| 3 | cq tool | `gomatic/cirql` | Standalone CLI/REPL — output formats, readline, history. Built from `gomatic/template.cli`. | 1, 2 |
| 4 | cmd-json adapter | `gloo-foo/cmd-json` | Adapt a cirql pipeline to `gloo.Command[[]byte,[]byte]`; the incoming gloo stream is the `stdin` source. Reconcile/retire the existing `pluck`/`select` Go API. | 1, 2 |

### Data flow (the whole language)

```
query text ──▶ Parse (ANTLR) ──▶ Pipeline AST ──▶ Run(ResultSet) ──▶ ResultSet ──▶ encode
                  │                                    │
            parse errors                         stage executors
          (sentinel + pos)                  (transform: pure; source: IO via DI)
```

A **`ResultSet` is a `[]Object`** (spec §5.3) — the unit flowing between stages. Each `Stage.Execute(ResultSet) (ResultSet, error)`. `Pipeline.Run` threads the result set through the stages left to right, carrying a variable environment for `$var` propagation.

## cirql-core (sub-project #1) — detailed design

### Module layout (gomatic + ANTLR conventions)

Mirrors [`sqlrest/graft`](https://github.com/sqlrest/graft) (the in-ecosystem ANTLR exemplar) and the gomatic shared-Makefile/`go.mod` tool-stanza standard.

```
gomatic/cirql/
  pkg/dialect/cirql/
    cirql.g4            # hand-authored grammar (the WHOLE language — transforms + sources)
    parser.go          # covered wrapper: ErrorListener→sentinel, parse tree→AST
    errors.go          # parse sentinels
    parser_test.go
  src/grammar/cirql/   # GENERATED lexer/parser (committed; excluded from gates)
  value/               # cirql value model + typed accessors
  ast/                 # AST node types
  eval/                # expression evaluator + builtins
  stage/               # transform stage executors
  pipeline/            # ResultSet + sequential runner + variable env
  cirql.go             # public API: Parse(string)→Pipeline; Pipeline.Run(ResultSet)→ResultSet
  docker/antlr/Dockerfile
  scripts/grammars-gen.sh
  Makefile  .goreleaser.yaml  go.mod (tool stanza)  .golangci.yaml
  specs/   (this file + quality-budget.yaml)
```

The **full grammar** is authored in #1 (one `.g4` is the source of truth), but #1 only wires executors for the **transform** stages and the implicit input. Source-stage executors (`query`/`http`/`file`) land in #2; until then, parsing a source stage other than the given input yields `ErrStageUnsupported`.

### Grammar (ANTLR, mandatory)

Per the Go quality standards, every DSL is an **ANTLR4 `.g4` → committed Go** grammar — never hand-rolled and **not** `participle`/`pigeon` (overriding spec §7/§10). `cirql.g4` encodes spec §5.1 (pipeline, stages, expressions with the precedence ladder `or → and → cmp → add → mul → unary → primary`, field access `.a.b[]`, `$var`, literals, `func(...)`). ANTLR handles the operator precedence natively.

- Generation runs **only** via Docker (`make grammars` → `docker/antlr` JRE+jar image → `scripts/grammars-gen.sh`), emitting to `src/grammar/cirql/`. Generated Go is **committed** so plain `go build`/`go test`/CI stay pure-Go and Docker-free. `*.interp`/`*.tokens` are git-ignored.
- `pkg/dialect/cirql/parser.go` is the **fully-covered seam**: it installs a custom `antlr.ErrorListener` that turns a syntax error into the `ErrParse` sentinel carrying `line:col`, then walks the parse tree into the typed `ast` nodes. The generated parser is excluded from coverage; the wrapper is 100%.

### Value model (`value/`)

cirql is a dynamically-typed JSON language; the value space is the union in spec §5.2 (`Null/Bool/Int/Float/String/List/Object`). **Decision:** represent values as the JSON-compatible Go set — `nil`, `bool`, `int64`, `float64`, `string`, `[]Value`, `map[string]Value` — behind a `value.Value` alias, with a **typed accessor layer** providing the named-type discipline the standards require.

_Alternative considered:_ a closed interface hierarchy (`value.Int`, `value.Str`, …). Rejected: it forces boxing/unboxing at every `encoding/json` boundary and across the ANTLR/jq-style surface, adding cognitive load and allocation without buying safety a dynamic language can actually use. The accessor layer gets us testable, named, constant-error type checks without the boxing tax — the same pragmatic choice `gojq` makes, reconciled with the standards via the accessors.

- `value.Value = any` (documented constrained union).
- Accessors return `(typed, error)` with constant sentinels: `AsObject`, `AsList`, `AsString`, `AsInt`, `AsFloat`, `AsBool`, `Truthy`, `Equal`, `Compare`, plus `Kind(Value) Kind` (a named enum). Coercion rules from spec §5.2 (`Int`→`Float` in mixed arithmetic; `null` propagates through field access; `+` concatenates when either operand is `String`) live here as pure functions.
- All operations return new values; nothing mutates its inputs (immutability standard).

### AST (`ast/`)

Typed nodes mirroring spec §6.3, minus the IO-stage internals deferred to #2: `Pipeline{Stages, Vars}`; transform stages `MapStage`, `FlatMapStage`, `FilterStage`, `ReduceStage{Op,Expr}`, `SortStage{Expr,Desc}`, `LimitStage{N}`, `UniqStage{Expr}`; expressions `FieldAccess{Path []PathSegment}`, `VarRef`, `BinaryExpr`, `UnaryExpr`, `FuncCall`, `Literal`. Source-stage nodes (`QueryStage`, `HttpStage`, `FileStage`, `StdinStage`) are declared here so the grammar walker is complete, but their `Execute` lives in #2.

### Expression evaluator (`eval/`)

`Eval(e ast.Expr, env Env) (value.Value, error)` where `Env` carries the current `Object` (for `.field`) and the variable bindings (for `$var`). Each node kind is a small pure function (gocognit ≤ 7); operators dispatch through a table, not a switch ladder. Builtins (spec §5.5: `length keys values type toInt toFloat toString upper lower trim split join contains startsWith now flatten distinct coalesce`) are first-class funcs in a registry `map[FuncName]Builtin`; `now()` takes an **injected clock** so it is deterministic in tests (DI standard). Field access on `null` yields `null`; division by zero yields `null` (spec §5.8) — both asserted by failure-path tests.

### Transform stages (`stage/`)

Each stage is a constructor returning a `pipeline.Stage` whose `Execute` is a pure function of the input `ResultSet` (and the evaluator). Cardinalities per spec §5.3: `map` 1:1, `filter` N:≤N, `reduce` N:1 (ops `count sum min max avg first last group_by collect`), `sort` N:N (stable, `asc`/`desc`), `flatMap` N:M (list-valued mappings expand), `limit` N:≤N, `uniq` N:≤N (full-object equality, or by `expr`). `flatMap`'s list-expansion is the nested-navigation primitive (spec §5.3).

### Pipeline runtime (`pipeline/`)

`ResultSet []value.Object`; `Stage` interface; `Pipeline.Run(in ResultSet) (ResultSet, error)` folds the stages, short-circuiting on the first stage error. Source normalization (spec §5.3: wrap a lone object into a 1-element set; wrap primitive-array elements as `{"value": …}`) is applied to the **input** result set so `cmd-json` (and `stdin`) can hand cirql whatever the stream decoded to.

### Public API (`cirql.go`)

```
Parse(query string) (Pipeline, error)        // ANTLR parse → AST → Pipeline
Pipeline.Run(in ResultSet) (ResultSet, error) // execute transforms
```

This two-call surface is exactly what `cmd-json` (#4) wraps: decode the gloo byte stream → `ResultSet`, `Parse` the query once, `Run` per input, encode each output `Object` back to a line.

### Error handling

One `type Error string` sentinel set per package (constant errors standard); `errors.Is`-matchable; wrap causes with `%w` where a leading constant sentinel adds context. Named sentinels include `ErrParse` (with `line:col`), `ErrType` (type mismatch with stage context), `ErrUnknownFunc`, `ErrStageUnsupported` (source stage before #2), `ErrDivideByZero` (internally → `null` + warning per spec, not surfaced as a hard error). No `fmt.Errorf`/`errors.New` except `%w` wraps.

### Testing & quality

- 100% statement coverage on owned packages; `COVER_PKGS = go list ./... | grep -v '/src/grammar'`; generated parser excluded; the `pkg/dialect/cirql` wrapper is 100%.
- Failure paths asserted specifically (`errors.Is(err, ErrParse)` etc.), not "an error occurred."
- `gocognit -over 7` clean; `gofumpt`, `go vet`, `staticcheck`, `govulncheck` clean; `goreleaser check` valid. Full `make check` green.
- A `specs/quality-budget.yaml` accompanies this spec; planner/implementer/reviewer optimize against it.

## cmd-json adapter (sub-project #4) — design sketch

`cmd-json` gains a primary constructor that takes a cirql query string and returns a `gloo.Command[[]byte,[]byte]`:

```
Json(query string, opts ...any) gloo.Command[[]byte, []byte]
```

It decodes incoming lines into a `cirql.ResultSet` (the `stdin` source), runs the parsed pipeline, and emits each result `Object` as a compact JSON line — preserving cmd-json's existing line-oriented contract. The current `Pluck`/`Select`/`from*` packages are reconciled: `from*` converters stay (they are input adapters, not query language); `pluck`/`select` are re-expressed as cirql queries (`map {…}` / `filter …`) and either retired or kept as thin Go conveniences that build cirql ASTs. Full cirql — including the network source stages from #2 — is available through this one entry point, per the chosen scope.

## Build order

1. **cirql-core** (#1) — this spec. Nothing else compiles without it.
2. **cirql-sources** (#2) — `stdin`/`file`/`http`/`query` executors + propagation.
3. **cmd-json adapter** (#4) — depends only on #1+#2; delivers the headline goal.
4. **cq tool** (#3) — REPL/CLI; can proceed in parallel with #4 once #1+#2 land.
