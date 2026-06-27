// Code generated from cirql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package cirqlgrammar // cirql
import "github.com/antlr4-go/antlr/v4"

// BasecirqlListener is a complete listener for a parse tree produced by cirqlParser.
type BasecirqlListener struct{}

var _ cirqlListener = &BasecirqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecirqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecirqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecirqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecirqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPipeline is called when production pipeline is entered.
func (s *BasecirqlListener) EnterPipeline(ctx *PipelineContext) {}

// ExitPipeline is called when production pipeline is exited.
func (s *BasecirqlListener) ExitPipeline(ctx *PipelineContext) {}

// EnterStage is called when production stage is entered.
func (s *BasecirqlListener) EnterStage(ctx *StageContext) {}

// ExitStage is called when production stage is exited.
func (s *BasecirqlListener) ExitStage(ctx *StageContext) {}

// EnterSourceStage is called when production sourceStage is entered.
func (s *BasecirqlListener) EnterSourceStage(ctx *SourceStageContext) {}

// ExitSourceStage is called when production sourceStage is exited.
func (s *BasecirqlListener) ExitSourceStage(ctx *SourceStageContext) {}

// EnterQueryStage is called when production queryStage is entered.
func (s *BasecirqlListener) EnterQueryStage(ctx *QueryStageContext) {}

// ExitQueryStage is called when production queryStage is exited.
func (s *BasecirqlListener) ExitQueryStage(ctx *QueryStageContext) {}

// EnterQueryBody is called when production queryBody is entered.
func (s *BasecirqlListener) EnterQueryBody(ctx *QueryBodyContext) {}

// ExitQueryBody is called when production queryBody is exited.
func (s *BasecirqlListener) ExitQueryBody(ctx *QueryBodyContext) {}

