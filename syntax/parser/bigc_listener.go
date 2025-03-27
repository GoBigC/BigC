// Code generated from BigC.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // BigC

import "github.com/antlr4-go/antlr/v4"

// BigCListener is a complete listener for a parse tree produced by BigCParser.
type BigCListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterDeclarationRemainder is called when entering the declarationRemainder production.
	EnterDeclarationRemainder(c *DeclarationRemainderContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockItem is called when entering the blockItem production.
	EnterBlockItem(c *BlockItemContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterElseClause is called when entering the elseClause production.
	EnterElseClause(c *ElseClauseContext)

	// EnterNonIfStatement is called when entering the nonIfStatement production.
	EnterNonIfStatement(c *NonIfStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignmentRest is called when entering the assignmentRest production.
	EnterAssignmentRest(c *AssignmentRestContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterLogicalOrRest is called when entering the logicalOrRest production.
	EnterLogicalOrRest(c *LogicalOrRestContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterLogicalAndRest is called when entering the logicalAndRest production.
	EnterLogicalAndRest(c *LogicalAndRestContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterEqualityRest is called when entering the equalityRest production.
	EnterEqualityRest(c *EqualityRestContext)

	// EnterEqualityOperator is called when entering the equalityOperator production.
	EnterEqualityOperator(c *EqualityOperatorContext)

	// EnterComparisonExpression is called when entering the comparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterComparisonRest is called when entering the comparisonRest production.
	EnterComparisonRest(c *ComparisonRestContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterAdditionExpression is called when entering the additionExpression production.
	EnterAdditionExpression(c *AdditionExpressionContext)

	// EnterAdditionExpressionRest is called when entering the additionExpressionRest production.
	EnterAdditionExpressionRest(c *AdditionExpressionRestContext)

	// EnterAddSubtractOperator is called when entering the addSubtractOperator production.
	EnterAddSubtractOperator(c *AddSubtractOperatorContext)

	// EnterMultiplicationExpression is called when entering the multiplicationExpression production.
	EnterMultiplicationExpression(c *MultiplicationExpressionContext)

	// EnterMultiplicationExpressionRest is called when entering the multiplicationExpressionRest production.
	EnterMultiplicationExpressionRest(c *MultiplicationExpressionRestContext)

	// EnterMultDivModOperator is called when entering the multDivModOperator production.
	EnterMultDivModOperator(c *MultDivModOperatorContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterArrayAccess is called when entering the arrayAccess production.
	EnterArrayAccess(c *ArrayAccessContext)

	// EnterFunctionCallArgs is called when entering the functionCallArgs production.
	EnterFunctionCallArgs(c *FunctionCallArgsContext)

	// EnterIncreaseDecrease is called when entering the increaseDecrease production.
	EnterIncreaseDecrease(c *IncreaseDecreaseContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitDeclarationRemainder is called when exiting the declarationRemainder production.
	ExitDeclarationRemainder(c *DeclarationRemainderContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockItem is called when exiting the blockItem production.
	ExitBlockItem(c *BlockItemContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitElseClause is called when exiting the elseClause production.
	ExitElseClause(c *ElseClauseContext)

	// ExitNonIfStatement is called when exiting the nonIfStatement production.
	ExitNonIfStatement(c *NonIfStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignmentRest is called when exiting the assignmentRest production.
	ExitAssignmentRest(c *AssignmentRestContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitLogicalOrRest is called when exiting the logicalOrRest production.
	ExitLogicalOrRest(c *LogicalOrRestContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitLogicalAndRest is called when exiting the logicalAndRest production.
	ExitLogicalAndRest(c *LogicalAndRestContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitEqualityRest is called when exiting the equalityRest production.
	ExitEqualityRest(c *EqualityRestContext)

	// ExitEqualityOperator is called when exiting the equalityOperator production.
	ExitEqualityOperator(c *EqualityOperatorContext)

	// ExitComparisonExpression is called when exiting the comparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitComparisonRest is called when exiting the comparisonRest production.
	ExitComparisonRest(c *ComparisonRestContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitAdditionExpression is called when exiting the additionExpression production.
	ExitAdditionExpression(c *AdditionExpressionContext)

	// ExitAdditionExpressionRest is called when exiting the additionExpressionRest production.
	ExitAdditionExpressionRest(c *AdditionExpressionRestContext)

	// ExitAddSubtractOperator is called when exiting the addSubtractOperator production.
	ExitAddSubtractOperator(c *AddSubtractOperatorContext)

	// ExitMultiplicationExpression is called when exiting the multiplicationExpression production.
	ExitMultiplicationExpression(c *MultiplicationExpressionContext)

	// ExitMultiplicationExpressionRest is called when exiting the multiplicationExpressionRest production.
	ExitMultiplicationExpressionRest(c *MultiplicationExpressionRestContext)

	// ExitMultDivModOperator is called when exiting the multDivModOperator production.
	ExitMultDivModOperator(c *MultDivModOperatorContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitArrayAccess is called when exiting the arrayAccess production.
	ExitArrayAccess(c *ArrayAccessContext)

	// ExitFunctionCallArgs is called when exiting the functionCallArgs production.
	ExitFunctionCallArgs(c *FunctionCallArgsContext)

	// ExitIncreaseDecrease is called when exiting the increaseDecrease production.
	ExitIncreaseDecrease(c *IncreaseDecreaseContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)
}
