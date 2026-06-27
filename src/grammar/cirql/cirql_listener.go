// Code generated from cirql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package cirqlgrammar // cirql
import "github.com/antlr4-go/antlr/v4"

// cirqlListener is a complete listener for a parse tree produced by cirqlParser.
type cirqlListener interface {
	antlr.ParseTreeListener

	// EnterPipeline is called when entering the pipeline production.
	EnterPipeline(c *PipelineContext)

	// EnterStage is called when entering the stage production.
	EnterStage(c *StageContext)

	// EnterSourceStage is called when entering the sourceStage production.
	EnterSourceStage(c *SourceStageContext)

	// EnterQueryStage is called when entering the queryStage production.
	EnterQueryStage(c *QueryStageContext)

	// EnterQueryBody is called when entering the queryBody production.
	EnterQueryBody(c *QueryBodyContext)

	// EnterSelectionSet is called when entering the selectionSet production.
	EnterSelectionSet(c *SelectionSetContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterHttpStage is called when entering the httpStage production.
	EnterHttpStage(c *HttpStageContext)

	// EnterFileStage is called when entering the fileStage production.
	EnterFileStage(c *FileStageContext)

	// EnterStdinStage is called when entering the stdinStage production.
	EnterStdinStage(c *StdinStageContext)

	// EnterTransformStage is called when entering the transformStage production.
	EnterTransformStage(c *TransformStageContext)

	// EnterMapStage is called when entering the mapStage production.
	EnterMapStage(c *MapStageContext)

	// EnterFlatMapStage is called when entering the flatMapStage production.
	EnterFlatMapStage(c *FlatMapStageContext)

	// EnterMapping is called when entering the mapping production.
	EnterMapping(c *MappingContext)

	// EnterFilterStage is called when entering the filterStage production.
	EnterFilterStage(c *FilterStageContext)

	// EnterReduceStage is called when entering the reduceStage production.
	EnterReduceStage(c *ReduceStageContext)

	// EnterReduceOp is called when entering the reduceOp production.
	EnterReduceOp(c *ReduceOpContext)

	// EnterSortStage is called when entering the sortStage production.
	EnterSortStage(c *SortStageContext)

	// EnterLimitStage is called when entering the limitStage production.
	EnterLimitStage(c *LimitStageContext)

	// EnterUniqStage is called when entering the uniqStage production.
	EnterUniqStage(c *UniqStageContext)

	// EnterMulExpr is called when entering the MulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterAndExpr is called when entering the AndExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterLitExpr is called when entering the LitExpr production.
	EnterLitExpr(c *LitExprContext)

	// EnterCmpExpr is called when entering the CmpExpr production.
	EnterCmpExpr(c *CmpExprContext)

	// EnterVarExpr is called when entering the VarExpr production.
	EnterVarExpr(c *VarExprContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterAddExpr is called when entering the AddExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterFieldExpr is called when entering the FieldExpr production.
	EnterFieldExpr(c *FieldExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterUnaryExpr is called when entering the UnaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterFieldAccess is called when entering the fieldAccess production.
	EnterFieldAccess(c *FieldAccessContext)

	// EnterPathSeg is called when entering the pathSeg production.
	EnterPathSeg(c *PathSegContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterArgValue is called when entering the argValue production.
	EnterArgValue(c *ArgValueContext)

	// EnterObjectLit is called when entering the objectLit production.
	EnterObjectLit(c *ObjectLitContext)

	// EnterListLit is called when entering the listLit production.
	EnterListLit(c *ListLitContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitPipeline is called when exiting the pipeline production.
	ExitPipeline(c *PipelineContext)

	// ExitStage is called when exiting the stage production.
	ExitStage(c *StageContext)

	// ExitSourceStage is called when exiting the sourceStage production.
	ExitSourceStage(c *SourceStageContext)

	// ExitQueryStage is called when exiting the queryStage production.
	ExitQueryStage(c *QueryStageContext)

	// ExitQueryBody is called when exiting the queryBody production.
	ExitQueryBody(c *QueryBodyContext)

	// ExitSelectionSet is called when exiting the selectionSet production.
	ExitSelectionSet(c *SelectionSetContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitHttpStage is called when exiting the httpStage production.
	ExitHttpStage(c *HttpStageContext)

	// ExitFileStage is called when exiting the fileStage production.
	ExitFileStage(c *FileStageContext)

	// ExitStdinStage is called when exiting the stdinStage production.
	ExitStdinStage(c *StdinStageContext)

	// ExitTransformStage is called when exiting the transformStage production.
	ExitTransformStage(c *TransformStageContext)

	// ExitMapStage is called when exiting the mapStage production.
	ExitMapStage(c *MapStageContext)

	// ExitFlatMapStage is called when exiting the flatMapStage production.
	ExitFlatMapStage(c *FlatMapStageContext)

	// ExitMapping is called when exiting the mapping production.
	ExitMapping(c *MappingContext)

	// ExitFilterStage is called when exiting the filterStage production.
	ExitFilterStage(c *FilterStageContext)

	// ExitReduceStage is called when exiting the reduceStage production.
	ExitReduceStage(c *ReduceStageContext)

	// ExitReduceOp is called when exiting the reduceOp production.
	ExitReduceOp(c *ReduceOpContext)

	// ExitSortStage is called when exiting the sortStage production.
	ExitSortStage(c *SortStageContext)

	// ExitLimitStage is called when exiting the limitStage production.
	ExitLimitStage(c *LimitStageContext)

	// ExitUniqStage is called when exiting the uniqStage production.
	ExitUniqStage(c *UniqStageContext)

	// ExitMulExpr is called when exiting the MulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitAndExpr is called when exiting the AndExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitLitExpr is called when exiting the LitExpr production.
	ExitLitExpr(c *LitExprContext)

	// ExitCmpExpr is called when exiting the CmpExpr production.
	ExitCmpExpr(c *CmpExprContext)

	// ExitVarExpr is called when exiting the VarExpr production.
	ExitVarExpr(c *VarExprContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitAddExpr is called when exiting the AddExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitFieldExpr is called when exiting the FieldExpr production.
	ExitFieldExpr(c *FieldExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitUnaryExpr is called when exiting the UnaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitFieldAccess is called when exiting the fieldAccess production.
	ExitFieldAccess(c *FieldAccessContext)

	// ExitPathSeg is called when exiting the pathSeg production.
	ExitPathSeg(c *PathSegContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitArgValue is called when exiting the argValue production.
	ExitArgValue(c *ArgValueContext)

	// ExitObjectLit is called when exiting the objectLit production.
	ExitObjectLit(c *ObjectLitContext)

	// ExitListLit is called when exiting the listLit production.
	ExitListLit(c *ListLitContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
