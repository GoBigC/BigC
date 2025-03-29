package main

import (
	"fmt"
	"os"
	"strings"
	"bytes"
	"io"

	"BigCooker/syntax/parser"
	"BigCooker/syntax/ast"
	
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please supply source file")
		os.Exit(1)
	}
	
	// cst always write to artifact/cst.txt
	cst := getCST(os.Args[1])
	cstFormatted := FormatParseTree(cst)
	err := os.WriteFile("artifact/cst.txt", []byte(cstFormatted), 0644)
	if err != nil {
		fmt.Printf("Error writing CST to file: %v\n", err)
	}
	
	program, err := parser.ParseFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}
	
	var astBuffer bytes.Buffer
	astBuffer.WriteString("========== AST ==========\n")
	PrintAST(program, "", &astBuffer)

	// ast always write to artifact/ast.txt
	err = os.WriteFile("artifact/ast.txt", astBuffer.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Error writing AST to file: %v\n", err)
	}
	
	// -- uncomment the following if want to print ast to terminal too
	// fmt.Println("========== AST ==========")
	// PrintAST(program, "", os.Stdout)
}

func getCST(filename string) string {
	// need this function because we dont want to mess with `parser.go` unnecessarily
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		fmt.Printf("Error creating file stream: %v\n", err)
		return ""
	}
	
	lexer := parser.NewBigCLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBigCParser(tokenStream)
	tree := p.Program()
	
	return tree.ToStringTree(p.GetRuleNames(), p)
}

