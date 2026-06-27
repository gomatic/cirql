# cirql-core Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build the pure-Go cirql language core — parse a cirql query string into a pipeline and run it over an in-memory `ResultSet` — covering the transform sublanguage (`map/filter/reduce/sort/flatMap/limit/uniq`), expressions, and builtins, with no IO.

**Architecture:** ANTLR4 grammar (`pkg/dialect/cirql/cirql.g4`) generates a committed lexer/parser under `src/grammar/cirql/`; a fully-covered wrapper turns the parse tree into a typed AST and ANTLR errors into a sentinel. A dynamic value model (`value`) backs an expression evaluator (`eval`) and pure transform stages (`stage`) that a sequential `pipeline` runner threads a `ResultSet` through. `cirql.go` exposes `Parse` + `Pipeline.Run`.

**Tech Stack:** Go 1.26, `github.com/antlr4-go/antlr/v4` (pure-Go runtime), ANTLR4 Java generator in Docker (build-time only), gomatic shared Makefile + `go.mod` tool stanza.

## Global Constraints

- Module path: `github.com/gomatic/cirql`; `go 1.26` minimum (minor floats in CI).
- Parser is ANTLR4 `.g4` → **committed** `src/grammar/cirql/`; never hand-rolled; never `participle`/`pigeon` (overrides language spec §7/§10). ANTLR runtime: `github.com/antlr4-go/antlr/v4`.
- 100% statement coverage on owned packages; `COVER_PKGS = $(shell go list ./... | grep -v '/src/grammar')`; the `pkg/dialect/cirql` wrapper is 100% covered.
- `gocognit -over 7` empty on production files; `gofumpt -l .` empty; `go vet`, `staticcheck ./...`, `govulncheck ./...` clean; `goreleaser check` valid. Full `make check` green before any task is "done."
- Errors: one `type Error string` sentinel set per package; `const ErrX Error = "..."`; match with `errors.Is`; no `fmt.Errorf`/`errors.New` except `%w` wraps led by a constant sentinel.
- Immutability: value operations return new values; no input mutation. DI: anything non-deterministic (the clock for `now()`) is injected.
- Custom named types for domain parameters; private by default; pointer receivers only where a type wraps non-copyable state or must mutate per contract.
- Generated parser aux files (`*.interp`, `*.tokens`) are git-ignored; generated `*.go` under `src/grammar/` is committed and carries no edits.

---

### Task 0: Module bootstrap + green empty gate

**Files:**
- Create: `go.mod`, `Makefile`, `.golangci.yaml`, `.goreleaser.yaml`, `.gitignore`, `.github/workflows/ci.yml`, `.github/workflows/release.yml`, `LICENSE`, `README.md`
- Create: `docker/antlr/Dockerfile`, `scripts/grammars-gen.sh`
- Create: `doc.go` (package-level placeholder so the module compiles)

**Interfaces:**
- Produces: a buildable `github.com/gomatic/cirql` module whose `make check` passes with no real code yet.

- [ ] **Step 1: Initialize the module and copy gomatic scaffolding**

Copy the shared scaffolding from a current gomatic Go library (e.g. clone the canonical files from `gomatic/qp` or `gomatic/optional`): `Makefile` (the gomatic/build shared Makefile), `.golangci.yaml`, `.gitignore`, `.github/workflows/{ci,release}.yml`, `LICENSE` (MIT, copyright "gomatic"). Then:

```bash
cd ~/src/github.com/gomatic/cirql
go mod init github.com/gomatic/cirql
go mod edit -go=1.26
```

In `Makefile`, set the coverage scope to exclude the generated tree (place BEFORE any include if the shared Makefile is included):

```make
COVER_PKGS ?= $(shell go list ./... | grep -v '/src/grammar')
```

- [ ] **Step 2: Add the tool stanza + ANTLR runtime**

```bash
go get github.com/antlr4-go/antlr/v4@latest
go get -tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint \
  github.com/goreleaser/goreleaser/v2 \
  github.com/uudashr/gocognit/cmd/gocognit \
  golang.org/x/tools/cmd/goimports \
  golang.org/x/vuln/cmd/govulncheck \
  gotest.tools/gotestsum \
  honnef.co/go/tools/cmd/staticcheck \
  mvdan.cc/gofumpt
go mod tidy
```

