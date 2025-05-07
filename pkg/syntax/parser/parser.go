package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"BigCooker/pkg/error_formatter"
	"BigCooker/pkg/syntax/ast"

	"github.com/antlr4-go/antlr/v4"
)

func ParseFile(filename string) (*ast.Program, error) {
	// 1. Read file manually
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	lines := strings.Split(string(raw), "\n")

	// 2. Use NewInputStream instead of NewFileStream
	input := antlr.NewInputStream(string(raw))

	// 3. Continue as before
	lexer := NewBigCLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewBigCParser(tokenStream)

	errorHandler := error_formatter.NewSyntaxErrorHandler(lines)

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorHandler)

	p.RemoveErrorListeners()
	p.AddErrorListener(errorHandler)

	tree := p.Program()

	// Check if there were any syntax errors
	if len(errorHandler.Errors) > 0 {
		// Return syntax error without attempting to build AST
		return nil, fmt.Errorf("%s", strings.Join(errorHandler.Errors, "\n"))
	}

	// Only build AST if there are no syntax errors
	builder := NewASTBuilder()
	astRoot := builder.VisitProgram(tree.(*ProgramContext))
	program, ok := astRoot.(*ast.Program)
	if !ok {
		return nil, fmt.Errorf("AST construction failed: expected *ast.Program but got %T", astRoot)
	}

	return program, nil
}

type ASTBuilder struct {
	*BaseBigCVisitor
}

func (v *ASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	// Direct dispatch based on node type
	// if not override this Visit method, the antlr parsetree
	// waits for the method to be called (instead of dispatching)
	// which returns all <nil> :)
	switch ctx := tree.(type) {
	case *DeclarationContext:
		return v.VisitDeclaration(ctx)
	case *ProgramContext:
		return v.VisitProgram(ctx)
	case *ArrayNotationContext:
		return v.VisitArrayNotation(ctx)
	case *TypeContext:
		return v.VisitType(ctx)
	case *DeclarationRemainderContext:
		return v.VisitDeclarationRemainder(ctx)
	case *ParameterListContext:
		return v.VisitParameterList(ctx)
	case *ParameterContext:
		return v.VisitParameter(ctx)
	case *BlockContext:
		return v.VisitBlock(ctx)
	case *BlockItemContext:
		return v.VisitBlockItem(ctx)
	case *StatementContext:
		return v.VisitStatement(ctx)
	case *IfStatementContext:
		return v.VisitIfStatement(ctx)
	case *ElseClauseContext:
		return v.VisitElseClause(ctx)
	case *NonIfStatementContext:
		return v.VisitNonIfStatement(ctx)
	case *WhileStatementContext:
		return v.VisitWhileStatement(ctx)
	case *ReturnStatementContext:
		return v.VisitReturnStatement(ctx)
	case *ExpressionContext:
		return v.VisitExpression(ctx)
	case *AssignmentExpressionContext:
		return v.VisitAssignmentExpression(ctx)
	case *AssignmentRestContext:
		return v.VisitAssignmentRest(ctx)
	case *VariableInitializerContext:
		return v.VisitVariableInitializer(ctx)
	case *LogicalOrExpressionContext:
		return v.VisitLogicalOrExpression(ctx)
	case *LogicalOrRestContext:
		return v.VisitLogicalOrRest(ctx)
	case *LogicalAndExpressionContext:
		return v.VisitLogicalAndExpression(ctx)
	case *LogicalAndRestContext:
		return v.VisitLogicalAndRest(ctx)
	case *EqualityExpressionContext:
		return v.VisitEqualityExpression(ctx)
	case *EqualityRestContext:
		return v.VisitEqualityRest(ctx)
	case *EqualityOperatorContext:
		return v.VisitEqualityOperator(ctx)
	case *ComparisonExpressionContext:
		return v.VisitComparisonExpression(ctx)
	case *ComparisonRestContext:
		return v.VisitComparisonRest(ctx)
	case *ComparisonOperatorContext:
		return v.VisitComparisonOperator(ctx)
	case *AdditionExpressionContext:
		return v.VisitAdditionExpression(ctx)
	case *AdditionExpressionRestContext:
		return v.VisitAdditionExpressionRest(ctx)
	case *AddSubtractOperatorContext:
		return v.VisitAddSubtractOperator(ctx)
	case *MultiplicationExpressionContext:
		return v.VisitMultiplicationExpression(ctx)
	case *MultiplicationExpressionRestContext:
		return v.VisitMultiplicationExpressionRest(ctx)
	case *MultDivModOperatorContext:
		return v.VisitMultDivModOperator(ctx)
	case *UnaryExpressionContext:
		return v.VisitUnaryExpression(ctx)
	case *UnaryOperatorContext:
		return v.VisitUnaryOperator(ctx)
	case *PostfixExpressionContext:
		return v.VisitPostfixExpression(ctx)
	case *ArrayAccessContext:
		return v.VisitArrayAccess(ctx)
	case *FunctionCallArgsContext:
		return v.VisitFunctionCallArgs(ctx)
	case *ArgListContext:
		return v.VisitArgList(ctx)
	case *PrimaryExpressionContext:
		return v.VisitPrimaryExpression(ctx)
	case *ConstantContext:
		return v.VisitConstant(ctx)
	default:
		fmt.Printf("WARNING: No specific visitor for type %T, using base visitor\n", tree)
		result := v.BaseBigCVisitor.Visit(tree)
		fmt.Printf("Base visitor returned: %v for %T\n", result, tree)
		return result
	}
}

