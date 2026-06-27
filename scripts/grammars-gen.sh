#!/usr/bin/env bash
# Generate a Go parser for every dialect grammar (pkg/dialect/<name>/*.g4) into
# src/grammar/<name> via ANTLR. ANTLR's Java codegen runs only here, in a pinned
# Docker image; the generated Go is COMMITTED so normal build/test/CI stay
# pure-Go and Docker-free — only this script (grammar authoring) needs Docker.
set -euo pipefail
CDPATH= cd "$(dirname "${BASH_SOURCE[0]}")/.." >/dev/null

go="${GO:-go}"
antlr_image="cirql-antlr:4.13.2"
built_antlr=""

build_antlr() {
	[[ -n "$built_antlr" ]] && return
	docker build -q -t "$antlr_image" docker/antlr >/dev/null
	built_antlr=1
}

shopt -s nullglob
for dir in pkg/dialect/*/; do
	name="$(basename "$dir")"
	g4s=("$dir"*.g4)
	[[ ${#g4s[@]} -gt 0 ]] || continue

	build_antlr
	out="src/grammar/$name"
	rm -rf "$out"
	mkdir -p "$out"
	# Run with the grammar's own directory as cwd and an absolute -o so ANTLR
	# emits flat into src/grammar/<name> instead of mirroring the input path.
	bases=()
	for g in "${g4s[@]}"; do bases+=("$(basename "$g")"); done
	docker run --rm -v "$PWD:/work" -w "/work/$dir" "$antlr_image" \
		-Dlanguage=Go -package "${name}grammar" -o "/work/$out" "${bases[@]}"
	rm -f "$out"/*.interp "$out"/*.tokens
	# ANTLR 4.13's Go output carries artifacts the standard tools flag even though
	# the files are generated: an unreachable `goto errorExit` guard (go vet) and a
	# redundant pre-loop `_la` assignment for one-or-more subrules (staticcheck
	# SA4006). Strip both so the committed parsers pass the gate, then reformat.
	for gen in "$out"/*_parser.go; do
		awk -f scripts/antlr-vet-fix.awk "$gen" >"$gen.tmp" && mv "$gen.tmp" "$gen"
		awk -f scripts/antlr-staticcheck-fix.awk "$gen" >"$gen.tmp" && mv "$gen.tmp" "$gen"
	done
	"$go" tool gofumpt -w "$out"
	echo "antlr: $out (${bases[*]})"
done