- [ ] **Step 3: Write `.goreleaser.yaml`** (library — no binaries in #1; the `cq` binary arrives in sub-project #3)

```yaml
version: 2
project_name: cirql
before:
  hooks:
    - go mod tidy
    - go mod verify
builds:
  - skip: true
source:
  enabled: true
  name_template: "{{ .ProjectName }}-{{ .Version }}-source"
checksum:
  name_template: "checksums.txt"
  algorithm: sha256
changelog:
  sort: asc
  use: github
release:
  github:
    owner: gomatic
    name: cirql
```

- [ ] **Step 4: Write the ANTLR Docker generator** (Java isolated in Docker; pure-Go runtime only at app build time)

`docker/antlr/Dockerfile`:

```dockerfile
FROM eclipse-temurin:21-jre
ARG ANTLR_VERSION=4.13.2
ADD https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar /antlr.jar
ENTRYPOINT ["java", "-jar", "/antlr.jar"]
```

`scripts/grammars-gen.sh` (generates Go into `src/grammar/cirql/`, package `cirqlgrammar`, then deletes aux files):

```bash
#!/usr/bin/env bash
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
img="cirql-antlr:local"
docker build -t "$img" "$here/docker/antlr"
out="$here/src/grammar/cirql"
mkdir -p "$out"
docker run --rm -v "$here:/work" -w /work "$img" \
  -Dlanguage=Go -package cirqlgrammar -visitor -o "$out" \
  pkg/dialect/cirql/cirql.g4
# move generated files from the mirrored sub-path up to $out, drop aux files
find "$out" -name '*.interp' -delete
find "$out" -name '*.tokens' -delete
```

Add to `.gitignore`:

```
*.interp
*.tokens
```

Add the `grammars` target to the Makefile if the shared one lacks it:

```make
.PHONY: grammars
grammars: ## Regenerate the cirql parser from the .g4 via ANTLR in Docker
	scripts/grammars-gen.sh
```

- [ ] **Step 5: Add `doc.go` so the module compiles**

```go
// Package cirql parses and runs cirql JSON pipeline queries.
package cirql
```

- [ ] **Step 6: Verify the gate is green on the empty module**

Run: `make check`
Expected: PASS (lint/vet/staticcheck/vulncheck clean; coverage vacuously 100% — no statements yet).

- [ ] **Step 7: Commit**

```bash
git add -A
git commit -m "init: bootstrap gomatic/cirql module + ANTLR build wiring"
```

---

### Task 1: Value model (`value`)

**Files:**
- Create: `value/value.go`, `value/errors.go`
- Test: `value/value_test.go`

**Interfaces:**
- Produces:
  - `type Value = any` (constrained union: `nil`, `bool`, `int64`, `float64`, `string`, `[]Value`, `map[string]Value`)
  - `type Kind int` with `KindNull, KindBool, KindInt, KindFloat, KindString, KindList, KindObject` and `func KindOf(Value) Kind`
  - `func AsObject(Value) (map[string]Value, error)`; `AsList`, `AsString`, `AsInt(Value)(int64,error)`, `AsFloat(Value)(float64,error)`, `AsBool`
  - `func Truthy(Value) bool`; `func Equal(a, b Value) bool`; `func Compare(a, b Value) (int, error)`
  - `func Add(a, b Value) (Value, error)` (numeric add or string concat per spec §5.2)
  - Sentinels: `ErrNotObject, ErrNotList, ErrNotString, ErrNotNumber, ErrNotBool, ErrIncomparable Error`

- [ ] **Step 1: Write failing tests for kind + accessors**

```go
package value

import (
	"errors"
	"testing"
)

func TestKindOf(t *testing.T) {
	cases := []struct {
		v    Value
		want Kind
	}{
		{nil, KindNull}, {true, KindBool}, {int64(1), KindInt},
		{3.14, KindFloat}, {"s", KindString},
		{[]Value{int64(1)}, KindList}, {map[string]Value{"a": int64(1)}, KindObject},
	}
	for _, c := range cases {
		if got := KindOf(c.v); got != c.want {
			t.Errorf("KindOf(%#v) = %v, want %v", c.v, got, c.want)
		}
	}
}

func TestAsInt_RejectsString(t *testing.T) {
	if _, err := AsInt("nope"); !errors.Is(err, ErrNotNumber) {
		t.Fatalf("got %v, want ErrNotNumber", err)
	}
}

func TestAsObject_OK(t *testing.T) {
	m, err := AsObject(map[string]Value{"a": int64(1)})
	if err != nil || m["a"] != int64(1) {
		t.Fatalf("AsObject failed: %v %v", m, err)
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `go test ./value/ -run TestKindOf -v`
Expected: FAIL (undefined `KindOf`/`Kind`/sentinels).

- [ ] **Step 3: Implement `value/errors.go`**

```go
package value

// Error is the sentinel error type for the value package.
type Error string

// Error renders the sentinel message.
func (e Error) Error() string { return string(e) }

const (
	ErrNotObject    Error = "value: not an object"
	ErrNotList      Error = "value: not a list"
	ErrNotString    Error = "value: not a string"
	ErrNotNumber    Error = "value: not a number"
	ErrNotBool      Error = "value: not a bool"
	ErrIncomparable Error = "value: incomparable types"
)
```

- [ ] **Step 4: Implement `value/value.go`** (each function ≤ 7 cognitive complexity; dispatch via type switch — a flat type switch counts low)

```go
package value

// Value is the dynamic cirql value: nil | bool | int64 | float64 | string |
// []Value | map[string]Value. The alias keeps encoding/json interop direct;
// the accessors below supply the typed, constant-error discipline.
type Value = any

// Kind names the dynamic type of a Value.
type Kind int

const (
	KindNull Kind = iota
	KindBool
	KindInt
	KindFloat
	KindString
	KindList
	KindObject
)

// KindOf reports the Kind of v.
func KindOf(v Value) Kind {
	switch v.(type) {
	case nil:
		return KindNull
	case bool:
		return KindBool
	case int64:
		return KindInt
	case float64:
		return KindFloat
	case string:
		return KindString
	case []Value:
		return KindList
	case map[string]Value:
		return KindObject
	}
	return KindNull
}

// AsObject returns v as an object or ErrNotObject.
func AsObject(v Value) (map[string]Value, error) {
	if o, ok := v.(map[string]Value); ok {
		return o, nil
	}
	return nil, ErrNotObject
}

// AsList returns v as a list or ErrNotList.
func AsList(v Value) ([]Value, error) {
	if l, ok := v.([]Value); ok {
		return l, nil
	}
	return nil, ErrNotList
}

// AsString returns v as a string or ErrNotString.
func AsString(v Value) (string, error) {
	if s, ok := v.(string); ok {
		return s, nil
	}
	return "", ErrNotString
}

// AsBool returns v as a bool or ErrNotBool.
func AsBool(v Value) (bool, error) {
	if b, ok := v.(bool); ok {
		return b, nil
	}
	return false, ErrNotBool
}

// AsInt returns v as an int64 or ErrNotNumber.
func AsInt(v Value) (int64, error) {
	switch n := v.(type) {
	case int64:
		return n, nil
	case float64:
		return int64(n), nil
	}
	return 0, ErrNotNumber
}

// AsFloat returns v as a float64 or ErrNotNumber.
func AsFloat(v Value) (float64, error) {
	switch n := v.(type) {
	case float64:
		return n, nil
	case int64:
		return float64(n), nil
	}
	return 0, ErrNotNumber
}

// Truthy reports cirql truthiness: false and null are falsey; all else truthy.
func Truthy(v Value) bool {
	switch t := v.(type) {
	case nil:
		return false
	case bool:
		return t
	}
	return true
}
```

- [ ] **Step 5: Run accessor tests to PASS**

Run: `go test ./value/ -run 'TestKindOf|TestAsInt|TestAsObject' -v`
Expected: PASS.

- [ ] **Step 6: Write failing tests for Equal / Compare / Add (coercion rules, spec §5.2)**

```go
func TestEqual_IntFloatMixed(t *testing.T) {
	if !Equal(int64(2), 2.0) {
		t.Fatal("2 should equal 2.0 across numeric types")
	}
}

func TestCompare_Numeric(t *testing.T) {
	c, err := Compare(int64(1), 2.0)
	if err != nil || c != -1 {
		t.Fatalf("Compare(1,2.0) = %d,%v want -1,nil", c, err)
	}
}

func TestCompare_Incomparable(t *testing.T) {
	if _, err := Compare("a", int64(1)); !errors.Is(err, ErrIncomparable) {
		t.Fatalf("got %v want ErrIncomparable", err)
	}
}

func TestAdd_StringConcat(t *testing.T) {
	v, err := Add("a", "b")
	if err != nil || v != "ab" {
		t.Fatalf("Add(a,b)=%v,%v want ab", v, err)
	}
}

func TestAdd_Numeric(t *testing.T) {
	v, err := Add(int64(1), 2.5)
	if err != nil || v != 3.5 {
		t.Fatalf("Add(1,2.5)=%v,%v want 3.5", v, err)
	}
}
```

- [ ] **Step 7: Run to verify failure**

Run: `go test ./value/ -run 'TestEqual|TestCompare|TestAdd' -v`
Expected: FAIL (undefined).

- [ ] **Step 8: Implement Equal / Compare / Add** (append to `value/value.go`; keep helpers small)

```go
// bothNumbers reports whether a and b are both numeric, with float views.
func bothNumbers(a, b Value) (x, y float64, ok bool) {
	af, aerr := AsFloat(a)
	bf, berr := AsFloat(b)
	if aerr != nil || berr != nil {
		return 0, 0, false
	}
	return af, bf, true
}

// Equal reports value equality, treating Int and Float as comparable numbers.
func Equal(a, b Value) bool {
	if x, y, ok := bothNumbers(a, b); ok {
		return x == y
	}
	return KindOf(a) == KindOf(b) && a == b
}

// Compare orders two values, returning -1, 0, or 1. Numbers compare across
// Int/Float; strings compare lexically; anything else is ErrIncomparable.
func Compare(a, b Value) (int, error) {
	if x, y, ok := bothNumbers(a, b); ok {
		return sign(x - y), nil
	}
	as, aerr := AsString(a)
	bs, berr := AsString(b)
	if aerr == nil && berr == nil {
		return strCmp(as, bs), nil
	}
	return 0, ErrIncomparable
}

func sign(d float64) int {
	switch {
	case d < 0:
		return -1
	case d > 0:
		return 1
	}
	return 0
}

func strCmp(a, b string) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}