func NewASTBuilder() *ASTBuilder {
	// [Important] gotta init this baby :)
	baseVisitor := &BaseBigCVisitor{
		BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
	}

	return &ASTBuilder{BaseBigCVisitor: baseVisitor}
}

func (v *ASTBuilder) VisitProgram(ctx *ProgramContext) interface{} {
	// set BaseNode properties
	program := &ast.Program{
		BaseNode: ast.BaseNode{
			Line:      ctx.GetStart().GetLine(),
			Column:    ctx.GetStart().GetColumn(),
			EndLine:   ctx.GetStop().GetLine(),
			EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
		},
	}

	for i, declCtx := range ctx.AllDeclaration() {
		// visit
		result := v.Visit(declCtx)

		if result == nil {
			fmt.Printf("ERROR: Visit returned nil for declaration %d at Line %d: %s\n",
				i+1, declCtx.GetStart().GetLine(), declCtx.GetText())
			continue
		}
		decl, ok := result.(ast.Declaration)
		if !ok {
			fmt.Printf("ERROR: Expected ast.Declaration but got %T for declaration %d at Line %d: %s\n",
				result, i+1, declCtx.GetStart().GetLine(), declCtx.GetText())
			continue
		}
		program.Declarations = append(program.Declarations, decl)
	}

	return program
}

