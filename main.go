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
	
	// Original code
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

	err = os.WriteFile("artifact/ast.txt", astBuffer.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Error writing AST to file: %v\n", err)
	}
	
	// Add new tree visualization code
	err = DrawASTTree(program, "artifact/ast_tree.txt")
	if err != nil {
		fmt.Printf("Error drawing AST tree: %v\n", err)
	}
	
	err = DrawCSTTree(os.Args[1], "artifact/cst_tree.txt")
	if err != nil {
		fmt.Printf("Error drawing CST tree: %v\n", err)
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

func DrawASTTree(program *ast.Program, outputFile string) error {
	var buffer bytes.Buffer
	buffer.WriteString("AST Tree:\n")
	drawASTNode(program, "", true, &buffer)
	
	return os.WriteFile(outputFile, buffer.Bytes(), 0644)
}

func drawASTNode(node ast.Node, prefix string, isLast bool, buffer *bytes.Buffer) {
	// Set up the branch indicators
	var connector, childPrefix string
	if isLast {
		connector = "└── "
		childPrefix = prefix + "    "
	} else {
		connector = "├── "
		childPrefix = prefix + "│   "
	}
	
	// Write the current node with its connector
	switch n := node.(type) {
	case *ast.Program:
		fmt.Fprintf(buffer, "%s%sProgram (Line %d, Col %d) with %d declarations\n", 
			prefix, connector, n.Line, n.Column, len(n.Declarations))
		for i, decl := range n.Declarations {
			drawASTNode(decl, childPrefix, i == len(n.Declarations)-1, buffer)
		}
	case *ast.VarDeclaration:
		fmt.Fprintf(buffer, "%s%sVarDeclaration: %s of type ", 
			prefix, connector, n.Name)
		// For types, print inline without tree format
		switch t := n.Type.(type) {
		case *ast.PrimitiveType:
			fmt.Fprintf(buffer, "%s", t.Name)
		case *ast.ArrayType:
			fmt.Fprintf(buffer, "array of ")
			// For simplicity, handle array type inline
			switch et := t.ElementType.(type) {
			case *ast.PrimitiveType:
				fmt.Fprintf(buffer, "%s", et.Name)
			default:
				fmt.Fprintf(buffer, "[complex type]")
			}
			fmt.Fprintf(buffer, " with size ")
			// Print size inline
			if intLit, ok := t.Size.(*ast.IntegerLiteral); ok {
				fmt.Fprintf(buffer, "%d", intLit.Value)
			} else {
				fmt.Fprintf(buffer, "[expression]")
			}
		}
		
		fmt.Fprintln(buffer)
		
		if n.Initializer != nil {
			nextChildPrefix := childPrefix + "│   "
			fmt.Fprintf(buffer, "%s%sInitializer:\n", childPrefix, "├── ")
			drawASTNode(n.Initializer, nextChildPrefix, true, buffer)
		}
	case *ast.FunctionDeclaration:
		fmt.Fprintf(buffer, "%s%sFunction: %s returns ", 
			prefix, connector, n.Name)
		// Print return type inline
		switch t := n.ReturnType.(type) {
		case *ast.PrimitiveType:
			fmt.Fprintf(buffer, "%s", t.Name)
		default:
			fmt.Fprintf(buffer, "[complex type]")
		}
		fmt.Fprintf(buffer, " with %d parameters\n", len(n.Parameters))
		
		// Parameters
		if len(n.Parameters) > 0 {
			paramChildPrefix := childPrefix + "│   "
			fmt.Fprintf(buffer, "%s%sParameters:\n", childPrefix, "├── ")
			for i, param := range n.Parameters {
				isLastParam := i == len(n.Parameters)-1
				var paramConnector string
				if isLastParam {
					paramConnector = "└── "
				} else {
					paramConnector = "├── "
				}
				fmt.Fprintf(buffer, "%s%s%s of type ", 
					paramChildPrefix, paramConnector, param.Name)
				// Print parameter type inline
				switch t := param.Type.(type) {
				case *ast.PrimitiveType:
					fmt.Fprintf(buffer, "%s\n", t.Name)
				case *ast.ArrayType:
					fmt.Fprintf(buffer, "array of ")
					if pt, ok := t.ElementType.(*ast.PrimitiveType); ok {
						fmt.Fprintf(buffer, "%s", pt.Name)
					} else {
						fmt.Fprintf(buffer, "[complex type]")
					}
					fmt.Fprintf(buffer, " with size ")
					if intLit, ok := t.Size.(*ast.IntegerLiteral); ok {
						fmt.Fprintf(buffer, "%d\n", intLit.Value)
					} else {
						fmt.Fprintf(buffer, "[expression]\n")
					}
				default:
					fmt.Fprintf(buffer, "[complex type]\n")
				}
			}
		}
		
		// Body
		fmt.Fprintf(buffer, "%s%sBody:\n", childPrefix, "└── ")
		drawASTNode(n.Body, childPrefix+"    ", true, buffer)
	case *ast.Block:
		fmt.Fprintf(buffer, "%s%sBlock with %d items\n", 
			prefix, connector, len(n.Items))
		for i, item := range n.Items {
			drawASTNode(item, childPrefix, i == len(n.Items)-1, buffer)
		}
	case *ast.IfStatement:
		fmt.Fprintf(buffer, "%s%sIfStatement\n", prefix, connector)
		fmt.Fprintf(buffer, "%s%sCondition:\n", childPrefix, "├── ")
		drawASTNode(n.Condition, childPrefix+"│   ", true, buffer)
		fmt.Fprintf(buffer, "%s%sThenBlock:\n", childPrefix, "├── ")
		drawASTNode(n.ThenBlock, childPrefix+"│   ", n.ElseBlock == nil, buffer)
		if n.ElseBlock != nil {
			fmt.Fprintf(buffer, "%s%sElseBlock:\n", childPrefix, "└── ")
			drawASTNode(n.ElseBlock, childPrefix+"    ", true, buffer)
		}
	case *ast.WhileStatement:
		fmt.Fprintf(buffer, "%s%sWhileStatement\n", prefix, connector)
		fmt.Fprintf(buffer, "%s%sCondition:\n", childPrefix, "├── ")
		drawASTNode(n.Condition, childPrefix+"│   ", true, buffer)
		fmt.Fprintf(buffer, "%s%sBody:\n", childPrefix, "└── ")
		drawASTNode(n.Body, childPrefix+"    ", true, buffer)
	case *ast.ReturnStatement:
		fmt.Fprintf(buffer, "%s%sReturnStatement\n", prefix, connector)
		drawASTNode(n.Value, childPrefix, true, buffer)
	case *ast.ExpressionStatement:
		fmt.Fprintf(buffer, "%s%sExpressionStatement\n", prefix, connector)
		drawASTNode(n.Expr, childPrefix, true, buffer)
	case *ast.BinaryExpression:
		fmt.Fprintf(buffer, "%s%sBinaryExpression: '%s'\n", 
			prefix, connector, n.Operator)
		fmt.Fprintf(buffer, "%s%sLeft:\n", childPrefix, "├── ")
		drawASTNode(n.Left, childPrefix+"│   ", false, buffer)
		fmt.Fprintf(buffer, "%s%sRight:\n", childPrefix, "└── ")
		drawASTNode(n.Right, childPrefix+"    ", true, buffer)
	case *ast.UnaryExpression:
		fmt.Fprintf(buffer, "%s%sUnaryExpression: '%s'\n", 
			prefix, connector, n.Operator)
		drawASTNode(n.Operand, childPrefix, true, buffer)
	case *ast.ArrayAccessExpression:
		fmt.Fprintf(buffer, "%s%sArrayAccess\n", prefix, connector)
		fmt.Fprintf(buffer, "%s%sArray:\n", childPrefix, "├── ")
		drawASTNode(n.Array, childPrefix+"│   ", false, buffer)
		fmt.Fprintf(buffer, "%s%sIndex:\n", childPrefix, "└── ")
		drawASTNode(n.Index, childPrefix+"    ", true, buffer)
	case *ast.FunctionCallExpression:
		fmt.Fprintf(buffer, "%s%sFunctionCall with %d arguments\n", 
			prefix, connector, len(n.Arguments))
		fmt.Fprintf(buffer, "%s%sFunction:\n", childPrefix, "├── ")
		drawASTNode(n.Function, childPrefix+"│   ", len(n.Arguments) == 0, buffer)
		
		if len(n.Arguments) > 0 {
			fmt.Fprintf(buffer, "%s%sArguments:\n", childPrefix, "└── ")
			argPrefix := childPrefix + "    "
			for i, arg := range n.Arguments {
				drawASTNode(arg, argPrefix, i == len(n.Arguments)-1, buffer)
			}
		}
	case *ast.Identifier:
		fmt.Fprintf(buffer, "%s%sIdentifier: %s\n", 
			prefix, connector, n.Name)
	case *ast.IntegerLiteral:
		fmt.Fprintf(buffer, "%s%sIntegerLiteral: %d\n", 
			prefix, connector, n.Value)
	case *ast.FloatLiteral:
		fmt.Fprintf(buffer, "%s%sFloatLiteral: %f\n", 
			prefix, connector, n.Value)
	case *ast.BoolLiteral:
		fmt.Fprintf(buffer, "%s%sBoolLiteral: %t\n", 
			prefix, connector, n.Value)
	case *ast.CharLiteral:
		fmt.Fprintf(buffer, "%s%sCharLiteral: '%c'\n", 
			prefix, connector, n.Value)
	default:
		fmt.Fprintf(buffer, "%s%sUnknown node type: %T\n", 
			prefix, connector, node)
	}
}

func DrawCSTTree(filename string, outputFile string) error {
	cst := getCST(filename)
	
	var buffer bytes.Buffer
	buffer.WriteString("CST Tree:\n")
	formatCSTAsTree(cst, "", true, &buffer)
	
	return os.WriteFile(outputFile, buffer.Bytes(), 0644)
}

func formatCSTAsTree(cst string, prefix string, isLast bool, buffer *bytes.Buffer) {
	if strings.TrimSpace(cst) == "" {
		return
	}
	
	var connector, childPrefix string
	if isLast {
		connector = "└── "
		childPrefix = prefix + "    "
	} else {
		connector = "├── "
		childPrefix = prefix + "│   "
	}
	
	nodeName, children := parseNode(cst)
	
	fmt.Fprintf(buffer, "%s%s%s\n", prefix, connector, nodeName)
	
	for i, child := range children {
		isLastChild := (i == len(children)-1)
		formatCSTAsTree(child, childPrefix, isLastChild, buffer)
	}
}

func parseNode(cst string) (string, []string) {
	cst = strings.TrimSpace(cst)
	
	if !strings.HasPrefix(cst, "(") {
		return cst, []string{}
	}
	
	// Remove outer parentheses
	cst = strings.TrimPrefix(cst, "(")
	cst = strings.TrimSuffix(cst, ")")
	
	// Extract node name (first word)
	var nodeName string
	spaceIndex := strings.Index(cst, " ")
	if spaceIndex == -1 {
		// No children, just the node name
		return cst, []string{}
	}
	
	nodeName = cst[:spaceIndex]
	content := cst[spaceIndex+1:]
	
	// Parse the children
	children := extractChildren(content)
	return nodeName, children
}

// extractChildren extracts the children nodes from CST content
func extractChildren(content string) []string {
	var children []string
	var currentChild strings.Builder
	parenthesesCount := 0
	inChild := false
	
	for i := 0; i < len(content); i++ {
		char := content[i]
		
		if char == '(' {
			parenthesesCount++
			inChild = true
			currentChild.WriteByte(char)
		} else if char == ')' {
			parenthesesCount--
			currentChild.WriteByte(char)
			
			if parenthesesCount == 0 && inChild {
				children = append(children, currentChild.String())
				currentChild.Reset()
				inChild = false
			}
		} else if inChild {
			currentChild.WriteByte(char)
		} else if char != ' ' && !inChild { // Start of a non-parenthesized child
			// Collect the token/terminal until the next whitespace or parenthesis
			start := i
			for i < len(content) && content[i] != ' ' && content[i] != '(' && content[i] != ')' {
				i++
			}
			token := content[start:i]
			if strings.TrimSpace(token) != "" {
				children = append(children, token)
			}
			i-- // Adjust for the loop increment
		}
	}
	
	return children
}