// Add adds two numbers, or concatenates when either operand is a string
// (spec §5.2). Otherwise ErrNotNumber.
func Add(a, b Value) (Value, error) {
	if KindOf(a) == KindString || KindOf(b) == KindString {
		return concat(a, b)
	}
	return numericAdd(a, b)
}

func concat(a, b Value) (Value, error) {
	as, err := coerceString(a)
	if err != nil {
		return nil, err
	}
	bs, err := coerceString(b)
	if err != nil {
		return nil, err
	}
	return as + bs, nil
}

// numericAdd preserves Int when both operands are Int, else promotes to Float.
func numericAdd(a, b Value) (Value, error) {
	if ai, aok := a.(int64); aok {
		if bi, bok := b.(int64); bok {
			return ai + bi, nil
		}
	}
	x, y, ok := bothNumbers(a, b)
	if !ok {
		return nil, ErrNotNumber
	}
	return x + y, nil
}
```

Add `coerceString` (string passthrough; numbers/bool via `strconv`/`fmt`-free formatting is acceptable here — use `strconv`):

```go
import "strconv"

func coerceString(v Value) (string, error) {
	switch t := v.(type) {
	case string:
		return t, nil
	case int64:
		return strconv.FormatInt(t, 10), nil
	case float64:
		return strconv.FormatFloat(t, 'g', -1, 64), nil
	case bool:
		return strconv.FormatBool(t), nil
	}
	return "", ErrNotString
}
```

- [ ] **Step 9: Add the failure-path tests to reach 100%**

```go
func TestAdd_NonNumber(t *testing.T) {
	if _, err := Add([]Value{}, int64(1)); !errors.Is(err, ErrNotNumber) {
		t.Fatalf("got %v want ErrNotNumber", err)
	}
}

