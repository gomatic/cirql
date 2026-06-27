package pipeline

import (
	"errors"
	"testing"

	"github.com/gomatic/cirql/value"
)

func TestNormalize_ListOfObjects(t *testing.T) {
	in := []value.Value{
		map[string]value.Value{"a": int64(1)},
		map[string]value.Value{"a": int64(2)},
	}
	rs := Normalize(in)
	if len(rs) != 2 {
		t.Fatalf("len=%d want 2", len(rs))
	}
	if _, ok := rs[0].(map[string]value.Value); !ok {
		t.Fatalf("element 0 = %T want object", rs[0])
	}
}

func TestNormalize_ListOfPrimitives(t *testing.T) {
	rs := Normalize([]value.Value{int64(1), int64(2)})
	o, ok := rs[0].(map[string]value.Value)
	if !ok || o["value"] != int64(1) {
		t.Fatalf("primitive not wrapped: %#v", rs[0])
	}
}

func TestNormalize_SingleObject(t *testing.T) {
	rs := Normalize(map[string]value.Value{"a": int64(1)})
	if len(rs) != 1 {
		t.Fatalf("len=%d want 1", len(rs))
	}
	if _, ok := rs[0].(map[string]value.Value); !ok {
		t.Fatalf("element = %T want object", rs[0])
	}
}

func TestNormalize_SinglePrimitive(t *testing.T) {
	rs := Normalize("hello")
	o := rs[0].(map[string]value.Value)
	if o["value"] != "hello" {
		t.Fatalf("primitive not wrapped: %#v", rs[0])
	}
}

// okStage passes input through; failStage returns an error.
type okStage struct{}

func (okStage) Execute(in ResultSet) (ResultSet, error) { return in, nil }

type failStage struct{ err error }

func (s failStage) Execute(ResultSet) (ResultSet, error) { return nil, s.err }

func TestRunStages_Success(t *testing.T) {
	in := ResultSet{int64(1)}
	out, err := RunStages([]Stage{okStage{}, okStage{}}, in)
	if err != nil || len(out) != 1 {
		t.Fatalf("out=%v err=%v", out, err)
	}
}

func TestRunStages_PropagatesError(t *testing.T) {
	sentinel := errors.New("boom")
	_, err := RunStages([]Stage{okStage{}, failStage{err: sentinel}}, ResultSet{})
	if !errors.Is(err, sentinel) {
		t.Fatalf("err=%v want boom", err)
	}
}
