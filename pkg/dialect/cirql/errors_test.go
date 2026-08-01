package cirql

import (
	"errors"
	"strings"
	"testing"
)

// TestMaxQueryBytesBoundsNestingBelowTheStackLimit is a SECURITY test, and the
// only kind that can verify this bound. The generated parser recurses once per
// nesting level, and a goroutine stack overflow is a fatal error the runtime
// will not let anyone recover from — so if 256 KiB of pure nesting could still
// overflow, the guard would be decorative and any caller accepting untrusted
// query text could be killed by one request.
//
// The fixture is a query at the largest length the guard admits, made entirely
// of nesting, so it exercises the deepest recursion the bound permits. It must
// PARSE — a syntax error would mean the parser stopped early and the test
// proved nothing about depth, which is exactly how a first attempt at this test
// passed while recursing three levels.
func TestMaxQueryBytesBoundsNestingBelowTheStackLimit(t *testing.T) {
	const prefix = "filter "
	depth := (MaxQueryBytes - len(prefix) - 1) / 2
	query := prefix + strings.Repeat("(", depth) + "1" + strings.Repeat(")", depth)

	if len(query) > MaxQueryBytes {
		t.Fatalf("fixture of %d bytes exceeds the limit it is meant to sit under", len(query))
	}

	pipeline, err := Parse(Query(query))
	if err != nil {
		t.Fatalf("a maximally-nested query at the limit must parse, not error: %v", err)
	}
	if len(pipeline.Stages) != 1 {
		t.Fatalf("got %d stages, want the one filter stage — the parser must have consumed the nesting",
			len(pipeline.Stages))
	}
}

// TestErrQueryTooLargeRejectsBeforeParsing pins the guard's other half: the
// rejection happens BEFORE the parser sees the text, so the check cannot be
// satisfied by a parser that survives one particular oversized input. One byte
// over the limit is refused, and refused with a matchable sentinel rather than
// a generic parse error, so a caller can tell "too big" from "malformed".
func TestErrQueryTooLargeRejectsBeforeParsing(t *testing.T) {
	oversized := Query(strings.Repeat("a", MaxQueryBytes+1))

	_, err := Parse(oversized)

	if !errors.Is(err, ErrQueryTooLarge) {
		t.Fatalf("got %v, want ErrQueryTooLarge", err)
	}
	if errors.Is(err, ErrParse) {
		t.Fatal("an oversized query must not be reported as a syntax error; it was never parsed")
	}
}

// TestMaxQueryBytesDoesNotConstrainARealQuery pins the sizing claim — "query
// text is DSL, not data" — that justifies the limit being safe to impose. A
// realistic query must sit orders of magnitude under it; if a plausible query
// ever approached the bound, the guard would start rejecting legitimate use and
// the right answer would be a depth counter, not a longer byte limit.
func TestMaxQueryBytesDoesNotConstrainARealQuery(t *testing.T) {
	realistic := Query(`stdin | filter (.age > 30 && .name != "") | ` +
		`map {name: .name, decade: (.age / 10)} | sort .decade | uniq .name | limit 100`)

	if len(realistic)*100 > MaxQueryBytes {
		t.Fatalf("a realistic query is %d bytes, within 100x of the %d-byte limit; the limit is too tight",
			len(realistic), MaxQueryBytes)
	}
	if _, err := Parse(realistic); err != nil {
		t.Fatalf("the realistic fixture must be a valid query, or the margin above means nothing: %v", err)
	}
}