func (v *ASTBuilder) VisitDeclaration(ctx *DeclarationContext) interface{} {
	typeName := ctx.Type_().GetText()
	identifier := ctx.Identifier().GetText()

	var typeNode ast.Type = &ast.PrimitiveType{
		BaseType: ast.BaseType{
			BaseNode: ast.BaseNode{
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
			},
		},
		Name: typeName,
	}

	arrayNotation := ctx.ArrayNotation()
	if arrayNotation != nil {
		result := v.Visit(arrayNotation.Expression())
		if result == nil {
			fmt.Printf("ERROR: Array size expression returned nil at Line %d\n",
				arrayNotation.GetStart().GetLine())
			return nil
		}

		sizeExpr, ok := result.(ast.Expression)
		if !ok {
			fmt.Printf("ERROR: Expected ast.Expression for array size but got %T at Line %d\n",
				result, arrayNotation.GetStart().GetLine())
			return nil
		}

		typeNode = &ast.ArrayType{
			BaseType: ast.BaseType{
				BaseNode: ast.BaseNode{
					Line:      arrayNotation.GetStart().GetLine(),
					Column:    arrayNotation.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			ElementType: typeNode,
			Size:        sizeExpr,
		}
	}

	declRemainder := ctx.DeclarationRemainder()
	if declRemainder.GetChildCount() == 0 {
		fmt.Printf("ERROR: DeclarationRemainder has no children at Line %d\n",
			declRemainder.GetStart().GetLine())
		return nil
	}

	firstChild := declRemainder.GetChild(0)
	if firstChild == nil {
		fmt.Printf("ERROR: First child of DeclarationRemainder is nil at Line %d\n",
			declRemainder.GetStart().GetLine())
		return nil
	}

	treeNode, ok := firstChild.(antlr.TerminalNode)
	if ok && treeNode.GetText() == "(" {
		// Case: function declaration
		funcDecl := &ast.FunctionDeclaration{
			BaseDeclaration: ast.BaseDeclaration{
				BaseNode: ast.BaseNode{
					Line:      ctx.GetStart().GetLine(),
					Column:    ctx.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Name:       identifier,
			ReturnType: typeNode,
			Parameters: []ast.Parameter{},
		}

		// process parameters (if present)
		paramList := declRemainder.ParameterList()
		if paramList != nil {
			result := v.Visit(paramList)
			if result == nil {
				fmt.Printf("ERROR: ParameterList visit returned nil at Line %d\n",
					paramList.GetStart().GetLine())
				return nil
			}

			params, ok := result.([]ast.Parameter)
			if !ok {
				fmt.Printf("ERROR: Expected []ast.Parameter but got %T at Line %d\n",
					result, paramList.GetStart().GetLine())
				return nil
			}

			funcDecl.Parameters = params
		}

		// process function body
		block := declRemainder.Block()
		if block == nil {
			fmt.Printf("ERROR: Function body (Block) is nil at Line %d\n",
				declRemainder.GetStart().GetLine())
			return nil
		}

		result := v.Visit(block)
		if result == nil {
			fmt.Printf("ERROR: Block visit returned nil at Line %d\n",
				block.GetStart().GetLine())
			return nil
		}

		body, ok := result.(*ast.Block)
		if !ok {
			fmt.Printf("ERROR: Expected *ast.Block but got %T at Line %d\n",
				result, block.GetStart().GetLine())
			return nil
		}

		funcDecl.Body = body
		return funcDecl
	}

	// Case: variable declaration
	varDecl := &ast.VarDeclaration{
		BaseDeclaration: ast.BaseDeclaration{
			BaseNode: ast.BaseNode{
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
			},
		},
		Name: identifier,
		Type: typeNode,
	}

	// initializer (if present)
	varInit := declRemainder.VariableInitializer()
	if varInit != nil {
		exprCtx := varInit.Expression()
		if exprCtx == nil {
			fmt.Printf("ERROR: Expression in VariableInitializer is nil at Line %d\n",
				varInit.GetStart().GetLine())
			return nil
		}

		result := v.Visit(exprCtx)
		if result == nil {
			fmt.Printf("ERROR: Expression visit returned nil at Line %d\n",
				exprCtx.GetStart().GetLine())
			return nil
		}

		expr, ok := result.(ast.Expression)
		if !ok {
			fmt.Printf("ERROR: Expected ast.Expression but got %T at Line %d\n",
				result, exprCtx.GetStart().GetLine())
			return nil
		}

		varDecl.Initializer = expr
	}

	return varDecl
}

func (v *ASTBuilder) VisitParameter(ctx *ParameterContext) interface{} {
	typeName := ctx.Type_().GetText()
	identifier := ctx.Identifier().GetText()

	var typeNode ast.Type = &ast.PrimitiveType{
		BaseType: ast.BaseType{
			BaseNode: ast.BaseNode{
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
			},
		},
		Name: typeName,
	}

	arrayNotation := ctx.ArrayNotation()
	if arrayNotation != nil {
		sizeExpr := v.Visit(arrayNotation.Expression()).(ast.Expression)

		typeNode = &ast.ArrayType{
			BaseType: ast.BaseType{
				BaseNode: ast.BaseNode{
					Line:      arrayNotation.GetStart().GetLine(),
					Column:    arrayNotation.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			ElementType: typeNode,
			Size:        sizeExpr,
		}
	}

	return ast.Parameter{
		BaseNode: ast.BaseNode{
			Line:      ctx.GetStart().GetLine(),
			Column:    ctx.GetStart().GetColumn(),
			EndLine:   ctx.GetStop().GetLine(),
			EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
		},
		Name: identifier,
		Type: typeNode,
	}
}

func (v *ASTBuilder) VisitParameterList(ctx *ParameterListContext) interface{} {
	var params []ast.Parameter

	for _, paramCtx := range ctx.AllParameter() {
		param := v.Visit(paramCtx).(ast.Parameter)
		params = append(params, param)
	}

	return params
}

// ---------- Statements ----------
func (v *ASTBuilder) VisitStatement(ctx *StatementContext) interface{} {
	if ifStmt := ctx.IfStatement(); ifStmt != nil {
		return v.Visit(ifStmt)
	}
	return v.Visit(ctx.NonIfStatement())
}

// If statements
func (v *ASTBuilder) VisitIfStatement(ctx *IfStatementContext) interface{} {
	ifStmt := &ast.IfStatement{
		BaseStatement: ast.BaseStatement{
			BaseNode: ast.BaseNode{
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
			},
		},
		Condition: v.Visit(ctx.Expression()).(ast.Expression),
		ThenBlock: v.Visit(ctx.Block()).(*ast.Block),
	}

	elseClause := ctx.ElseClause()
	if elseClause != nil {
		if elseBlock := elseClause.Block(); elseBlock != nil {
			ifStmt.ElseBlock = v.Visit(elseBlock).(ast.Node)
		} else if elseIf := elseClause.IfStatement(); elseIf != nil {
			ifStmt.ElseBlock = v.Visit(elseIf).(ast.Node)
		}
	}

	return ifStmt
}

// NonIfStatement (while, return)
func (v *ASTBuilder) VisitNonIfStatement(ctx *NonIfStatementContext) interface{} {
	if expr := ctx.Expression(); expr != nil {
		return &ast.ExpressionStatement{
			BaseStatement: ast.BaseStatement{
				BaseNode: ast.BaseNode{
					Line:      expr.GetStart().GetLine(),
					Column:    expr.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Expr: v.Visit(expr).(ast.Expression),
		}
	} else if whileStmt := ctx.WhileStatement(); whileStmt != nil {
		return v.Visit(whileStmt)
	} else if returnStmt := ctx.ReturnStatement(); returnStmt != nil {
		return v.Visit(returnStmt)
	}

	panic("Unknown non-if statement type")
}

func (v *ASTBuilder) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return &ast.WhileStatement{
		BaseStatement: ast.BaseStatement{
			BaseNode: ast.BaseNode{
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
			},
		},
		Condition: v.Visit(ctx.Expression()).(ast.Expression),
		Body:      v.Visit(ctx.Block()).(*ast.Block),
	}
}

func (v *ASTBuilder) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return &ast.ReturnStatement{
		BaseStatement: ast.BaseStatement{
			BaseNode: ast.BaseNode{
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
			},
		},
		Value: v.Visit(ctx.Expression()).(ast.Expression),
	}
}

func (v *ASTBuilder) VisitBlock(ctx *BlockContext) interface{} {
	block := &ast.Block{
		BaseStatement: ast.BaseStatement{
			BaseNode: ast.BaseNode{
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
			},
		},
		Items: []ast.BlockItem{},
	}

	for _, itemCtx := range ctx.AllBlockItem() {
		item := v.Visit(itemCtx).(ast.BlockItem)
		block.Items = append(block.Items, item)
	}

	return block
}

func (v *ASTBuilder) VisitBlockItem(ctx *BlockItemContext) interface{} {
	if declCtx := ctx.Declaration(); declCtx != nil {
		return v.Visit(declCtx)
	}
	return v.Visit(ctx.Statement())
}

// ---------- Expressions ----------
// Helper method to create binary expressions
func (v *ASTBuilder) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.Visit(ctx.AssignmentExpression())
}

func (v *ASTBuilder) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	expr := v.Visit(ctx.LogicalOrExpression()).(ast.Expression)

	rest := ctx.AssignmentRest()
	if rest != nil {
		rightExpr := v.Visit(rest.AssignmentExpression()).(ast.Expression)
		expr = &ast.BinaryExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      rest.GetStart().GetLine(),
					Column:    rest.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Left:     expr,
			Operator: "=",
			Right:    rightExpr,
		}
	}

	return expr
}

func (v *ASTBuilder) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	expr := v.Visit(ctx.LogicalAndExpression()).(ast.Expression)

	for _, rest := range ctx.AllLogicalOrRest() {
		rightExpr := v.Visit(rest.LogicalAndExpression()).(ast.Expression)
		expr = &ast.BinaryExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      rest.GetStart().GetLine(),
					Column:    rest.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Left:     expr,
			Operator: "||",
			Right:    rightExpr,
		}
	}

	return expr
}

func (v *ASTBuilder) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	expr := v.Visit(ctx.EqualityExpression()).(ast.Expression)

	for _, rest := range ctx.AllLogicalAndRest() {
		rightExpr := v.Visit(rest.EqualityExpression()).(ast.Expression)
		expr = &ast.BinaryExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      rest.GetStart().GetLine(),
					Column:    rest.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Left:     expr,
			Operator: "&&",
			Right:    rightExpr,
		}
	}

	return expr
}

