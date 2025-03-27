// Code generated from BigC.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // BigC

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BigCParser.
type BigCVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BigCParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by BigCParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by BigCParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by BigCParser#declarationRemainder.
	VisitDeclarationRemainder(ctx *DeclarationRemainderContext) interface{}

	// Visit a parse tree produced by BigCParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by BigCParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by BigCParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by BigCParser#blockItem.
	VisitBlockItem(ctx *BlockItemContext) interface{}

	// Visit a parse tree produced by BigCParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by BigCParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by BigCParser#elseClause.
	VisitElseClause(ctx *ElseClauseContext) interface{}

	// Visit a parse tree produced by BigCParser#nonIfStatement.
	VisitNonIfStatement(ctx *NonIfStatementContext) interface{}

	// Visit a parse tree produced by BigCParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by BigCParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by BigCParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#assignmentRest.
	VisitAssignmentRest(ctx *AssignmentRestContext) interface{}

	// Visit a parse tree produced by BigCParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by BigCParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#logicalOrRest.
	VisitLogicalOrRest(ctx *LogicalOrRestContext) interface{}

	// Visit a parse tree produced by BigCParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#logicalAndRest.
	VisitLogicalAndRest(ctx *LogicalAndRestContext) interface{}

	// Visit a parse tree produced by BigCParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#equalityRest.
	VisitEqualityRest(ctx *EqualityRestContext) interface{}

	// Visit a parse tree produced by BigCParser#equalityOperator.
	VisitEqualityOperator(ctx *EqualityOperatorContext) interface{}

	// Visit a parse tree produced by BigCParser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#comparisonRest.
	VisitComparisonRest(ctx *ComparisonRestContext) interface{}

	// Visit a parse tree produced by BigCParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by BigCParser#additionExpression.
	VisitAdditionExpression(ctx *AdditionExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#additionExpressionRest.
	VisitAdditionExpressionRest(ctx *AdditionExpressionRestContext) interface{}

	// Visit a parse tree produced by BigCParser#addSubtractOperator.
	VisitAddSubtractOperator(ctx *AddSubtractOperatorContext) interface{}

	// Visit a parse tree produced by BigCParser#multiplicationExpression.
	VisitMultiplicationExpression(ctx *MultiplicationExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#multiplicationExpressionRest.
	VisitMultiplicationExpressionRest(ctx *MultiplicationExpressionRestContext) interface{}

	// Visit a parse tree produced by BigCParser#multDivModOperator.
	VisitMultDivModOperator(ctx *MultDivModOperatorContext) interface{}

	// Visit a parse tree produced by BigCParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by BigCParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#arrayAccess.
	VisitArrayAccess(ctx *ArrayAccessContext) interface{}

	// Visit a parse tree produced by BigCParser#functionCallArgs.
	VisitFunctionCallArgs(ctx *FunctionCallArgsContext) interface{}

	// Visit a parse tree produced by BigCParser#increaseDecrease.
	VisitIncreaseDecrease(ctx *IncreaseDecreaseContext) interface{}

	// Visit a parse tree produced by BigCParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by BigCParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by BigCParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}
}
