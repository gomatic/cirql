# Strip ANTLR 4.13's redundant pre-loop `_la` assignment from generated Go so the
# committed parsers pass `staticcheck` (which, like go vet, does analyze generated
# files). Companion to antlr-vet-fix.awk; run as a second pass over *_parser.go.
#
# For a one-or-more subrule `(X)+` over a token set, ANTLR emits a do-while:
#     _la = p.GetTokenStream().LA(1)        <- assigned, then immediately
#     for ok := true; ok; ok = (... _la ...) {   overwritten at the loop tail
#         ...
#         _la = p.GetTokenStream().LA(1)         before the condition reads it,
#     }                                          so the pre-loop value is dead
# which staticcheck flags as SA4006 ("this value of _la is never used"). The
# zero-or-more form `(X)*` instead emits `for _la cond {...}` whose pre-loop
# assignment IS read by the condition — so we drop the pre-loop `_la` ONLY when
# the next non-blank line is the `for ok := true;` do-while header, leaving the
# while-form untouched. Behaviour is unchanged: `_la` is reassigned inside the
# loop before its first use.

/^\t_la = p\.GetTokenStream\(\)\.LA\(1\)$/ {
	pending = $0
	blank = ""
	next
}

pending != "" {
	if ($0 ~ /^[[:space:]]*$/) {
		blank = $0
		next
	}
	if ($0 !~ /^\tfor ok := true;/) {
		print pending
	}
	if (blank != "") {
		print blank
	}
	pending = ""
	blank = ""
}

{ print }