func (v *ASTBuilder) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	expr := v.Visit(ctx.ComparisonExpression()).(ast.Expression)

	for _, rest := range ctx.AllEqualityRest() {
		rightExpr := v.Visit(rest.ComparisonExpression()).(ast.Expression)
		operator := rest.EqualityOperator().GetText()
		expr = &ast.BinaryExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      rest.GetStart().GetLine(),
					Column:    rest.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Left:     expr,
			Operator: operator,
			Right:    rightExpr,
		}
	}

	return expr
}

func (v *ASTBuilder) VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{} {
	expr := v.Visit(ctx.AdditionExpression()).(ast.Expression)

	for _, rest := range ctx.AllComparisonRest() {
		rightExpr := v.Visit(rest.AdditionExpression()).(ast.Expression)
		operator := rest.ComparisonOperator().GetText()
		expr = &ast.BinaryExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      rest.GetStart().GetLine(),
					Column:    rest.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Left:     expr,
			Operator: operator,
			Right:    rightExpr,
		}
	}

	return expr
}

func (v *ASTBuilder) VisitAdditionExpression(ctx *AdditionExpressionContext) interface{} {
	expr := v.Visit(ctx.MultiplicationExpression()).(ast.Expression)

	for _, rest := range ctx.AllAdditionExpressionRest() {
		rightExpr := v.Visit(rest.MultiplicationExpression()).(ast.Expression)
		operator := rest.AddSubtractOperator().GetText()
		expr = &ast.BinaryExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      rest.GetStart().GetLine(),
					Column:    rest.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Left:     expr,
			Operator: operator,
			Right:    rightExpr,
		}
	}

	return expr
}

