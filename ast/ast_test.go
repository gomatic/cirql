package ast

import "testing"

func TestNodesImplementInterfaces(t *testing.T) {
	stages := []Stage{
		MapStage{},
		FlatMapStage{},
		FilterStage{},
		ReduceStage{},
		SortStage{},
		LimitStage{},
		UniqStage{},
		StdinStage{},
		FileStage{},
		HTTPStage{},
		QueryStage{},
	}
	if len(stages) != 11 {
		t.Fatalf("expected 11 stage kinds, got %d", len(stages))
	}
	for _, s := range stages {
		s.isStage() // exercise the closed-set marker
	}
	exprs := []Expr{
		FieldAccess{}, VarRef{}, BinaryExpr{}, UnaryExpr{}, FuncCall{}, Literal{},
	}
	if len(exprs) != 6 {
		t.Fatalf("expected 6 expr kinds, got %d", len(exprs))
	}
	for _, e := range exprs {
		e.isExpr() // exercise the closed-set marker
	}
}

func TestOperatorConsts(t *testing.T) {
	if OpEq != "==" || OpAnd != "&&" || OpNot != "!" || OpGroupBy != "group_by" {
		t.Fatal("operator/op constants have unexpected values")
	}
}

// TestUniqStageKeyDistinguishesTheTwoDeduplicationModes names UniqStage's
// claim: "by Key when non-nil else whole-object". The nil Key is not a missing
// value to be defaulted — it IS the whole-object mode, and it is the only thing
// telling `uniq` apart from `uniq .name`. A constructor that filled in a
// default key, or an executor that treated nil as an empty path, would silently
// answer a different question than the query asked.
func TestUniqStageKeyDistinguishesTheTwoDeduplicationModes(t *testing.T) {
	if (UniqStage{}).Key != nil {
		t.Fatal("the zero UniqStage must carry a nil Key — that is whole-object mode")
	}

	keyed := UniqStage{Key: FieldAccess{Path: []PathSegment{{Name: "name"}}}}
	if keyed.Key == nil {
		t.Fatal("a keyed UniqStage must carry its key expression")
	}

	// The identity path "." is a real, non-nil key and must not be mistaken for
	// the absent one: `uniq .` asks to deduplicate by the whole object as a
	// VALUE, which is a key expression, not the keyless mode.
	identity := UniqStage{Key: FieldAccess{}}
	if identity.Key == nil {
		t.Fatal("an empty-path FieldAccess is still a key; only a nil Key means keyless")
	}
}
