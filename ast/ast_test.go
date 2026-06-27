package ast

import "testing"

func TestNodesImplementInterfaces(t *testing.T) {
	stages := []Stage{
		MapStage{}, FlatMapStage{}, FilterStage{}, ReduceStage{},
		SortStage{}, LimitStage{}, UniqStage{},
		StdinStage{}, FileStage{}, HTTPStage{}, QueryStage{},
	}
	if len(stages) != 11 {
		t.Fatalf("expected 11 stage kinds, got %d", len(stages))
	}
	exprs := []Expr{
		FieldAccess{}, VarRef{}, BinaryExpr{}, UnaryExpr{}, FuncCall{}, Literal{},
	}
	if len(exprs) != 6 {
		t.Fatalf("expected 6 expr kinds, got %d", len(exprs))
	}
}

func TestOperatorConsts(t *testing.T) {
	if OpEq != "==" || OpAnd != "&&" || OpNot != "!" || OpGroupBy != "group_by" {
		t.Fatal("operator/op constants have unexpected values")
	}
}
