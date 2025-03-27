// Code generated from BigC.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // BigC

import "github.com/antlr4-go/antlr/v4"

// BaseBigCListener is a complete listener for a parse tree produced by BigCParser.
type BaseBigCListener struct{}

var _ BigCListener = &BaseBigCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBigCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBigCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBigCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBigCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBigCListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBigCListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseBigCListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseBigCListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterType is called when production type is entered.
func (s *BaseBigCListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseBigCListener) ExitType(ctx *TypeContext) {}

// EnterDeclarationRemainder is called when production declarationRemainder is entered.
func (s *BaseBigCListener) EnterDeclarationRemainder(ctx *DeclarationRemainderContext) {}

// ExitDeclarationRemainder is called when production declarationRemainder is exited.
func (s *BaseBigCListener) ExitDeclarationRemainder(ctx *DeclarationRemainderContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseBigCListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseBigCListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseBigCListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseBigCListener) ExitParameter(ctx *ParameterContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBigCListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBigCListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockItem is called when production blockItem is entered.
func (s *BaseBigCListener) EnterBlockItem(ctx *BlockItemContext) {}

// ExitBlockItem is called when production blockItem is exited.
func (s *BaseBigCListener) ExitBlockItem(ctx *BlockItemContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseBigCListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseBigCListener) ExitStatement(ctx *StatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseBigCListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseBigCListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterElseClause is called when production elseClause is entered.
func (s *BaseBigCListener) EnterElseClause(ctx *ElseClauseContext) {}

// ExitElseClause is called when production elseClause is exited.
func (s *BaseBigCListener) ExitElseClause(ctx *ElseClauseContext) {}

// EnterNonIfStatement is called when production nonIfStatement is entered.
func (s *BaseBigCListener) EnterNonIfStatement(ctx *NonIfStatementContext) {}

// ExitNonIfStatement is called when production nonIfStatement is exited.
func (s *BaseBigCListener) ExitNonIfStatement(ctx *NonIfStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseBigCListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseBigCListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseBigCListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseBigCListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBigCListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBigCListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseBigCListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseBigCListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignmentRest is called when production assignmentRest is entered.
func (s *BaseBigCListener) EnterAssignmentRest(ctx *AssignmentRestContext) {}

// ExitAssignmentRest is called when production assignmentRest is exited.
func (s *BaseBigCListener) ExitAssignmentRest(ctx *AssignmentRestContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseBigCListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseBigCListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseBigCListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseBigCListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterLogicalOrRest is called when production logicalOrRest is entered.
func (s *BaseBigCListener) EnterLogicalOrRest(ctx *LogicalOrRestContext) {}

// ExitLogicalOrRest is called when production logicalOrRest is exited.
func (s *BaseBigCListener) ExitLogicalOrRest(ctx *LogicalOrRestContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseBigCListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseBigCListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterLogicalAndRest is called when production logicalAndRest is entered.
func (s *BaseBigCListener) EnterLogicalAndRest(ctx *LogicalAndRestContext) {}

// ExitLogicalAndRest is called when production logicalAndRest is exited.
func (s *BaseBigCListener) ExitLogicalAndRest(ctx *LogicalAndRestContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseBigCListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseBigCListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterEqualityRest is called when production equalityRest is entered.
func (s *BaseBigCListener) EnterEqualityRest(ctx *EqualityRestContext) {}

// ExitEqualityRest is called when production equalityRest is exited.
func (s *BaseBigCListener) ExitEqualityRest(ctx *EqualityRestContext) {}

// EnterEqualityOperator is called when production equalityOperator is entered.
func (s *BaseBigCListener) EnterEqualityOperator(ctx *EqualityOperatorContext) {}

// ExitEqualityOperator is called when production equalityOperator is exited.
func (s *BaseBigCListener) ExitEqualityOperator(ctx *EqualityOperatorContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseBigCListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseBigCListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterComparisonRest is called when production comparisonRest is entered.
func (s *BaseBigCListener) EnterComparisonRest(ctx *ComparisonRestContext) {}

// ExitComparisonRest is called when production comparisonRest is exited.
func (s *BaseBigCListener) ExitComparisonRest(ctx *ComparisonRestContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseBigCListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseBigCListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterAdditionExpression is called when production additionExpression is entered.
func (s *BaseBigCListener) EnterAdditionExpression(ctx *AdditionExpressionContext) {}

// ExitAdditionExpression is called when production additionExpression is exited.
func (s *BaseBigCListener) ExitAdditionExpression(ctx *AdditionExpressionContext) {}

// EnterAdditionExpressionRest is called when production additionExpressionRest is entered.
func (s *BaseBigCListener) EnterAdditionExpressionRest(ctx *AdditionExpressionRestContext) {}

// ExitAdditionExpressionRest is called when production additionExpressionRest is exited.
func (s *BaseBigCListener) ExitAdditionExpressionRest(ctx *AdditionExpressionRestContext) {}

// EnterAddSubtractOperator is called when production addSubtractOperator is entered.
func (s *BaseBigCListener) EnterAddSubtractOperator(ctx *AddSubtractOperatorContext) {}

// ExitAddSubtractOperator is called when production addSubtractOperator is exited.
func (s *BaseBigCListener) ExitAddSubtractOperator(ctx *AddSubtractOperatorContext) {}

// EnterMultiplicationExpression is called when production multiplicationExpression is entered.
func (s *BaseBigCListener) EnterMultiplicationExpression(ctx *MultiplicationExpressionContext) {}

// ExitMultiplicationExpression is called when production multiplicationExpression is exited.
func (s *BaseBigCListener) ExitMultiplicationExpression(ctx *MultiplicationExpressionContext) {}

// EnterMultiplicationExpressionRest is called when production multiplicationExpressionRest is entered.
func (s *BaseBigCListener) EnterMultiplicationExpressionRest(ctx *MultiplicationExpressionRestContext) {
}

// ExitMultiplicationExpressionRest is called when production multiplicationExpressionRest is exited.
func (s *BaseBigCListener) ExitMultiplicationExpressionRest(ctx *MultiplicationExpressionRestContext) {
}

// EnterMultDivModOperator is called when production multDivModOperator is entered.
func (s *BaseBigCListener) EnterMultDivModOperator(ctx *MultDivModOperatorContext) {}

// ExitMultDivModOperator is called when production multDivModOperator is exited.
func (s *BaseBigCListener) ExitMultDivModOperator(ctx *MultDivModOperatorContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseBigCListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseBigCListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseBigCListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseBigCListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseBigCListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseBigCListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterArrayAccess is called when production arrayAccess is entered.
func (s *BaseBigCListener) EnterArrayAccess(ctx *ArrayAccessContext) {}

// ExitArrayAccess is called when production arrayAccess is exited.
func (s *BaseBigCListener) ExitArrayAccess(ctx *ArrayAccessContext) {}

// EnterFunctionCallArgs is called when production functionCallArgs is entered.
func (s *BaseBigCListener) EnterFunctionCallArgs(ctx *FunctionCallArgsContext) {}

// ExitFunctionCallArgs is called when production functionCallArgs is exited.
func (s *BaseBigCListener) ExitFunctionCallArgs(ctx *FunctionCallArgsContext) {}

// EnterIncreaseDecrease is called when production increaseDecrease is entered.
func (s *BaseBigCListener) EnterIncreaseDecrease(ctx *IncreaseDecreaseContext) {}

// ExitIncreaseDecrease is called when production increaseDecrease is exited.
func (s *BaseBigCListener) ExitIncreaseDecrease(ctx *IncreaseDecreaseContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseBigCListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseBigCListener) ExitArgList(ctx *ArgListContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseBigCListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseBigCListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseBigCListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseBigCListener) ExitConstant(ctx *ConstantContext) {}