func (v *ASTBuilder) VisitMultiplicationExpression(ctx *MultiplicationExpressionContext) interface{} {
	expr := v.Visit(ctx.UnaryExpression()).(ast.Expression)

	for _, rest := range ctx.AllMultiplicationExpressionRest() {
		rightExpr := v.Visit(rest.UnaryExpression()).(ast.Expression)
		operator := rest.MultDivModOperator().GetText()
		expr = &ast.BinaryExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      rest.GetStart().GetLine(),
					Column:    rest.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Left:     expr,
			Operator: operator,
			Right:    rightExpr,
		}
	}

	return expr
}

func (v *ASTBuilder) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	opCtx := ctx.UnaryOperator()
	if opCtx != nil {
		operand := v.Visit(ctx.UnaryExpression()).(ast.Expression)
		operator := opCtx.GetText()

		return &ast.UnaryExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      opCtx.GetStart().GetLine(),
					Column:    opCtx.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Operator: operator,
			Operand:  operand,
		}
	}

	return v.Visit(ctx.PostfixExpression())
}

// (array access, function calls)
func (v *ASTBuilder) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	primaryExpr := v.Visit(ctx.PrimaryExpression()).(ast.Expression)

	arrAccess := ctx.ArrayAccess()
	if arrAccess != nil {
		indexExpr := v.Visit(arrAccess.Expression()).(ast.Expression)
		return &ast.ArrayAccessExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      arrAccess.GetStart().GetLine(),
					Column:    arrAccess.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Array: primaryExpr,
			Index: indexExpr,
		}
	}

	funcCall := ctx.FunctionCallArgs()
	if funcCall != nil {
		callExpr := &ast.FunctionCallExpression{
			BaseExpression: ast.BaseExpression{
				BaseNode: ast.BaseNode{
					Line:      funcCall.GetStart().GetLine(),
					Column:    funcCall.GetStart().GetColumn(),
					EndLine:   ctx.GetStop().GetLine(),
					EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()) - 1,
				},
			},
			Function:  primaryExpr,
			Arguments: []ast.Expression{},
		}

		// Process arguments if they exist
		argList := funcCall.ArgList()
		if argList != nil {
			for _, argCtx := range argList.AllAssignmentExpression() {
				argExpr := v.Visit(argCtx).(ast.Expression)
				callExpr.Arguments = append(callExpr.Arguments, argExpr)
			}
		}
		return callExpr
	}

	return primaryExpr
}