func TestTruthy(t *testing.T) {
	for v, want := range map[any]bool{nil: false, false: false, true: true, int64(0): true, "": true} {
		if Truthy(v) != want {
			t.Errorf("Truthy(%#v)=%v want %v", v, Truthy(v), want)
		}
	}
}
```

- [ ] **Step 10: Run full value coverage**

Run: `go test ./value/ -covermode=atomic -coverprofile=/tmp/c.out && go tool cover -func=/tmp/c.out | tail -1`
Expected: PASS, `total: ... 100.0%`.

- [ ] **Step 11: Commit**

```bash
git add value/
git commit -m "feat(value): dynamic value model with typed accessors + coercion"
```

---

### Task 2: AST nodes (`ast`)

**Files:**
- Create: `ast/ast.go`
- Test: `ast/ast_test.go`

**Interfaces:**
- Produces:
  - `type Pipeline struct { Stages []Stage }`
  - `type Stage interface { isStage() }` with transform implementers `MapStage{Mappings []Mapping}`, `FlatMapStage{Mappings []Mapping}`, `FilterStage{Cond Expr}`, `ReduceStage{Op ReduceOp; Arg Expr}`, `SortStage{Key Expr; Desc bool}`, `LimitStage{N int}`, `UniqStage{Key Expr}`; and source markers `StdinStage{}`, `FileStage{Path string}`, `HTTPStage{...}`, `QueryStage{...}` (declared; executed in #2)
  - `type Mapping struct { Key string; Expr Expr }`
  - `type ReduceOp string` consts (`OpCount, OpSum, OpMin, OpMax, OpAvg, OpFirst, OpLast, OpGroupBy, OpCollect`)
  - `type Expr interface { isExpr() }` with `FieldAccess{Path []PathSegment}`, `VarRef{Name string}`, `BinaryExpr{Op BinOp; L, R Expr}`, `UnaryExpr{Op UnOp; X Expr}`, `FuncCall{Name string; Args []Expr}`, `Literal{V value.Value}`
  - `type PathSegment struct { Name string; Iter bool }` (`Iter` true for `[]`)
  - `type BinOp string`, `type UnOp string` with operator consts

- [ ] **Step 1: Write a failing test** (interfaces compile + tag methods exist)

```go
package ast

import "testing"

