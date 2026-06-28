package cirql_test

import (
	"errors"
	"testing"

	"github.com/gomatic/cirql"
	dialect "github.com/gomatic/cirql/pkg/dialect/cirql"
	"github.com/gomatic/cirql/stage"
	value "github.com/gomatic/go-json"
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