func (v *ASTBuilder) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	id := ctx.Identifier()
	if id != nil {
		return v.plantIdentifier(id.GetSymbol())
	}

	constant := ctx.Constant()
	if constant != nil {
		return v.Visit(constant)
	}

	expr := ctx.Expression()
	if expr != nil {
		return v.Visit(expr)
	}

	panic("Unhandled primary expression!!!")
}

func (v *ASTBuilder) VisitConstant(ctx *ConstantContext) interface{} {
	intConst := ctx.IntegerConstant()
	if intConst != nil {
		return v.plantIntegerLiteral(intConst.GetSymbol())
	}

	floatConst := ctx.FloatingConstant()
	if floatConst != nil {
		return v.plantFloatLiteral(floatConst.GetSymbol())
	}

	boolConst := ctx.BooleanConstant()
	if boolConst != nil {
		return v.plantBoolLiteral(boolConst.GetSymbol())
	}

	charConst := ctx.CharConstant()
	if charConst != nil {
		return v.plantCharLiteral(charConst.GetSymbol())
	}

	panic("Unknown constant type!!!")
}

// ---------- Terminals Node ---------
func (v *ASTBuilder) plantIdentifier(token antlr.Token) *ast.Identifier {
	name := token.GetText()

	return &ast.Identifier{
		BaseExpression: ast.BaseExpression{
			BaseNode: ast.BaseNode{
				Line:      token.GetLine(),
				Column:    token.GetColumn(),
				EndLine:   token.GetLine(),
				EndColumn: token.GetColumn() + len(token.GetText()) - 1,
			},
		},
		Name: name,
	}
}
func (v *ASTBuilder) plantIntegerLiteral(token antlr.Token) *ast.IntegerLiteral {
	value, _ := strconv.ParseInt(token.GetText(), 10, 64)

	return &ast.IntegerLiteral{
		BaseExpression: ast.BaseExpression{
			BaseNode: ast.BaseNode{
				Line:      token.GetLine(),
				Column:    token.GetColumn(),
				EndLine:   token.GetLine(),
				EndColumn: token.GetColumn() + len(token.GetText()) - 1,
			},
		},
		Value: value,
	}
}

func (v *ASTBuilder) plantFloatLiteral(token antlr.Token) *ast.FloatLiteral {
	value, _ := strconv.ParseFloat(token.GetText(), 64)

	return &ast.FloatLiteral{
		BaseExpression: ast.BaseExpression{
			BaseNode: ast.BaseNode{
				Line:      token.GetLine(),
				Column:    token.GetColumn(),
				EndLine:   token.GetLine(),
				EndColumn: token.GetColumn() + len(token.GetText()) - 1,
			},
		},
		Value: value,
	}
}

func (v *ASTBuilder) plantBoolLiteral(token antlr.Token) *ast.BoolLiteral {
	value, _ := strconv.ParseBool(token.GetText())

	return &ast.BoolLiteral{
		BaseExpression: ast.BaseExpression{
			BaseNode: ast.BaseNode{
				Line:      token.GetLine(),
				Column:    token.GetColumn(),
				EndLine:   token.GetLine(),
				EndColumn: token.GetColumn() + len(token.GetText()) - 1,
			},
		},
		Value: value,
	}
}

func (v *ASTBuilder) plantCharLiteral(token antlr.Token) *ast.CharLiteral {
	text := token.GetText()
	text = text[1 : len(text)-1] // extract quotes

	var value rune
	if len(text) == 1 {
		value = rune(text[0])
	} else if text[0] == '\\' {
		// escape character
		switch text[1] {
		case 'n':
			value = '\n'
		case 't':
			value = '\t'
		case 'r':
			value = '\r'
		case '\\':
			value = '\\'
		case '\'':
			value = '\''
		case '0':
			value = 0
		default:
			value = rune(text[1])
		}
	}

	return &ast.CharLiteral{
		BaseExpression: ast.BaseExpression{
			BaseNode: ast.BaseNode{
				Line:      token.GetLine(),
				Column:    token.GetColumn(),
				EndLine:   token.GetLine(),
				EndColumn: token.GetColumn() + len(token.GetText()) - 1,
			},
		},
		Value: value,
	}
}