func PrintAST(node ast.Node, indent string, writer io.Writer) {
	switch n := node.(type) {
	case *ast.Program:
		fmt.Fprintf(writer, "%sProgram (Line %d, Col %d) with %d declarations\n",
			indent, n.Line, n.Column, len(n.Declarations))
		for i, decl := range n.Declarations {
			fmt.Fprintf(writer, "%sDeclaration %d:\n", indent+" ", i+1)
			PrintAST(decl, indent+" ", writer)
		}
	case *ast.VarDeclaration:
		fmt.Fprintf(writer, "%sVarDeclaration (Line %d, Col %d): %s of type ",
			indent, n.Line, n.Column, n.Name)
		PrintAST(n.Type, "", writer)
		if n.Initializer != nil {
			fmt.Fprintf(writer, " with initializer:\n")
			PrintAST(n.Initializer, indent+"  ", writer)
		} else {
			fmt.Fprintln(writer)
		}
	case *ast.FunctionDeclaration:
		fmt.Fprintf(writer, "%sFunctionDeclaration (Line %d, Col %d): %s returns ",
			indent, n.Line, n.Column, n.Name)
		PrintAST(n.ReturnType, "", writer)
		fmt.Fprintf(writer, " with %d parameters\n", len(n.Parameters))
		for i, param := range n.Parameters {
			fmt.Fprintf(writer, "%sParameter %d: %s of type ", indent+"  ", i+1, param.Name)
			PrintAST(param.Type, "", writer)
			fmt.Fprintln(writer)
		}
		fmt.Fprintf(writer, "%sBody:\n", indent+"  ")
		PrintAST(n.Body, indent+"  ", writer)
	case *ast.PrimitiveType:
		fmt.Fprintf(writer, "%s", n.Name)
	case *ast.ArrayType:
		fmt.Fprintf(writer, "array of ")
		PrintAST(n.ElementType, "", writer)
		fmt.Fprintf(writer, " with size ")
		PrintAST(n.Size, "", writer)
	case *ast.Block:
		fmt.Fprintf(writer, "%sBlock (Line %d, Col %d) with %d items\n",
			indent, n.Line, n.Column, len(n.Items))
		for i, item := range n.Items {
			fmt.Fprintf(writer, "%sItem %d:\n", indent+"  ", i+1)
			PrintAST(item, indent+"  ", writer)
		}
	case *ast.IfStatement:
		fmt.Fprintf(writer, "%sIfStatement (Line %d, Col %d):\n", indent, n.Line, n.Column)
		fmt.Fprintf(writer, "%sCondition:\n", indent+"  ")
		PrintAST(n.Condition, indent+"  ", writer)
		fmt.Fprintf(writer, "%sThen:\n", indent+"  ")
		PrintAST(n.ThenBlock, indent+"  ", writer)
		if n.ElseBlock != nil {
			fmt.Fprintf(writer, "%sElse:\n", indent+"  ")
			PrintAST(n.ElseBlock, indent+"  ", writer)
		}
	case *ast.WhileStatement:
		fmt.Fprintf(writer, "%sWhileStatement (Line %d, Col %d):\n", indent, n.Line, n.Column)
		fmt.Fprintf(writer, "%sCondition:\n", indent+"  ")
		PrintAST(n.Condition, indent+"  ", writer)
		fmt.Fprintf(writer, "%sBody:\n", indent+"  ")
		PrintAST(n.Body, indent+"  ", writer)
	case *ast.ReturnStatement:
		fmt.Fprintf(writer, "%sReturnStatement (Line %d, Col %d):\n", indent, n.Line, n.Column)
		PrintAST(n.Value, indent+"  ", writer)
	case *ast.ExpressionStatement:
		fmt.Fprintf(writer, "%sExpressionStatement (Line %d, Col %d):\n", indent, n.Line, n.Column)
		PrintAST(n.Expr, indent+"  ", writer)
	case *ast.BinaryExpression:
		fmt.Fprintf(writer, "%sBinaryExpression (Line %d, Col %d): Operator '%s'\n",
			indent, n.Line, n.Column, n.Operator)
		fmt.Fprintf(writer, "%sLeft:\n", indent+"  ")
		PrintAST(n.Left, indent+"  ", writer)
		fmt.Fprintf(writer, "%sRight:\n", indent+"  ")
		PrintAST(n.Right, indent+"  ", writer)
	case *ast.UnaryExpression:
		fmt.Fprintf(writer, "%sUnaryExpression (Line %d, Col %d): Operator '%s'\n",
			indent, n.Line, n.Column, n.Operator)
		PrintAST(n.Operand, indent+"  ", writer)
	case *ast.ArrayAccessExpression:
		fmt.Fprintf(writer, "%sArrayAccessExpression (Line %d, Col %d):\n", indent, n.Line, n.Column)
		fmt.Fprintf(writer, "%sArray:\n", indent+"  ")
		PrintAST(n.Array, indent+"  ", writer)
		fmt.Fprintf(writer, "%sIndex:\n", indent+"  ")
		PrintAST(n.Index, indent+"  ", writer)
	case *ast.FunctionCallExpression:
		fmt.Fprintf(writer, "%sFunctionCallExpression (Line %d, Col %d) with %d arguments:\n",
			indent, n.Line, n.Column, len(n.Arguments))
		fmt.Fprintf(writer, "%sFunction:\n", indent+"  ")
		PrintAST(n.Function, indent+"  ", writer)
		for i, arg := range n.Arguments {
			fmt.Fprintf(writer, "%sArgument %d:\n", indent+"  ", i+1)
			PrintAST(arg, indent+"    ", writer)
		}
	case *ast.Identifier:
		fmt.Fprintf(writer, "%sIdentifier (Line %d, Col %d): %s\n",
			indent, n.Line, n.Column, n.Name)
	case *ast.IntegerLiteral:
		fmt.Fprintf(writer, "%sIntegerLiteral (Line %d, Col %d): %d\n",
			indent, n.Line, n.Column, n.Value)
	case *ast.FloatLiteral:
		fmt.Fprintf(writer, "%sFloatLiteral (Line %d, Col %d): %f\n",
			indent, n.Line, n.Column, n.Value)
	case *ast.BoolLiteral:
		fmt.Fprintf(writer, "%sBoolLiteral (Line %d, Col %d): %t\n",
			indent, n.Line, n.Column, n.Value)
	case *ast.CharLiteral:
		fmt.Fprintf(writer, "%sCharLiteral (Line %d, Col %d): '%c'\n",
			indent, n.Line, n.Column, n.Value)
	default: 
		fmt.Fprintf(writer, "%sUnknown node type: %T\n", indent, node)
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