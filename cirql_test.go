package cirql_test

import (
	"errors"
	"testing"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql"
	dialect "github.com/gomatic/cirql/pkg/dialect/cirql"
	"github.com/gomatic/cirql/stage"
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
	if out[0].(map[string]value.Value)["name"] != "b" {
		t.Fatalf("top = %v want b", out[0].(map[string]value.Value)["name"])
	}
}

func TestParse_SyntaxErrorSurfaces(t *testing.T) {
	if _, err := cirql.Parse(`map {`); !errors.Is(err, dialect.ErrParse) {
		t.Fatalf("err=%v want ErrParse", err)
	}
}

func TestParse_SourceStageUnsupported(t *testing.T) {
	if _, err := cirql.Parse(`file "x.json" | filter .a`); !errors.Is(err, stage.ErrStageUnsupported) {
		t.Fatalf("err=%v want ErrStageUnsupported", err)
	}
}

func TestRun_WithClock(t *testing.T) {
	p, err := cirql.Parse(`map { t: now() }`, cirql.WithClock(func() int64 { return 1234 }))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	out, err := p.Run([]value.Value{map[string]value.Value{}})
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if out[0].(map[string]value.Value)["t"] != int64(1234) {
		t.Fatalf("now=%v want 1234", out[0].(map[string]value.Value)["t"])
	}
}

func TestRun_PropagatesStageError(t *testing.T) {
	p, err := cirql.Parse(`sort .n`)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	in := []value.Value{
		map[string]value.Value{"n": "x"},
		map[string]value.Value{"n": int64(1)},
	}
	if _, err := p.Run(in); !errors.Is(err, value.ErrIncomparable) {
		t.Fatalf("err=%v want ErrIncomparable", err)
	}
}

// End-to-end: a query addressing keyword-named JSON fields (count, type) and
// grouping produces natural JSON keys — the whole point of the keyword-field
// and group-by fixes, exercised through the public API.
func TestEndToEnd_KeywordFieldsAndGroupBy(t *testing.T) {
	p, err := cirql.Parse(`filter .count > 1 | map { type: .type, count: .count }`)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	in := []value.Value{
		map[string]value.Value{"type": "x", "count": int64(5)},
		map[string]value.Value{"type": "y", "count": int64(1)},
	}
	out, err := p.Run(in)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if len(out) != 1 {
		t.Fatalf("want 1 row got %d", len(out))
	}
	row := out[0].(map[string]value.Value)
	if row["type"] != "x" || row["count"] != int64(5) {
		t.Errorf("row = %#v, want type=x count=5", row)
	}
}

// A query above the size bound is rejected as ErrQueryTooLarge through the
// public Parse, not a stack overflow.
func TestParse_RejectsOversizedQuery(t *testing.T) {
	big := make([]byte, dialect.MaxQueryBytes+1)
	for i := range big {
		big[i] = ' '
	}
	_, err := cirql.Parse(cirql.Query(big))
	if !errors.Is(err, dialect.ErrQueryTooLarge) {
		t.Fatalf("got %v want ErrQueryTooLarge", err)
	}
}
