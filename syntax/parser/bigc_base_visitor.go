// Code generated from BigC.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // BigC

import "github.com/antlr4-go/antlr/v4"

type BaseBigCVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBigCVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitDeclarationRemainder(ctx *DeclarationRemainderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitBlockItem(ctx *BlockItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitElseClause(ctx *ElseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitNonIfStatement(ctx *NonIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitAssignmentRest(ctx *AssignmentRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitLogicalOrRest(ctx *LogicalOrRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitLogicalAndRest(ctx *LogicalAndRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitEqualityRest(ctx *EqualityRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitEqualityOperator(ctx *EqualityOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitComparisonRest(ctx *ComparisonRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitAdditionExpression(ctx *AdditionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitAdditionExpressionRest(ctx *AdditionExpressionRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitAddSubtractOperator(ctx *AddSubtractOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitMultiplicationExpression(ctx *MultiplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitMultiplicationExpressionRest(ctx *MultiplicationExpressionRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitMultDivModOperator(ctx *MultDivModOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitArrayAccess(ctx *ArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitFunctionCallArgs(ctx *FunctionCallArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitIncreaseDecrease(ctx *IncreaseDecreaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigCVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}
