package parser

import (
	"strconv"

	"BigCooker/syntax/ast"
	"github.com/antlr4-go/antlr/v4"
)

func ParseFile(filename string) (*ast.Program, error) {
	// 1. Stream input
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}

	// 2. Create lexer & parser instance 
	lexer := NewBigCLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewBigCParser(tokenStream)

	tree := p.Program() // "program" is the grammar entrypoint

	// Use the constructor function
	builder := NewASTBuilder()
	astRoot := builder.VisitProgram(tree.(*ProgramContext))

	return astRoot.(*ast.Program), nil
}

type ASTBuilder struct {
	*BaseBigCVisitor
}

func NewASTBuilder() *ASTBuilder { // constructor
	return &ASTBuilder{BaseBigCVisitor: &BaseBigCVisitor{}}
}

func (v *ASTBuilder) VisitProgram(ctx *ProgramContext) interface{} {
	// set BaseNode properties
	program := &ast.Program{
		BaseNode: ast.BaseNode{
			Line: 	ctx.GetStart().GetLine(), 
			Column: ctx.GetStart().GetColumn(),
		},
	}

	// visit 
	for _, declCtx := range ctx.AllDeclaration() {
		decl := v.Visit(declCtx).(ast.Declaration)
		program.Declarations = append(program.Declarations, decl)
	}

	return program 
}

func (v *ASTBuilder) VisitDeclaration(ctx *DeclarationContext) interface{} {
	typeName := ctx.Type_().GetText()
	identifier := ctx.Identifier().GetText()

	var typeNode ast.Type = &ast.PrimitiveType{
		BaseType: ast.BaseType{
			BaseNode: ast.BaseNode {
				Line: 	ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
			},
		},
		Name: typeName,
	}

	arrayNotation := ctx.ArrayNotation()
	if (arrayNotation != nil) {
		sizeExpr := v.Visit(arrayNotation.Expression()).(ast.Expression)

		typeNode = &ast.ArrayType{
			BaseType: ast.BaseType{
				BaseNode: ast.BaseNode{
					Line:	arrayNotation.GetStart().GetLine(), 
					Column: arrayNotation.GetStart().GetLine(),
				}, 
			}, 
			ElementType: typeNode, 
			Size: 		 sizeExpr,
		}
	}

	declRemainder := ctx.DeclarationRemainder()
	firstChild := declRemainder.GetChild(0)

	if (firstChild != nil) {
		treeNode, ok := firstChild.(antlr.TerminalNode)
		if (ok && treeNode.GetText() == "(") {
			// this is a function declaration 
			funcDecl := &ast.FunctionDeclaration{
				BaseDeclaration: ast.BaseDeclaration{
					BaseNode: ast.BaseNode{
						Line: 	ctx.GetStart().GetLine(),
						Column: ctx.GetStart().GetColumn(),
					},
				},
				Name: 		identifier, 
				ReturnType: typeNode, 
				Parameters: []ast.Parameter{},
			}

			paramList := declRemainder.ParameterList()
			if (paramList != nil) {
				params := v.Visit(paramList).([]ast.Parameter)
				funcDecl.Parameters = params
			}

			block := declRemainder.Block()
			funcDecl.Body = v.Visit(block).(*ast.Block)

			return funcDecl
		}
	}

	// else this is a variable declaration
	varDecl := &ast.VarDeclaration{
		BaseDeclaration: ast.BaseDeclaration{
			BaseNode: ast.BaseNode{
				Line: 	ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
			},
		},
		Name: 	identifier, 
		Type: 	typeNode, 
	}

	varInit := declRemainder.VariableInitializer()
	if (varInit != nil) {
		exprCtx := varInit.Expression()
		varDecl.Initializer = v.Visit(exprCtx).(ast.Expression)
	}

	return varDecl
}

// ---------- Expressions ----------
func (v* ASTBuilder) VisitArrayAccessExpression(ctx *ArrayAccessContext) interface{} {
	
}

func (v* ASTBuilder) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	id := ctx.Identifier()
	if (id != nil) {
		return v.CreateIdentifier(id.GetSymbol())
	}
	
	constant := ctx.Constant() 
	if (constant != nil) {
		return v.Visit(constant)
	}

	expr := ctx.Expression()
	if (expr != nil) {
		return v.Visit(expr)
	}

	panic("Unhandled primary expression!!!")
}

func (v *ASTBuilder) VisitConstant(ctx *ConstantContext) interface{} {
    intConst := ctx.IntegerConstant()
    if intConst != nil {
        return v.CreateIntegerLiteral(intConst.GetSymbol())
    }
    
    floatConst := ctx.FloatingConstant()
    if floatConst != nil {
        return v.CreateFloatLiteral(floatConst.GetSymbol())
    }
    
    boolConst := ctx.BooleanConstant()
    if boolConst != nil {
        return v.CreateBoolLiteral(boolConst.GetSymbol())
    }
    
    charConst := ctx.CharConstant()
    if charConst != nil {
        return v.CreateCharLiteral(charConst.GetSymbol())
    }
    
    panic("Unknown constant type!!!")
}

// ---------- Terminals Node ---------
func (v *ASTBuilder) CreateIdentifier(token antlr.Token) *ast.Identifier {
	name := token.GetText()

	return &ast.Identifier{
		BaseExpression: ast.BaseExpression{
			BaseNode: ast.BaseNode{
				Line: 	token.GetLine(), 
				Column: token.GetColumn(),
			},
		}, 
		Name: name,
	}
}
func (v *ASTBuilder) CreateIntegerLiteral(token antlr.Token) *ast.IntegerLiteral {
	value, _ := strconv.ParseInt(token.GetText(), 10, 64)

	return &ast.IntegerLiteral{
		BaseExpression: ast.BaseExpression{
			BaseNode: ast.BaseNode{
				Line: 	token.GetLine(), 
				Column: token.GetColumn(),
			},
		}, 
		Value: value,
	}
}

func (v *ASTBuilder) CreateFloatLiteral(token antlr.Token) *ast.FloatLiteral {
	value, _ := strconv.ParseFloat(token.GetText(), 64)

	return &ast.FloatLiteral{
		BaseExpression: ast.BaseExpression{
			BaseNode: ast.BaseNode{
				Line: 	token.GetLine(), 
				Column: token.GetColumn(),
			},
		}, 
		Value: value,
	}
}

func (v *ASTBuilder) CreateBoolLiteral(token antlr.Token) *ast.BoolLiteral {
	value, _ := strconv.ParseBool(token.GetText())

	return &ast.BoolLiteral{
		BaseExpression: ast.BaseExpression{
			BaseNode: ast.BaseNode{
				Line: 	token.GetLine(), 
				Column: token.GetColumn(),
			},
		}, 
		Value: value,
	}
}

func (v *ASTBuilder) CreateCharLiteral(token antlr.Token) *ast.CharLiteral {
	text := token.GetText()
	text = text[1 : len(text)-1] // extract quotes

	var value rune 
	if (len(text) == 1){
		value = rune(text[0])
	} else if text[0] == '\\' {
		// escape character
		switch text[1]{
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
				Line: 	token.GetLine(), 
				Column: token.GetColumn(),
			},
		}, 
		Value: value,
	}
}