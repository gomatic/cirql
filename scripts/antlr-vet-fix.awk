# Strip ANTLR 4.13's intentionally-dead `goto errorExit` guard from generated Go
# so the committed parsers pass `go vet` (which, unlike staticcheck/golangci-lint,
# does not skip "DO NOT EDIT" files and flags the dead goto as unreachable code).
#
# ANTLR appends, at each rule function's tail, an unreachable
#     return localctx
#     goto errorExit // Trick to prevent compiler error if the label is not used
# whose sole purpose is to mark the errorExit: label "used" in rules that have no
# real (reachable) `goto errorExit`. Per function we drop that trick line; when
# the function has no real goto we also drop the now-unused `errorExit:` label —
# its cleanup block still runs by fall-through, so behaviour is unchanged.
#
# The transform is scoped to ANTLR's fixed (pinned 4.13.2) output format.

# Only ANTLR rule functions carry the trick; they are uniquely identified by
# their `(localctx I…Context)` return (accessor one-liners do not have it), so
# scoping to them keeps the `^}` end-of-function marker reliable.
/^func \(p .*\(localctx / { infunc = 1; n = 0; real = 0 }

infunc {
	buf[++n] = $0
	# A real goto is the bare `…goto errorExit` emitted inside error checks; the
	# trick line ends in a comment, the label ends in a colon — neither matches.
	if ($0 ~ /[[:space:]]goto errorExit$/) real = 1
	if ($0 ~ /^}/) {
		for (i = 1; i <= n; i++) {
			if (buf[i] ~ /goto errorExit \/\/ Trick/) continue
			if (real == 0 && buf[i] ~ /^errorExit:$/) continue
			print buf[i]
		}
		infunc = 0
	}
	next
}

{ print }