// EnterSelectionSet is called when production selectionSet is entered.
func (s *BasecirqlListener) EnterSelectionSet(ctx *SelectionSetContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (s *BasecirqlListener) ExitSelectionSet(ctx *SelectionSetContext) {}

// EnterField is called when production field is entered.
func (s *BasecirqlListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BasecirqlListener) ExitField(ctx *FieldContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BasecirqlListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasecirqlListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasecirqlListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasecirqlListener) ExitArgument(ctx *ArgumentContext) {}

// EnterHttpStage is called when production httpStage is entered.
func (s *BasecirqlListener) EnterHttpStage(ctx *HttpStageContext) {}

// ExitHttpStage is called when production httpStage is exited.
func (s *BasecirqlListener) ExitHttpStage(ctx *HttpStageContext) {}

// EnterFileStage is called when production fileStage is entered.
func (s *BasecirqlListener) EnterFileStage(ctx *FileStageContext) {}

// ExitFileStage is called when production fileStage is exited.
func (s *BasecirqlListener) ExitFileStage(ctx *FileStageContext) {}

// EnterStdinStage is called when production stdinStage is entered.
func (s *BasecirqlListener) EnterStdinStage(ctx *StdinStageContext) {}

// ExitStdinStage is called when production stdinStage is exited.
func (s *BasecirqlListener) ExitStdinStage(ctx *StdinStageContext) {}

// EnterTransformStage is called when production transformStage is entered.
func (s *BasecirqlListener) EnterTransformStage(ctx *TransformStageContext) {}

// ExitTransformStage is called when production transformStage is exited.
func (s *BasecirqlListener) ExitTransformStage(ctx *TransformStageContext) {}

// EnterMapStage is called when production mapStage is entered.
func (s *BasecirqlListener) EnterMapStage(ctx *MapStageContext) {}

// ExitMapStage is called when production mapStage is exited.
func (s *BasecirqlListener) ExitMapStage(ctx *MapStageContext) {}

// EnterFlatMapStage is called when production flatMapStage is entered.
func (s *BasecirqlListener) EnterFlatMapStage(ctx *FlatMapStageContext) {}

// ExitFlatMapStage is called when production flatMapStage is exited.
func (s *BasecirqlListener) ExitFlatMapStage(ctx *FlatMapStageContext) {}

// EnterMapping is called when production mapping is entered.
func (s *BasecirqlListener) EnterMapping(ctx *MappingContext) {}

// ExitMapping is called when production mapping is exited.
func (s *BasecirqlListener) ExitMapping(ctx *MappingContext) {}

// EnterFilterStage is called when production filterStage is entered.
func (s *BasecirqlListener) EnterFilterStage(ctx *FilterStageContext) {}

// ExitFilterStage is called when production filterStage is exited.
func (s *BasecirqlListener) ExitFilterStage(ctx *FilterStageContext) {}

// EnterReduceStage is called when production reduceStage is entered.
func (s *BasecirqlListener) EnterReduceStage(ctx *ReduceStageContext) {}

// ExitReduceStage is called when production reduceStage is exited.
func (s *BasecirqlListener) ExitReduceStage(ctx *ReduceStageContext) {}

// EnterReduceOp is called when production reduceOp is entered.
func (s *BasecirqlListener) EnterReduceOp(ctx *ReduceOpContext) {}

// ExitReduceOp is called when production reduceOp is exited.
func (s *BasecirqlListener) ExitReduceOp(ctx *ReduceOpContext) {}

// EnterSortStage is called when production sortStage is entered.
func (s *BasecirqlListener) EnterSortStage(ctx *SortStageContext) {}

// ExitSortStage is called when production sortStage is exited.
func (s *BasecirqlListener) ExitSortStage(ctx *SortStageContext) {}

// EnterLimitStage is called when production limitStage is entered.
func (s *BasecirqlListener) EnterLimitStage(ctx *LimitStageContext) {}

// ExitLimitStage is called when production limitStage is exited.
func (s *BasecirqlListener) ExitLimitStage(ctx *LimitStageContext) {}

// EnterUniqStage is called when production uniqStage is entered.
func (s *BasecirqlListener) EnterUniqStage(ctx *UniqStageContext) {}

// ExitUniqStage is called when production uniqStage is exited.
func (s *BasecirqlListener) ExitUniqStage(ctx *UniqStageContext) {}

// EnterMulExpr is called when production MulExpr is entered.
func (s *BasecirqlListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production MulExpr is exited.
func (s *BasecirqlListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterAndExpr is called when production AndExpr is entered.
func (s *BasecirqlListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production AndExpr is exited.
func (s *BasecirqlListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterLitExpr is called when production LitExpr is entered.
func (s *BasecirqlListener) EnterLitExpr(ctx *LitExprContext) {}

// ExitLitExpr is called when production LitExpr is exited.
func (s *BasecirqlListener) ExitLitExpr(ctx *LitExprContext) {}

// EnterCmpExpr is called when production CmpExpr is entered.
func (s *BasecirqlListener) EnterCmpExpr(ctx *CmpExprContext) {}

// ExitCmpExpr is called when production CmpExpr is exited.
func (s *BasecirqlListener) ExitCmpExpr(ctx *CmpExprContext) {}

// EnterVarExpr is called when production VarExpr is entered.
func (s *BasecirqlListener) EnterVarExpr(ctx *VarExprContext) {}

// ExitVarExpr is called when production VarExpr is exited.
func (s *BasecirqlListener) ExitVarExpr(ctx *VarExprContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BasecirqlListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BasecirqlListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterAddExpr is called when production AddExpr is entered.
func (s *BasecirqlListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production AddExpr is exited.
func (s *BasecirqlListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterFieldExpr is called when production FieldExpr is entered.
func (s *BasecirqlListener) EnterFieldExpr(ctx *FieldExprContext) {}

// ExitFieldExpr is called when production FieldExpr is exited.
func (s *BasecirqlListener) ExitFieldExpr(ctx *FieldExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BasecirqlListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BasecirqlListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterUnaryExpr is called when production UnaryExpr is entered.
func (s *BasecirqlListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production UnaryExpr is exited.
func (s *BasecirqlListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterOrExpr is called when production OrExpr is entered.
func (s *BasecirqlListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production OrExpr is exited.
func (s *BasecirqlListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterFieldAccess is called when production fieldAccess is entered.
func (s *BasecirqlListener) EnterFieldAccess(ctx *FieldAccessContext) {}

// ExitFieldAccess is called when production fieldAccess is exited.
func (s *BasecirqlListener) ExitFieldAccess(ctx *FieldAccessContext) {}

// EnterPathSeg is called when production pathSeg is entered.
func (s *BasecirqlListener) EnterPathSeg(ctx *PathSegContext) {}

// ExitPathSeg is called when production pathSeg is exited.
func (s *BasecirqlListener) ExitPathSeg(ctx *PathSegContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BasecirqlListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BasecirqlListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasecirqlListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasecirqlListener) ExitVariable(ctx *VariableContext) {}

// EnterArgValue is called when production argValue is entered.
func (s *BasecirqlListener) EnterArgValue(ctx *ArgValueContext) {}

// ExitArgValue is called when production argValue is exited.
func (s *BasecirqlListener) ExitArgValue(ctx *ArgValueContext) {}

// EnterObjectLit is called when production objectLit is entered.
func (s *BasecirqlListener) EnterObjectLit(ctx *ObjectLitContext) {}

// ExitObjectLit is called when production objectLit is exited.
func (s *BasecirqlListener) ExitObjectLit(ctx *ObjectLitContext) {}

// EnterListLit is called when production listLit is entered.
func (s *BasecirqlListener) EnterListLit(ctx *ListLitContext) {}

// ExitListLit is called when production listLit is exited.
func (s *BasecirqlListener) ExitListLit(ctx *ListLitContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasecirqlListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasecirqlListener) ExitLiteral(ctx *LiteralContext) {}
