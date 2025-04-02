// Generated from /home/lanphgphm/Projects/BigCLang/BigC/grammar/BigC.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link BigCParser}.
 */
public interface BigCListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link BigCParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(BigCParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(BigCParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#declaration}.
	 * @param ctx the parse tree
	 */
	void enterDeclaration(BigCParser.DeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#declaration}.
	 * @param ctx the parse tree
	 */
	void exitDeclaration(BigCParser.DeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#arrayNotation}.
	 * @param ctx the parse tree
	 */
	void enterArrayNotation(BigCParser.ArrayNotationContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#arrayNotation}.
	 * @param ctx the parse tree
	 */
	void exitArrayNotation(BigCParser.ArrayNotationContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(BigCParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(BigCParser.TypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#declarationRemainder}.
	 * @param ctx the parse tree
	 */
	void enterDeclarationRemainder(BigCParser.DeclarationRemainderContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#declarationRemainder}.
	 * @param ctx the parse tree
	 */
	void exitDeclarationRemainder(BigCParser.DeclarationRemainderContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#parameterList}.
	 * @param ctx the parse tree
	 */
	void enterParameterList(BigCParser.ParameterListContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#parameterList}.
	 * @param ctx the parse tree
	 */
	void exitParameterList(BigCParser.ParameterListContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#parameter}.
	 * @param ctx the parse tree
	 */
	void enterParameter(BigCParser.ParameterContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#parameter}.
	 * @param ctx the parse tree
	 */
	void exitParameter(BigCParser.ParameterContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(BigCParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(BigCParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#blockItem}.
	 * @param ctx the parse tree
	 */
	void enterBlockItem(BigCParser.BlockItemContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#blockItem}.
	 * @param ctx the parse tree
	 */
	void exitBlockItem(BigCParser.BlockItemContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(BigCParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(BigCParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(BigCParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(BigCParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#elseClause}.
	 * @param ctx the parse tree
	 */
	void enterElseClause(BigCParser.ElseClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#elseClause}.
	 * @param ctx the parse tree
	 */
	void exitElseClause(BigCParser.ElseClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#nonIfStatement}.
	 * @param ctx the parse tree
	 */
	void enterNonIfStatement(BigCParser.NonIfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#nonIfStatement}.
	 * @param ctx the parse tree
	 */
	void exitNonIfStatement(BigCParser.NonIfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#whileStatement}.
	 * @param ctx the parse tree
	 */
	void enterWhileStatement(BigCParser.WhileStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#whileStatement}.
	 * @param ctx the parse tree
	 */
	void exitWhileStatement(BigCParser.WhileStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void enterReturnStatement(BigCParser.ReturnStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void exitReturnStatement(BigCParser.ReturnStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(BigCParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(BigCParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#assignmentExpression}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentExpression(BigCParser.AssignmentExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#assignmentExpression}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentExpression(BigCParser.AssignmentExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#assignmentRest}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentRest(BigCParser.AssignmentRestContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#assignmentRest}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentRest(BigCParser.AssignmentRestContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#variableInitializer}.
	 * @param ctx the parse tree
	 */
	void enterVariableInitializer(BigCParser.VariableInitializerContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#variableInitializer}.
	 * @param ctx the parse tree
	 */
	void exitVariableInitializer(BigCParser.VariableInitializerContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#logicalOrExpression}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOrExpression(BigCParser.LogicalOrExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#logicalOrExpression}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOrExpression(BigCParser.LogicalOrExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#logicalOrRest}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOrRest(BigCParser.LogicalOrRestContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#logicalOrRest}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOrRest(BigCParser.LogicalOrRestContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#logicalAndExpression}.
	 * @param ctx the parse tree
	 */
	void enterLogicalAndExpression(BigCParser.LogicalAndExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#logicalAndExpression}.
	 * @param ctx the parse tree
	 */
	void exitLogicalAndExpression(BigCParser.LogicalAndExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#logicalAndRest}.
	 * @param ctx the parse tree
	 */
	void enterLogicalAndRest(BigCParser.LogicalAndRestContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#logicalAndRest}.
	 * @param ctx the parse tree
	 */
	void exitLogicalAndRest(BigCParser.LogicalAndRestContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#equalityExpression}.
	 * @param ctx the parse tree
	 */
	void enterEqualityExpression(BigCParser.EqualityExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#equalityExpression}.
	 * @param ctx the parse tree
	 */
	void exitEqualityExpression(BigCParser.EqualityExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#equalityRest}.
	 * @param ctx the parse tree
	 */
	void enterEqualityRest(BigCParser.EqualityRestContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#equalityRest}.
	 * @param ctx the parse tree
	 */
	void exitEqualityRest(BigCParser.EqualityRestContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#equalityOperator}.
	 * @param ctx the parse tree
	 */
	void enterEqualityOperator(BigCParser.EqualityOperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#equalityOperator}.
	 * @param ctx the parse tree
	 */
	void exitEqualityOperator(BigCParser.EqualityOperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#comparisonExpression}.
	 * @param ctx the parse tree
	 */
	void enterComparisonExpression(BigCParser.ComparisonExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#comparisonExpression}.
	 * @param ctx the parse tree
	 */
	void exitComparisonExpression(BigCParser.ComparisonExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#comparisonRest}.
	 * @param ctx the parse tree
	 */
	void enterComparisonRest(BigCParser.ComparisonRestContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#comparisonRest}.
	 * @param ctx the parse tree
	 */
	void exitComparisonRest(BigCParser.ComparisonRestContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#comparisonOperator}.
	 * @param ctx the parse tree
	 */
	void enterComparisonOperator(BigCParser.ComparisonOperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#comparisonOperator}.
	 * @param ctx the parse tree
	 */
	void exitComparisonOperator(BigCParser.ComparisonOperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#additionExpression}.
	 * @param ctx the parse tree
	 */
	void enterAdditionExpression(BigCParser.AdditionExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#additionExpression}.
	 * @param ctx the parse tree
	 */
	void exitAdditionExpression(BigCParser.AdditionExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#additionExpressionRest}.
	 * @param ctx the parse tree
	 */
	void enterAdditionExpressionRest(BigCParser.AdditionExpressionRestContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#additionExpressionRest}.
	 * @param ctx the parse tree
	 */
	void exitAdditionExpressionRest(BigCParser.AdditionExpressionRestContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#addSubtractOperator}.
	 * @param ctx the parse tree
	 */
	void enterAddSubtractOperator(BigCParser.AddSubtractOperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#addSubtractOperator}.
	 * @param ctx the parse tree
	 */
	void exitAddSubtractOperator(BigCParser.AddSubtractOperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#multiplicationExpression}.
	 * @param ctx the parse tree
	 */
	void enterMultiplicationExpression(BigCParser.MultiplicationExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#multiplicationExpression}.
	 * @param ctx the parse tree
	 */
	void exitMultiplicationExpression(BigCParser.MultiplicationExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#multiplicationExpressionRest}.
	 * @param ctx the parse tree
	 */
	void enterMultiplicationExpressionRest(BigCParser.MultiplicationExpressionRestContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#multiplicationExpressionRest}.
	 * @param ctx the parse tree
	 */
	void exitMultiplicationExpressionRest(BigCParser.MultiplicationExpressionRestContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#multDivModOperator}.
	 * @param ctx the parse tree
	 */
	void enterMultDivModOperator(BigCParser.MultDivModOperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#multDivModOperator}.
	 * @param ctx the parse tree
	 */
	void exitMultDivModOperator(BigCParser.MultDivModOperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#unaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterUnaryExpression(BigCParser.UnaryExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#unaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitUnaryExpression(BigCParser.UnaryExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#unaryOperator}.
	 * @param ctx the parse tree
	 */
	void enterUnaryOperator(BigCParser.UnaryOperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#unaryOperator}.
	 * @param ctx the parse tree
	 */
	void exitUnaryOperator(BigCParser.UnaryOperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#postfixExpression}.
	 * @param ctx the parse tree
	 */
	void enterPostfixExpression(BigCParser.PostfixExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#postfixExpression}.
	 * @param ctx the parse tree
	 */
	void exitPostfixExpression(BigCParser.PostfixExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#arrayAccess}.
	 * @param ctx the parse tree
	 */
	void enterArrayAccess(BigCParser.ArrayAccessContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#arrayAccess}.
	 * @param ctx the parse tree
	 */
	void exitArrayAccess(BigCParser.ArrayAccessContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#functionCallArgs}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCallArgs(BigCParser.FunctionCallArgsContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#functionCallArgs}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCallArgs(BigCParser.FunctionCallArgsContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#argList}.
	 * @param ctx the parse tree
	 */
	void enterArgList(BigCParser.ArgListContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#argList}.
	 * @param ctx the parse tree
	 */
	void exitArgList(BigCParser.ArgListContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpression(BigCParser.PrimaryExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpression(BigCParser.PrimaryExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BigCParser#constant}.
	 * @param ctx the parse tree
	 */
	void enterConstant(BigCParser.ConstantContext ctx);
	/**
	 * Exit a parse tree produced by {@link BigCParser#constant}.
	 * @param ctx the parse tree
	 */
	void exitConstant(BigCParser.ConstantContext ctx);
}