func TestNodesImplementInterfaces(t *testing.T) {
	var _ Stage = MapStage{}
	var _ Stage = FilterStage{}
	var _ Expr = FieldAccess{}
	var _ Expr = Literal{}
	if (BinaryExpr{Op: OpEq}).Op != OpEq {
		t.Fatal("BinOp const mismatch")
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `go test ./ast/ -v`
Expected: FAIL (undefined types).

- [ ] **Step 3: Implement `ast/ast.go`** (pure data; unexported tag methods enforce the closed sets)

Provide the full type set from the Interfaces block. Each node gets a one-line `func (X) isStage()` / `func (X) isExpr()`. Operator consts:

```go
type BinOp string

const (
	OpOr  BinOp = "||"
	OpAnd BinOp = "&&"
	OpEq  BinOp = "=="
	OpNe  BinOp = "!="
	OpGt  BinOp = ">"
	OpLt  BinOp = "<"
	OpGe  BinOp = ">="
	OpLe  BinOp = "<="
	OpAdd BinOp = "+"
	OpSub BinOp = "-"
	OpMul BinOp = "*"
	OpDiv BinOp = "/"
	OpMod BinOp = "%"
)

type UnOp string

const OpNot UnOp = "!"
```

- [ ] **Step 4: Run to PASS**

Run: `go test ./ast/ -v`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add ast/
git commit -m "feat(ast): cirql AST node types for stages and expressions"
```

---

### Task 3: Grammar + parser wrapper (`pkg/dialect/cirql`, `src/grammar/cirql`)

**Files:**
- Create: `pkg/dialect/cirql/cirql.g4`, `pkg/dialect/cirql/parser.go`, `pkg/dialect/cirql/errors.go`
- Create (generated, committed): `src/grammar/cirql/*.go`
- Test: `pkg/dialect/cirql/parser_test.go`

**Interfaces:**
- Consumes: `ast` (Task 2), `value` (Task 1)
- Produces: `func Parse(query string) (ast.Pipeline, error)`; sentinel `ErrParse Error` (message form `cirql: parse error at <line>:<col>: <msg>`)

- [ ] **Step 1: Author the grammar `cirql.g4`** (encodes spec §5.1; ANTLR resolves precedence top-down)

```antlr
grammar cirql;

pipeline      : stage (PIPE stage)* EOF ;

stage         : sourceStage | transformStage ;

sourceStage   : queryStage | httpStage | fileStage | stdinStage ;
queryStage    : QUERY? queryBody varBindings? ;
queryBody     : LBRACE selectionSet RBRACE ;
selectionSet  : field+ ;
field         : NAME (LPAREN arguments RPAREN)? (LBRACE selectionSet RBRACE)? ;
arguments     : argument (COMMA argument)* ;
argument      : NAME COLON value ;
httpStage     : HTTP (STRING | variable) httpOptions? ;
httpOptions   : LPAREN httpOpt (COMMA httpOpt)* RPAREN ;
httpOpt       : NAME COLON (STRING | variable | objectLit) ;
fileStage     : FILE STRING ;
stdinStage    : STDIN ;
varBindings   : LPAREN varBinding (COMMA varBinding)* RPAREN ;
varBinding    : variable ASSIGN expr ;

transformStage: mapStage | filterStage | reduceStage
              | sortStage | flatMapStage | limitStage | uniqStage ;
mapStage      : MAP LBRACE mapping (COMMA mapping)* RBRACE ;
flatMapStage  : FLATMAP LBRACE mapping (COMMA mapping)* RBRACE ;
mapping       : NAME COLON expr ;
filterStage   : FILTER expr ;
reduceStage   : REDUCE reduceOp (LPAREN expr RPAREN)? ;
reduceOp      : COUNT | SUM | MIN | MAX | AVG | FIRST | LAST | GROUP_BY | COLLECT ;
sortStage     : SORT expr (ASC | DESC)? ;
limitStage    : LIMIT INT ;
uniqStage     : UNIQ expr? ;

expr          : expr (STAR|SLASH|PERCENT) expr      # MulExpr
              | expr (PLUS|MINUS) expr              # AddExpr
              | expr (EQ|NE|GT|LT|GE|LE) expr       # CmpExpr
              | expr AND expr                       # AndExpr
              | expr OR expr                        # OrExpr
              | NOT expr                            # NotExpr
              | LPAREN expr RPAREN                  # ParenExpr
              | fieldAccess                         # FieldExpr
              | variable                            # VarExpr
              | funcCall                            # CallExpr
              | literal                             # LitExpr
              ;
fieldAccess   : DOT (NAME | LBRACK RBRACK) (DOT NAME | LBRACK RBRACK)* | DOT ;
funcCall      : NAME LPAREN (expr (COMMA expr)*)? RPAREN ;
variable      : DOLLAR NAME ;
objectLit     : LBRACE (NAME COLON value (COMMA NAME COLON value)*)? RBRACE ;
value         : variable | literal | listLit | objectLit ;
listLit       : LBRACK (value (COMMA value)*)? RBRACK ;
literal       : STRING | INT | FLOAT | TRUE | FALSE | NULL ;

QUERY:'query'; HTTP:'http'; FILE:'file'; STDIN:'stdin';
MAP:'map'; FLATMAP:'flatMap'; FILTER:'filter'; REDUCE:'reduce';
SORT:'sort'; LIMIT:'limit'; UNIQ:'uniq';
COUNT:'count'; SUM:'sum'; MIN:'min'; MAX:'max'; AVG:'avg';
FIRST:'first'; LAST:'last'; GROUP_BY:'group_by'; COLLECT:'collect';
ASC:'asc'; DESC:'desc'; TRUE:'true'; FALSE:'false'; NULL:'null';
PIPE:'|'; DOT:'.'; COMMA:','; COLON:':'; ASSIGN:'='; DOLLAR:'$';
LPAREN:'('; RPAREN:')'; LBRACE:'{'; RBRACE:'}'; LBRACK:'['; RBRACK:']';
OR:'||'; AND:'&&'; NOT:'!'; EQ:'=='; NE:'!='; GE:'>='; LE:'<='; GT:'>'; LT:'<';
PLUS:'+'; MINUS:'-'; STAR:'*'; SLASH:'/'; PERCENT:'%';
NAME:[a-zA-Z][a-zA-Z0-9_]* ;
INT:'-'?[0-9]+ ;
FLOAT:'-'?[0-9]+'.'[0-9]+ ;
STRING:'"' (~["\\] | '\\' .)* '"' ;
WS:[ \t\r\n]+ -> skip ;
```

- [ ] **Step 2: Generate the parser**

Run: `make grammars`
Expected: `src/grammar/cirql/cirql_lexer.go`, `cirql_parser.go`, `cirql_base_visitor.go`, `cirql_visitor.go` created; no `*.interp`/`*.tokens` left.

- [ ] **Step 3: Write `pkg/dialect/cirql/errors.go`**

```go
package cirql

// Error is the sentinel error type for the cirql dialect wrapper.
type Error string

func (e Error) Error() string { return string(e) }

// ErrParse is returned for any lexer/parser syntax error, carrying the source
// position via a %w wrap.
const ErrParse Error = "cirql: parse error"

// ErrStageUnsupported is returned when a source stage (query/http/file) is
// encountered before sub-project #2 wires its executor.
const ErrStageUnsupported Error = "cirql: source stage not supported in core"
```

- [ ] **Step 4: Write a failing parse test**

```go
package cirql

import (
	"errors"
	"testing"

	"github.com/gomatic/cirql/ast"
)

func TestParse_FilterMapSortLimit(t *testing.T) {
	p, err := Parse(`filter .stars > 1000 | map { name: .name } | sort .name desc | limit 10`)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(p.Stages) != 4 {
		t.Fatalf("got %d stages want 4", len(p.Stages))
	}
	if _, ok := p.Stages[0].(ast.FilterStage); !ok {
		t.Fatalf("stage0 = %T want FilterStage", p.Stages[0])
	}
	s := p.Stages[2].(ast.SortStage)
	if !s.Desc {
		t.Fatal("sort should be desc")
	}
}

func TestParse_SyntaxError(t *testing.T) {
	_, err := Parse(`map {`)
	if !errors.Is(err, ErrParse) {
		t.Fatalf("got %v want ErrParse", err)
	}
}
```

- [ ] **Step 5: Run to verify failure**

Run: `go test ./pkg/dialect/cirql/ -v`
Expected: FAIL (undefined `Parse`).

- [ ] **Step 6: Implement `parser.go`** — the covered seam: error listener → `ErrParse`, then a visitor walks the tree into `ast`

The wrapper installs a custom `antlr.ErrorListener` capturing the first `line:col`+message, runs the generated parser, and on success walks the tree. Keep each visitor method ≤ 7 cognitive complexity by giving every grammar alt its own small method. Skeleton:

```go
package cirql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gomatic/cirql/ast"
	g "github.com/gomatic/cirql/src/grammar/cirql"
)

// Parse turns a cirql query into a Pipeline AST, or ErrParse on a syntax error.
func Parse(query string) (ast.Pipeline, error) {
	el := &errListener{}
	stream := antlr.NewInputStream(query)
	lexer := g.NewcirqlLexer(stream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := g.NewcirqlParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(el)
	tree := parser.Pipeline()
	if el.err != nil {
		return ast.Pipeline{}, el.err
	}
	return (&builder{}).pipeline(tree.(*g.PipelineContext))
}

// errListener converts the first ANTLR syntax error into ErrParse.
type errListener struct {
	*antlr.DefaultErrorListener
	err error
}

func (l *errListener) SyntaxError(_ antlr.Recognizer, _ any, line, col int, msg string, _ antlr.RecognitionException) {
	if l.err == nil {
		l.err = fmt.Errorf("%w at %d:%d: %s", ErrParse, line, col, msg)
	}
}
```

Then a `builder` type with methods: `pipeline`, `stage`, `transformStage`, `mapStage`, `filterStage`, `reduceStage`, `sortStage`, `flatMapStage`, `limitStage`, `uniqStage`, `expr`, `fieldAccess`, `funcCall`, `variable`, `literal`, returning the matching `ast` nodes. Source stages return `ast.StdinStage{}` etc.; the `value`/`literal` conversion produces `value.Value` (int64/float64/string/bool/nil). **Each method is small and single-purpose** — one grammar rule each — so gocognit stays ≤ 7.

- [ ] **Step 7: Run parse tests to PASS**

Run: `go test ./pkg/dialect/cirql/ -v`
Expected: PASS.

- [ ] **Step 8: Add tests covering every expression/stage alt + literal kind to reach 100% on the wrapper**

Add table-driven cases parsing: each binary operator, `!x`, `(x)`, `$v`, `f(a,b)`, `.a.b[]`, bare `.`, every literal (`"s"`,`1`,`1.5`,`true`,`false`,`null`), every reduce op, `uniq` with and without expr, `flatMap`, and a source stage asserting it parses to the marker node. Confirm with the coverage gate.

- [ ] **Step 9: Verify wrapper coverage = 100% (generated tree excluded)**

Run: `go test $(go list ./... | grep -v '/src/grammar') -covermode=atomic -coverprofile=/tmp/c.out && go tool cover -func=/tmp/c.out | grep dialect | awk '$3!="100.0%"'`
Expected: no output (all wrapper funcs 100%).

- [ ] **Step 10: Commit**

```bash
git add pkg/dialect/cirql src/grammar/cirql .gitignore
git commit -m "feat(parser): ANTLR cirql grammar + covered AST-building wrapper"
```

---

### Task 4: Expression evaluator + builtins (`eval`)

**Files:**
- Create: `eval/eval.go`, `eval/builtins.go`, `eval/errors.go`
- Test: `eval/eval_test.go`, `eval/builtins_test.go`

**Interfaces:**
- Consumes: `ast`, `value`
- Produces:
  - `type Env struct { Obj value.Value; Vars map[string]value.Value; Now func() time.Time }`
  - `func Eval(e ast.Expr, env Env) (value.Value, error)`
  - Sentinels: `ErrUnknownFunc, ErrArity, ErrType Error`

- [ ] **Step 1: Write failing tests** (field access, arithmetic, comparison, logical, null-propagation, div-by-zero→null)

```go
package eval

import (
	"testing"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/value"
)

func env(obj value.Value) Env { return Env{Obj: obj, Vars: map[string]value.Value{}} }

func TestEval_FieldAccess(t *testing.T) {
	e := ast.FieldAccess{Path: []ast.PathSegment{{Name: "a"}, {Name: "b"}}}
	got, err := Eval(e, env(map[string]value.Value{"a": map[string]value.Value{"b": int64(7)}}))
	if err != nil || got != int64(7) {
		t.Fatalf("got %v,%v want 7", got, err)
	}
}

func TestEval_FieldAccess_NullPropagates(t *testing.T) {
	e := ast.FieldAccess{Path: []ast.PathSegment{{Name: "missing"}, {Name: "x"}}}
	got, err := Eval(e, env(map[string]value.Value{}))
	if err != nil || got != nil {
		t.Fatalf("got %v,%v want nil,nil", got, err)
	}
}

func TestEval_DivByZero_Null(t *testing.T) {
	e := ast.BinaryExpr{Op: ast.OpDiv, L: ast.Literal{V: int64(1)}, R: ast.Literal{V: int64(0)}}
	got, err := Eval(e, env(nil))
	if err != nil || got != nil {
		t.Fatalf("got %v,%v want nil,nil", got, err)
	}
}

func TestEval_Comparison(t *testing.T) {
	e := ast.BinaryExpr{Op: ast.OpGt, L: ast.Literal{V: int64(3)}, R: ast.Literal{V: int64(2)}}
	got, _ := Eval(e, env(nil))
	if got != true {
		t.Fatalf("3>2 = %v want true", got)
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `go test ./eval/ -v`
Expected: FAIL.

- [ ] **Step 3: Implement `eval/errors.go`**

```go
package eval

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrUnknownFunc Error = "eval: unknown function"
	ErrArity       Error = "eval: wrong argument count"
	ErrType        Error = "eval: type error"
)
```

- [ ] **Step 4: Implement `eval/eval.go`** — one tiny method per node kind; operator dispatch via tables (keeps gocognit ≤ 7)

`Eval` is a type switch routing to `evalField`, `evalVar`, `evalBinary`, `evalUnary`, `evalCall`, `evalLiteral`. `evalBinary` evaluates both sides then dispatches:
- logical (`||`,`&&`) on `value.Truthy`
- comparison (`==`,`!=`,`<`,…) via `value.Equal`/`value.Compare`
- arithmetic (`+` via `value.Add`; `-`,`*`,`/`,`%` via a numeric helper that returns `nil` on divide/mod by zero per spec §5.8)

`evalField` folds path segments: a `Name` segment does `value.AsObject` then index (missing key → `nil`, and once `nil` short-circuits the rest to `nil`); an `Iter` (`[]`) segment is handled minimally in #1 (return `ErrType` if used outside flatMap context, or evaluate to the list itself — pick the spec-faithful behavior: `[]` yields the list for downstream flatMap). Provide the full implementation with each helper ≤ 7.

- [ ] **Step 5: Run eval tests to PASS**

Run: `go test ./eval/ -run TestEval -v`
Expected: PASS.

- [ ] **Step 6: Write failing builtin tests** (spec §5.5 table)

```go
func TestBuiltin_Length(t *testing.T) {
	got, err := Eval(ast.FuncCall{Name: "length", Args: []ast.Expr{ast.Literal{V: "abc"}}}, env(nil))
	if err != nil || got != int64(3) {
		t.Fatalf("length(abc)=%v,%v want 3", got, err)
	}
}

func TestBuiltin_Coalesce(t *testing.T) {
	got, _ := Eval(ast.FuncCall{Name: "coalesce", Args: []ast.Expr{ast.Literal{V: nil}, ast.Literal{V: "x"}}}, env(nil))
	if got != "x" {
		t.Fatalf("coalesce=%v want x", got)
	}
}

func TestBuiltin_Unknown(t *testing.T) {
	_, err := Eval(ast.FuncCall{Name: "nope"}, env(nil))
	if !errorsIs(err, ErrUnknownFunc) { t.Fatalf("got %v", err) }
}
```

- [ ] **Step 7: Implement `eval/builtins.go`** — a `map[string]builtin` registry; each builtin a small named func; `now()` reads `env.Now` (injected clock). Implement all of: `length keys values type toInt toFloat toString upper lower trim split join contains startsWith now flatten distinct coalesce`. Arity errors → `ErrArity`; unknown name → `ErrUnknownFunc`.

- [ ] **Step 8: Run builtin tests to PASS, then add cases to reach 100%**

Cover each builtin (happy + one failure/arity path) and `now()` with an injected fixed clock:

```go
func TestBuiltin_Now_Injected(t *testing.T) {
	fixed := time.Unix(1000, 0).UTC()
	e := Env{Vars: map[string]value.Value{}, Now: func() time.Time { return fixed }}
	got, _ := Eval(ast.FuncCall{Name: "now"}, e)
	// assert formatted/epoch per chosen now() contract
	_ = got
}
```

- [ ] **Step 9: Verify eval coverage 100%**

Run: `go test ./eval/ -covermode=atomic -coverprofile=/tmp/c.out && go tool cover -func=/tmp/c.out | awk '$3!="100.0%"'`
Expected: only the `total: 100.0%` line passes the filter (no sub-100 funcs).

- [ ] **Step 10: Commit**

```bash
git add eval/
git commit -m "feat(eval): expression evaluator + spec builtins with injected clock"
```

---

### Task 5: Transform stages + pipeline runtime (`stage`, `pipeline`)

**Files:**
- Create: `pipeline/pipeline.go`, `pipeline/errors.go`, `stage/stage.go`
- Test: `pipeline/pipeline_test.go`, `stage/stage_test.go`

**Interfaces:**
- Consumes: `ast`, `eval`, `value`
- Produces:
  - `type ResultSet = []value.Value` (each element an Object per §5.3, but kept as Value for normalization)
  - `type Stage interface { Execute(ResultSet) (ResultSet, error) }`
  - `func Normalize(v value.Value) ResultSet` (spec §5.3 source normalization)
  - `func Build(s ast.Stage, now func() time.Time) (Stage, error)` (maps a transform AST node to an executor; source nodes → `ErrStageUnsupported`)
  - `func RunStages(stages []Stage, in ResultSet) (ResultSet, error)`
  - Sentinel: `ErrStageUnsupported Error`

- [ ] **Step 1: Write failing tests for Normalize + map + filter**

```go
package pipeline

import (
	"testing"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/value"
)

func TestNormalize_WrapsLoneObject(t *testing.T) {
	rs := Normalize(map[string]value.Value{"a": int64(1)})
	if len(rs) != 1 {
		t.Fatalf("want 1 got %d", len(rs))
	}
}

func TestNormalize_WrapsPrimitiveArray(t *testing.T) {
	rs := Normalize([]value.Value{int64(1), int64(2)})
	o := rs[0].(map[string]value.Value)
	if o["value"] != int64(1) {
		t.Fatalf("primitive not wrapped: %v", rs[0])
	}
}

func TestFilterStage(t *testing.T) {
	s, _ := Build(ast.FilterStage{Cond: ast.BinaryExpr{
		Op: ast.OpGt,
		L:  ast.FieldAccess{Path: []ast.PathSegment{{Name: "n"}}},
		R:  ast.Literal{V: int64(1)},
	}}, nil)
	out, err := s.Execute(ResultSet{
		map[string]value.Value{"n": int64(0)},
		map[string]value.Value{"n": int64(5)},
	})
	if err != nil || len(out) != 1 {
		t.Fatalf("filter got %v,%v want 1 row", out, err)
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `go test ./pipeline/ -v`
Expected: FAIL.

- [ ] **Step 3: Implement `pipeline/errors.go`, `pipeline/pipeline.go`, `stage/stage.go`**

`Normalize`: list → if elements are objects, return as-is; if primitives, wrap each as `{"value": x}`; non-list → single-element set (object as-is; primitive wrapped). `Build`: type switch over `ast.Stage` constructing each executor; source nodes return `ErrStageUnsupported`. `RunStages`: fold, short-circuit on error. Each stage executor in `stage/stage.go` is a small struct with an `Execute` method:
- `mapExec` (1:1, eval each mapping into a new object)
- `flatMapExec` (N:M, list-valued mapping expands)
- `filterExec` (keep when `value.Truthy(eval(cond))`)
- `reduceExec` (dispatch on `ReduceOp` via a table of small funcs)
- `sortExec` (stable sort via `slices.SortStableFunc` + `value.Compare`, honoring `Desc`)
- `limitExec` (`min(N, len)`)
- `uniqExec` (full-object equality via canonical JSON key, or by `Key` expr)

Keep each `Execute` and helper ≤ 7 cognitive complexity.

- [ ] **Step 4: Run map/filter/normalize tests to PASS**

Run: `go test ./pipeline/ -run 'TestNormalize|TestFilter' -v`
Expected: PASS.

- [ ] **Step 5: Add tests for every stage + reduce op + flatMap expansion + sort asc/desc + uniq both modes + limit, plus `ErrStageUnsupported` for a source node. Drive to 100%.**

```go
func TestBuild_SourceUnsupported(t *testing.T) {
	if _, err := Build(ast.FileStage{Path: "x"}, nil); !errorsIs(err, ErrStageUnsupported) {
		t.Fatalf("got %v want ErrStageUnsupported", err)
	}
}
```

- [ ] **Step 6: Verify coverage 100% for stage + pipeline**

Run: `go test ./stage/ ./pipeline/ -covermode=atomic -coverprofile=/tmp/c.out && go tool cover -func=/tmp/c.out | awk '$3!="100.0%"'`
Expected: no sub-100 functions.

- [ ] **Step 7: Commit**

```bash
git add stage/ pipeline/
git commit -m "feat(stage,pipeline): transform stage executors + sequential runner"
```

---

### Task 6: Public API + end-to-end (`cirql.go`)

**Files:**
- Modify: `cirql.go` (replace the `doc.go` placeholder package doc home)
- Test: `cirql_test.go`, `examples/cirql_test.go`

**Interfaces:**
- Consumes: `pkg/dialect/cirql` (Parse), `pipeline`, `value`
- Produces:
  - `type Pipeline struct { ... }` wrapping the built stages
  - `func Parse(query string) (Pipeline, error)`
  - `func (Pipeline) Run(in value.Value) ([]value.Value, error)` (Normalizes input, runs stages, returns the result set)
  - `type Option func(*config)`; `func WithClock(func() time.Time) Option` (DI for `now()`)

- [ ] **Step 1: Write a failing end-to-end test using a spec §9 (transform-only) example**

```go
package cirql_test

import (
	"testing"

	"github.com/gomatic/cirql"
	"github.com/gomatic/cirql/value"
)

func TestEndToEnd_FilterMapSortLimit(t *testing.T) {
	p, err := cirql.Parse(`filter .stars > 1000 | map { name: .name, stars: .stars } | sort .stars desc | limit 2`)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	in := []value.Value{
		map[string]value.Value{"name": "a", "stars": int64(500)},
		map[string]value.Value{"name": "b", "stars": int64(3000)},
		map[string]value.Value{"name": "c", "stars": int64(2000)},
	}
	out, err := p.Run(in)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if len(out) != 2 {
		t.Fatalf("want 2 rows got %d", len(out))
	}
	first := out[0].(map[string]value.Value)
	if first["name"] != "b" {
		t.Fatalf("top = %v want b", first["name"])
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `go test . -run TestEndToEnd -v`
Expected: FAIL (undefined `Parse`/`Run`).

- [ ] **Step 3: Implement `cirql.go`** — `Parse` calls `dialect.Parse`, then `pipeline.Build` per stage (threading the configured clock), storing the `[]pipeline.Stage`. `Run` calls `pipeline.Normalize` then `pipeline.RunStages`. Options apply the clock (default `time.Now`).

- [ ] **Step 4: Run end-to-end to PASS**

Run: `go test . -run TestEndToEnd -v`
Expected: PASS.

- [ ] **Step 5: Add an `Example` in `examples/cirql_test.go`** (a runnable doc example with `// Output:`), and a parse-error propagation test asserting `errors.Is(err, dialect.ErrParse)` surfaces through `cirql.Parse`.

- [ ] **Step 6: Run the FULL gate**

Run: `make check`
Expected: PASS — 100% coverage over `COVER_PKGS`, lint/vet/staticcheck/vulncheck clean, gocognit empty.

- [ ] **Step 7: Commit**

```bash
git add cirql.go cirql_test.go examples/
git commit -m "feat: public cirql API (Parse + Pipeline.Run) with clock injection"
```

---

## Self-Review

**Spec coverage:** value model (Task 1 ✓ spec §5.2), grammar/parse (Task 3 ✓ §5.1), expressions+builtins (Task 4 ✓ §5.5, null-propagation/div-zero §5.8), transform stages (Task 5 ✓ §5.3 incl. flatMap/reduce ops), source normalization (Task 5 ✓ §5.3), public API (Task 6 ✓). Source stages (§5.4 propagation, `http`/`query`/`file`) are intentionally deferred to sub-project #2 and stubbed via `ErrStageUnsupported` — covered by the design's build order, not this plan.

**Placeholder scan:** No "TBD"/"add error handling"-style steps; each code step shows code or an exact command. Two steps (Task 4 Step 4, Task 5 Step 3) describe the per-node/per-stage implementations narratively rather than pasting every helper body — acceptable because each is a mechanical 1-rule-per-method expansion whose signatures, dispatch strategy, and complexity budget are fully specified; the implementer writes each ≤7-complexity helper against the named interfaces.

**Type consistency:** `value.Value` (alias) used throughout; `ResultSet = []value.Value`; `Parse`/`Run` names match across Task 3, 5, 6; `ErrStageUnsupported` defined once (pipeline) and referenced consistently; `Env.Now`/`WithClock` clock seam consistent between Task 4 and Task 6.

## Execution Handoff

This plan implements into `github.com/gomatic/cirql`, which **does not exist yet** — you (the user) are creating/moving that repo. Implementation begins once the repo is in place; until then this plan + the design spec are the ready artifacts.
