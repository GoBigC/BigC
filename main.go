package main

import (
	"fmt"
	"os"
	"strings"
	

	"BigCooker/syntax/parser"
	"BigCooker/syntax/ast"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please supply source file")
		os.Exit(1)
	}
	program, err := parser.ParseFile(os.Args[1])
	if (err != nil) {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("========== AST ==========")
	PrintAST(program, "")
}

func PrintAST(node ast.Node, indent string){
	switch n := node.(type) {
	case *ast.Program:
		fmt.Printf("%sProgram (Line %d, Col %d) with %d declarations\n",
			indent, n.Line, n.Column, len(n.Declarations))
		for i, decl := range n.Declarations {
			fmt.Printf("%sDeclaration %d:\n", indent+" ", i+1)
			PrintAST(decl, indent+" ")
		}
	case *ast.VarDeclaration:
		fmt.Printf("%sVarDeclaration (Line %d, Col %d): %s of type ",
			indent, n.Line, n.Column, n.Name)
		PrintAST(n.Type, "")
		if n.Initializer != nil {
			fmt.Printf(" with initializer:\n")
			PrintAST(n.Initializer, indent+"  ")
		} else {
			fmt.Println()
		}
	case *ast.FunctionDeclaration:
		fmt.Printf("%sFunctionDeclaration (Line %d, Col %d): %s returns ",
			indent, n.Line, n.Column, n.Name)
		PrintAST(n.ReturnType, "")
		fmt.Printf(" with %d parameters\n", len(n.Parameters))
		for i, param := range n.Parameters {
			fmt.Printf("%sParameter %d: %s of type ", indent+"  ", i+1, param.Name)
			PrintAST(param.Type, "")
			fmt.Println()
		}
		fmt.Printf("%sBody:\n", indent+"  ")
		PrintAST(n.Body, indent+"  ")
	case *ast.PrimitiveType:
		fmt.Printf("%s", n.Name)
	case *ast.ArrayType:
		fmt.Printf("array of ")
		PrintAST(n.ElementType, "")
		fmt.Printf(" with size ")
		PrintAST(n.Size, "")
	case *ast.Block:
		fmt.Printf("%sBlock (Line %d, Col %d) with %d items\n",
			indent, n.Line, n.Column, len(n.Items))
		for i, item := range n.Items {
			fmt.Printf("%sItem %d:\n", indent+"  ", i+1)
			PrintAST(item, indent+"  ")
		}
	case *ast.IfStatement:
		fmt.Printf("%sIfStatement (Line %d, Col %d):\n", indent, n.Line, n.Column)
		fmt.Printf("%sCondition:\n", indent+"  ")
		PrintAST(n.Condition, indent+"  ")
		fmt.Printf("%sThen:\n", indent+"  ")
		PrintAST(n.ThenBlock, indent+"  ")
		if n.ElseBlock != nil {
			fmt.Printf("%sElse:\n", indent+"  ")
			PrintAST(n.ElseBlock, indent+"  ")
		}
	case *ast.WhileStatement:
		fmt.Printf("%sWhileStatement (Line %d, Col %d):\n", indent, n.Line, n.Column)
		fmt.Printf("%sCondition:\n", indent+"  ")
		PrintAST(n.Condition, indent+"  ")
		fmt.Printf("%sBody:\n", indent+"  ")
		PrintAST(n.Body, indent+"  ")
	case *ast.ReturnStatement:
		fmt.Printf("%sReturnStatement (Line %d, Col %d):\n", indent, n.Line, n.Column)
		PrintAST(n.Value, indent+"  ")
	case *ast.ExpressionStatement:
		fmt.Printf("%sExpressionStatement (Line %d, Col %d):\n", indent, n.Line, n.Column)
		PrintAST(n.Expr, indent+"  ")
	case *ast.BinaryExpression:
		fmt.Printf("%sBinaryExpression (Line %d, Col %d): Operator '%s'\n",
			indent, n.Line, n.Column, n.Operator)
		fmt.Printf("%sLeft:\n", indent+"  ")
		PrintAST(n.Left, indent+"  ")
		fmt.Printf("%sRight:\n", indent+"  ")
		PrintAST(n.Right, indent+"  ")
	case *ast.UnaryExpression:
		fmt.Printf("%sUnaryExpression (Line %d, Col %d): Operator '%s'\n",
			indent, n.Line, n.Column, n.Operator)
		PrintAST(n.Operand, indent+"  ")
	case *ast.ArrayAccessExpression:
		fmt.Printf("%sArrayAccessExpression (Line %d, Col %d):\n", indent, n.Line, n.Column)
		fmt.Printf("%sArray:\n", indent+"  ")
		PrintAST(n.Array, indent+"  ")
		fmt.Printf("%sIndex:\n", indent+"  ")
		PrintAST(n.Index, indent+"  ")
	case *ast.FunctionCallExpression:
		fmt.Printf("%sFunctionCallExpression (Line %d, Col %d) with %d arguments:\n",
			indent, n.Line, n.Column, len(n.Arguments))
		fmt.Printf("%sFunction:\n", indent+"  ")
		PrintAST(n.Function, indent+"  ")
		for i, arg := range n.Arguments {
			fmt.Printf("%sArgument %d:\n", indent+"  ", i+1)
			PrintAST(arg, indent+"    ")
		}
	case *ast.Identifier:
		fmt.Printf("%sIdentifier (Line %d, Col %d): %s\n",
			indent, n.Line, n.Column, n.Name)
	case *ast.IntegerLiteral:
		fmt.Printf("%sIntegerLiteral (Line %d, Col %d): %d\n",
			indent, n.Line, n.Column, n.Value)
	case *ast.FloatLiteral:
		fmt.Printf("%sFloatLiteral (Line %d, Col %d): %f\n",
			indent, n.Line, n.Column, n.Value)
	case *ast.BoolLiteral:
		fmt.Printf("%sBoolLiteral (Line %d, Col %d): %t\n",
			indent, n.Line, n.Column, n.Value)
	case *ast.CharLiteral:
		fmt.Printf("%sCharLiteral (Line %d, Col %d): '%c'\n",
			indent, n.Line, n.Column, n.Value)
	default: 
		fmt.Printf("%sUnknown node type: %T\n", indent, node)
	}
}

func FormatParseTree(rawTree string) string {
	rawTree = strings.TrimSpace(rawTree)
	var fmtTree strings.Builder 
	n_indent := 0 
	inParen := false 

	for i := 0; i < len(rawTree); i++ {
		char := rawTree[i]

		switch (char) { 
		case '(': 
			fmtTree.WriteByte(char)
			n_indent++
			inParen = true 

			fmtTree.WriteString("\n")
			fmtTree.WriteString(strings.Repeat("\t", n_indent))

		case ')': 
			n_indent-- 
			if (inParen) {
				fmtTree.WriteString("\n")
				fmtTree.WriteString(strings.Repeat("\t", n_indent))
			}
			fmtTree.WriteByte(char)
			inParen = false 

		case ' ': 
			if (i > 0 && rawTree[i-1]=='('){
				continue
			}

			if (!inParen){
				fmtTree.WriteString("\n")
				fmtTree.WriteString(strings.Repeat("\t", n_indent))
			} else {
				fmtTree.WriteByte(char)
			}
		
		default: 
			fmtTree.WriteByte(char)
		}
	}

	return fmtTree.String()
}