package semantic

import (
	"BigCooker/pkg/semantic/table"
	"BigCooker/pkg/syntax/ast"
	"fmt"
	"math"
)

type SemanticAnalyzer struct {
	SymTable        *table.SymbolTable
	errors          []string
	currentFunction string
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
	return &SemanticAnalyzer{
		SymTable:        table.NewSymbolTable(),
		errors:          []string{},
		currentFunction: "",
	}
}

func (analyzer *SemanticAnalyzer) Error(line int, msg string) {
	analyzer.errors = append(analyzer.errors, fmt.Sprintf("Line %d: %s", line, msg))
}

func (analyzer *SemanticAnalyzer) PrintErrors() {
	for _, err := range analyzer.errors {
		fmt.Println(err)
	}
}

// 3 pass analyzer:
// First pass collect all global symbols
// Second pass, collect local symbols, check for usage and type checking
// Allows for declarations after function call
// Third pass finalizes all array size into constant
func (analyzer *SemanticAnalyzer) Analyze(program *ast.Program) []string {
	for _, decl := range program.Declarations {
		analyzer.collectDeclaration(decl)
	}

	for _, decl := range program.Declarations {
		analyzer.analyzeDeclaration(decl)
	}

	analyzer.finalizeArraySizes()
	
	fmt.Println("Symbol table dump:")
	analyzer.SymTable.PrintTable()

	return analyzer.errors
}

func (analyzer *SemanticAnalyzer) finalizeArraySizes() {
	for name, symbol := range analyzer.SymTable.Symbols {
		if isArray(symbol.Type) {
			arr := symbol.Type.(*ast.ArrayType)
			if lit, ok := arr.Size.(*ast.IntegerLiteral); ok {
				symbol.ArraySize = lit.Value
				analyzer.SymTable.Symbols[name] = symbol 
			} else {
				size, isConst := analyzer.evaluateArraySize(arr.Size) 
				if isConst {
					symbol.ArraySize = size 
					analyzer.SymTable.Symbols[name] = symbol 
				}
			}
		}
	}
}

// First pass to add all symbols to table first
func (analyzer *SemanticAnalyzer) collectDeclaration(declr ast.Declaration) {
	// Global default
	lastLine := math.MaxInt
	switch d := declr.(type) {
	case *ast.VarDeclaration:
		name := d.Name
		if analyzer.currentFunction != "" {
			name = fmt.Sprintf("%s.%s", analyzer.currentFunction, d.Name)
		}

		sym, ok := analyzer.SymTable.Lookup(name)
		if ok {
			analyzer.Error(d.Line, fmt.Sprintf("variable %s already declared at line %d", sym.Name, sym.Scope.ValidFirstLine))
		} else {

			analyzer.SymTable.Define(name, table.Symbol{
				Name:  d.Name,
				Type:  d.Type,
				Scope: table.ScopeInfo{ValidFirstLine: d.Line, ValidLastLine: lastLine},
			})
		}
	case *ast.FunctionDeclaration:
		funcName := d.Name
		if funcName != "main" {
			analyzer.Error(d.Line, "cannot declare function, only main is allowed")
		} else {
			if sym, ok := analyzer.SymTable.Lookup(funcName); ok {
				analyzer.Error(d.Line, fmt.Sprintf("function %s already declared at line %d", sym.Name, sym.Scope.ValidFirstLine))
			} else {

				// All functions are global, no nested functions allowed
				analyzer.SymTable.Define(d.Name, table.Symbol{
					Name:       d.Name,
					Type:       &ast.PrimitiveType{Name: "function"},
					Scope:      table.ScopeInfo{ValidFirstLine: 1, ValidLastLine: lastLine},
					Parameters: d.Parameters,
					ReturnType: d.ReturnType,
				})
			}
		}
	}
}

// Now we analyze -> can declare and use anywhere within scope
func (analyzer *SemanticAnalyzer) analyzeDeclaration(declr ast.Declaration) {
	switch d := declr.(type) {
	case *ast.VarDeclaration:
		name := d.Name
		if d.Initializer != nil {
			if analyzer.currentFunction != "" {
				name = analyzer.currentFunction + "." + d.Name
			}
			initType := analyzer.checkExpression(d.Initializer)
			if !typesMatch(d.Type, initType) {
				analyzer.Error(d.Line, fmt.Sprintf("type mismatch in initializer: expected %s, got %s", typeString(d.Type), typeString(initType)))
			}
			if val, ok := analyzer.evaluateConstantExpression(d.Initializer); ok { 
				sym := analyzer.SymTable.Symbols[name]
				sym.Value = val
				analyzer.SymTable.Symbols[name] = sym
			}
		}
	case *ast.FunctionDeclaration:
		analyzer.currentFunction = d.Name
		for _, param := range d.Parameters {
			paramName := fmt.Sprintf("%s.%s", d.Name, param.Name)
			analyzer.SymTable.Define(paramName, table.Symbol{
				Name:  param.Name,
				Type:  param.Type,
				Scope: table.ScopeInfo{ValidFirstLine: d.Line, ValidLastLine: d.Body.EndLine},
			})
		}
		analyzer.checkBlock(d.Body)
		analyzer.currentFunction = ""
	}
}

// --- Visitors ---
func (analyzer *SemanticAnalyzer) visitDeclaration(decl ast.Declaration, blockEndLine int) {
	switch d := decl.(type) {
	case *ast.VarDeclaration:
		analyzer.checkVarDeclaration(d, blockEndLine)
	case *ast.FunctionDeclaration:
		analyzer.checkFunctionDeclaration(d)
	}
}

func (analyzer *SemanticAnalyzer) visitStatement(stmt ast.Statement) {
	switch st := stmt.(type) {
	case *ast.Block:
		analyzer.checkBlock(st)
	case *ast.IfStatement:
		analyzer.checkIfStatement(st)
	case *ast.WhileStatement:
		analyzer.checkWhileStatement(st)
	case *ast.ReturnStatement:
		analyzer.checkReturnStatement(st)
	case *ast.ExpressionStatement:
		analyzer.checkExpression(st.Expr)
	}
}

func (analyzer *SemanticAnalyzer) checkExpression(expr ast.Expression) ast.Type {
	switch e := expr.(type) {
	case *ast.BinaryExpression:
		return analyzer.checkBinaryExpression(e)
	case *ast.UnaryExpression:
		return analyzer.checkUnaryExpression(e)
	case *ast.ArrayAccessExpression:
		return analyzer.checkArrayAccessExpression(e)
	case *ast.FunctionCallExpression:
		return analyzer.checkFunctionCallExpression(e)
	case *ast.Identifier:
		return analyzer.checkIdentifier(e)
	case *ast.IntegerLiteral:
		var p *ast.PrimitiveType = &ast.PrimitiveType{Name: "int"}
		p.Line = e.Line
		return p
	case *ast.FloatLiteral:
		var p *ast.PrimitiveType = &ast.PrimitiveType{Name: "float"}
		p.Line = e.Line
		return p
	case *ast.BoolLiteral:
		var p *ast.PrimitiveType = &ast.PrimitiveType{Name: "bool"}
		p.Line = e.Line
		return p
	case *ast.CharLiteral:
		var p *ast.PrimitiveType = &ast.PrimitiveType{Name: "char"}
		p.Line = e.Line
		return p
	}
	return nil
}

// --- SemanticAnalyzer Checks ---
func (analyzer *SemanticAnalyzer) checkVarDeclaration(varDeclr *ast.VarDeclaration, blockEndLine int) {
	name := varDeclr.Name
	lastLine := math.MaxInt // Global default

	if analyzer.currentFunction != "" {
		name = analyzer.currentFunction + "." + varDeclr.Name
		lastLine = blockEndLine
	}

	sym, ok := analyzer.SymTable.Lookup(name)
	if ok {
		analyzer.Error(varDeclr.Line, fmt.Sprintf("variable %s already declared at line %d", sym.Name, sym.Scope.ValidFirstLine))
	} else {
		var size int64 = 0
		
		if isArray(varDeclr.Type) {
			arrayType := varDeclr.Type.(*ast.ArrayType)

			if lit, ok := arrayType.Size.(*ast.IntegerLiteral); ok {
				size = lit.Value
			} else {
				constSize, isConst := analyzer.evaluateArraySize(arrayType.Size)

				if !isConst {
					analyzer.Error(varDeclr.Line, "array size must be a compile-time constant expression")
				} else {
					size = constSize 
				}
			}
			if size < 1 { 
				fmt.Printf("array size is %d", size)
				analyzer.Error(varDeclr.Line, "array size must be a positive constant integer")
			}
		}

		var sym table.Symbol = table.Symbol{
			Name:      varDeclr.Name,
			Type:      varDeclr.Type,
			Scope:     table.ScopeInfo{ValidFirstLine: varDeclr.Line, ValidLastLine: lastLine},
			ArraySize: size,
		}

		if varDeclr.Initializer != nil {
			var initType ast.Type = analyzer.checkExpression(varDeclr.Initializer)
			if !typesMatch(varDeclr.Type, initType) {
				analyzer.Error(varDeclr.Line, fmt.Sprintf("type mismatch in initializer: expected %s, got %s", typeString(varDeclr.Type), typeString(initType)))
			}
			if val, ok := analyzer.evaluateConstantExpression(varDeclr.Initializer); ok { // TODO
				sym.Value = val
			}
		}
		analyzer.SymTable.Define(name, sym)
	}
}

func (analyzer *SemanticAnalyzer) checkFunctionDeclaration(funcDeclr *ast.FunctionDeclaration) {
	analyzer.currentFunction = funcDeclr.Name

	for _, param := range funcDeclr.Parameters {
		paramName := fmt.Sprintf("%s.%s", funcDeclr.Name, param.Name)
		analyzer.SymTable.Define(paramName, table.Symbol{
			Name:  param.Name,
			Type:  param.Type,
			Scope: table.ScopeInfo{ValidFirstLine: funcDeclr.Line, ValidLastLine: funcDeclr.Body.EndLine},
		})
		analyzer.checkBlock(funcDeclr.Body)

		// fmt.Printf("Table in function scope at line %d \n", funcDeclr.Line)
		// analyzer.SymTable.PrintTable()
		// analyzer.SymTable = analyzer.Sym Table.Parent
	}
}

func (analyzer *SemanticAnalyzer) checkBlock(block *ast.Block) {
	blockEndLine := block.EndLine
	for _, item := range block.Items {
		switch it := item.(type) {
		case ast.Declaration:
			switch d := it.(type) {
			case *ast.VarDeclaration:
				analyzer.visitDeclaration(d, blockEndLine)
				if sym, ok := analyzer.SymTable.Symbols[analyzer.currentFunction+"."+d.Name]; ok {
					sym.Scope.ValidLastLine = blockEndLine
					analyzer.SymTable.Symbols[analyzer.currentFunction+"."+d.Name] = sym
				}
			case *ast.FunctionDeclaration:
				// Nested function declarations aren't valid in BigC, so error
				analyzer.Error(d.Line, "nested function declarations are not allowed")
			}
		case ast.Statement:
			analyzer.visitStatement(it)
		}
	}
	// fmt.Printf("Table in block scope at line %d \n", block.Line)
	// analyzer.SymTable.PrintTable()
}

func (analyzer *SemanticAnalyzer) checkIfStatement(ifStmt *ast.IfStatement) {
	condType := analyzer.checkExpression(ifStmt.Condition)
	if !isBoolType(condType) {
		analyzer.Error(ifStmt.Line, "if condition must be boolean")
	}
	analyzer.checkBlock(ifStmt.ThenBlock)
	if ifStmt.ElseBlock != nil {
		if elseBlock, ok := ifStmt.ElseBlock.(*ast.Block); ok {
			analyzer.checkBlock(elseBlock)
		} else if elseIf, ok := ifStmt.ElseBlock.(*ast.IfStatement); ok {
			analyzer.checkIfStatement(elseIf)
		}
	}
}

func (analyzer *SemanticAnalyzer) checkWhileStatement(WhileStatement *ast.WhileStatement) {
	condType := analyzer.checkExpression(WhileStatement.Condition)
	if !isBoolType(condType) {
		analyzer.Error(WhileStatement.Line, "while condition must be boolean")
	}
	analyzer.checkBlock(WhileStatement.Body)
}

func (analyzer *SemanticAnalyzer) checkReturnStatement(r *ast.ReturnStatement) {
	retType := analyzer.checkExpression(r.Value)
	fnScope := analyzer.findFunctionScope()
	if fnScope != nil {
		if !typesMatch(fnScope.ReturnType, retType) {
			analyzer.Error(r.Line, fmt.Sprintf("return type mismatch: expected %s, got %s", typeString(fnScope.ReturnType), typeString(retType)))
		}
	}
}

func (analyzer *SemanticAnalyzer) checkBinaryExpression(binaryExpr *ast.BinaryExpression) ast.Type {
	var leftType ast.Type = analyzer.checkExpression(binaryExpr.Left)
	var rightType ast.Type = analyzer.checkExpression(binaryExpr.Right)
	switch binaryExpr.Operator {
	case "+", "-", "*", "/":
		// Check if both operands are numeric types
		if !isNumericType(leftType) || !isNumericType(rightType) {
			analyzer.Error(binaryExpr.Line, fmt.Sprintf("operator %s requires numeric types", binaryExpr.Operator))
		} else {

			// Check if both operands are of the same type for binary operations
			// e.g int + int, float + float
			if !typesMatch(leftType, rightType) {
				analyzer.Error(binaryExpr.Line, fmt.Sprintf("operator %s requires matching numeric types", binaryExpr.Operator))
			}
			if binaryExpr.Operator == "/" {
				if literal, ok := binaryExpr.Right.(*ast.IntegerLiteral); ok && literal.Value == 0 {
					analyzer.Error(binaryExpr.Line, "division by zero")
				}
			}
		}
		return leftType // Result type matches left operand (simplified)

	case "==", "!=", "<", "<=", ">", ">=":
		if !typesMatch(leftType, rightType) {
			analyzer.Error(binaryExpr.Line, fmt.Sprintf("comparison %s requires matching types", binaryExpr.Operator))
		}
		return &ast.PrimitiveType{Name: "bool"}
	case "&&", "||":
		if !isBoolType(leftType) || !isBoolType(rightType) {
			analyzer.Error(binaryExpr.Line, fmt.Sprintf("logical %s requires boolean types", binaryExpr.Operator))
		}
		return &ast.PrimitiveType{Name: "bool"}
	}
	return nil
}

func (analyzer *SemanticAnalyzer) checkUnaryExpression(unaryExpr *ast.UnaryExpression) ast.Type {
	var operandType ast.Type = analyzer.checkExpression(unaryExpr.Operand)
	switch unaryExpr.Operator {
	case "!":
		if !isBoolType(operandType) {
			analyzer.Error(unaryExpr.Line, "logical not requires boolean type")
		}
		return &ast.PrimitiveType{Name: "bool"}
	case "-":
		if !isNumericType(operandType) {
			analyzer.Error(unaryExpr.Line, "unary minus requires numeric type")
		}
		return operandType
	}
	return nil
}

func (analyzer *SemanticAnalyzer) checkArrayAccessExpression(arrAccessExpr *ast.ArrayAccessExpression) ast.Type {
	var arrayType ast.Type = analyzer.checkExpression(arrAccessExpr.Array)

	if arr, ok := arrayType.(*ast.ArrayType); ok {
		var size int64
		var arrayName string 

		if id, ok := arrAccessExpr.Array.(*ast.Identifier); ok {
			arrayName = id.Name 
			var symName string = arrayName 
			if analyzer.currentFunction != "" {
				symName = fmt.Sprintf("main.%s", arrayName) // not global
			}
			if symbol, ok := analyzer.SymTable.Lookup(symName); ok {
				size = symbol.ArraySize
			} else if symbol, ok := analyzer.SymTable.Lookup(arrayName); ok {
				size = symbol.ArraySize
			}
		}

		if lit, ok := arrAccessExpr.Index.(*ast.IntegerLiteral); ok {
			if lit.Value < 0 || (size >= 0 && lit.Value >= size) {
				analyzer.Error(arrAccessExpr.Line, fmt.Sprintf("index out of bounds, cannot access index %d with array size %d", lit.Value, size))
			}
		} else if variable, ok := arrAccessExpr.Index.(*ast.Identifier); ok {
			name := variable.Name
			if analyzer.currentFunction != "" {
				name = fmt.Sprintf("%s.%s", analyzer.currentFunction, variable.Name)
			}
			sym, ok := analyzer.SymTable.Lookup(name)
			if !ok {
				analyzer.Error(arrAccessExpr.Line, fmt.Sprintf("undefined symbol: %s", variable.Name))
			} else if !isIntType(sym.Type) {
				analyzer.Error(arrAccessExpr.Line, fmt.Sprintf("index must be an integer literal or identifier, not %s", typeString(sym.Type)))
			} else if ok && isIntType(sym.Type) {
				if sym.Value != nil {
					if val, ok := sym.Value.(int64); ok {
						if val < 0 || (size >= 0 && val >= size) {
							analyzer.Error(arrAccessExpr.Line, fmt.Sprintf("index out of bounds, cannot access index %d with array size %d", val, size))
						}
					}
				}
			}
		} else {
			analyzer.Error(arrAccessExpr.Line, "index must be an integer literal or identifier")
		}
		return arr.ElementType
	}
	analyzer.Error(arrAccessExpr.Line, "array access on non-array type")
	return nil
}

func (analyzer *SemanticAnalyzer) checkFunctionCallExpression(funcCallExpr *ast.FunctionCallExpression) ast.Type {

	if id, ok := funcCallExpr.Function.(*ast.Identifier); ok {
		sym, ok := analyzer.SymTable.Lookup(id.Name)
		if !ok || sym.Type.(*ast.PrimitiveType).Name != "function" {
			analyzer.Error(funcCallExpr.Line, fmt.Sprintf("invalid function: %s", id.Name))
			return nil
		}
		if len(funcCallExpr.Arguments) != len(sym.Parameters) {
			analyzer.Error(funcCallExpr.Line, fmt.Sprintf("argument count mismatch, expected %d, got %d", len(sym.Parameters), len(funcCallExpr.Arguments)))
			return sym.ReturnType
		}
		for i, arg := range funcCallExpr.Arguments {
			var argType ast.Type = analyzer.checkExpression(arg)
			if !typesMatch(sym.Parameters[i].Type, argType) {
				analyzer.Error(funcCallExpr.Line, fmt.Sprintf("parameter %d type mismatch: expected %s, got %s", i+1, typeString(sym.Parameters[i].Type), typeString(argType)))
			}
		}
		return sym.ReturnType
	}
	analyzer.Error(funcCallExpr.Line, "function call on non-identifier")
	return nil
}

func (analyzer *SemanticAnalyzer) checkIdentifier(id *ast.Identifier) ast.Type {

	// Check if the identifier is a local variable
	if analyzer.currentFunction != "" {
		scopedVarName := fmt.Sprintf("%s.%s", analyzer.currentFunction, id.Name)
		if sym, ok := analyzer.SymTable.Lookup(scopedVarName); ok {
			if id.Line < sym.Scope.ValidFirstLine || id.Line > sym.Scope.ValidLastLine {
				analyzer.Error(id.Line, fmt.Sprintf("variable %s out of scope", id.Name))
				return nil
			}
			return sym.Type
		}
	}

	// Check global scope
	sym, ok := analyzer.SymTable.Lookup(id.Name)
	if !ok {
		analyzer.Error(id.Line, fmt.Sprintf("undefined symbol: %s", id.Name))
		return nil
	}
	if id.Line < sym.Scope.ValidFirstLine || id.Line > sym.Scope.ValidLastLine {
		analyzer.Error(id.Line, fmt.Sprintf("variable %s out of scope", id.Name))
	}
	return sym.Type
}

// --- Utilities ---
func typesMatch(t1, t2 ast.Type) bool {
	if t1 == nil || t2 == nil {
		return false
	}
	p1, ok1 := t1.(*ast.PrimitiveType)
	p2, ok2 := t2.(*ast.PrimitiveType)
	if ok1 && ok2 {
		return p1.Name == p2.Name
	}
	a1, ok1 := t1.(*ast.ArrayType)
	a2, ok2 := t2.(*ast.ArrayType)
	if ok1 && ok2 {
		return typesMatch(a1.ElementType, a2.ElementType)
	}
	return false
}

func typeString(t ast.Type) string {
	if p, ok := t.(*ast.PrimitiveType); ok {
		return p.Name
	}
	if a, ok := t.(*ast.ArrayType); ok {
		return fmt.Sprintf("%s[]", typeString(a.ElementType)) 
	}
	return "unknown"
}

func (analyzer *SemanticAnalyzer) evaluateConstantExpression(expr ast.Expression) (any, bool) {
	switch e := expr.(type) {
	case *ast.IntegerLiteral:
		return e.Value, true
	case *ast.FloatLiteral:
		return e.Value, true
	case *ast.BoolLiteral:
		return e.Value, true
	case *ast.CharLiteral:
		return e.Value, true
	}
	return nil, false
}

func (analyzer *SemanticAnalyzer) evaluateArraySize(expr ast.Expression) (int64, bool) {
	switch e := expr.(type) {
	case *ast.IntegerLiteral: 
		return e.Value, true
	case *ast.BinaryExpression: 
		leftVal, leftOk := analyzer.evaluateArraySize(e.Left)
		rightVal, rightOk := analyzer.evaluateArraySize(e.Right)

		if !leftOk {
			analyzer.Error(e.Line, "Invalid left operand of array size expression")
			return 0, false
		}

		if !rightOk {
			analyzer.Error(e.Line, "Invalid right operand of array size expression")
			return 0, false
		}

		switch e.Operator {
		case "+": 
			return leftVal+rightVal, true
		case "-": 
			return leftVal-rightVal, true
		case "*": 
			return leftVal*rightVal, true
		case "/": 
			if rightVal == 0 {
				analyzer.Error(e.Line, "division by zero in array size")
				return 0, false
			}
			return leftVal/rightVal, true
		default: // TODO: this function is evaluating all kinds of constant operations not just array types --> bad, we want it to only target array size expression
            analyzer.Error(e.Line, fmt.Sprintf("operator %s not supported in constant expressions", e.Operator))
            return 0, false
		}
	case *ast.Identifier:
		symName := e.Name 

		symbol, found := analyzer.SymTable.Lookup("main." + symName) // check local first
		if !found { // then check global
			symbol, found = analyzer.SymTable.Lookup(symName)
		}
		 
		if found {
			if symbol.Value != nil {
				if intVal, ok := symbol.Value.(int64); ok {
					return intVal, true
				}
			}
		}
		analyzer.Error(e.Line, fmt.Sprintf("identifier %s is not a compile-time constant", e.Name))
    	return 0, false
	default: 
		analyzer.Error(getExprLine(expr), "cannot evaluate constant expression (69420 is placeholder int only)")
		return 0, true
	}
	// return 0, false 
}

func getExprLine(expr ast.Expression) int {
	switch e := expr.(type) {
	case *ast.BinaryExpression: 
		return e.Line
	case *ast.IntegerLiteral: 
		return e.Line
	case *ast.Identifier: 
		return e.Line
	default: 
		return 69420
	}
}

func isArray(t ast.Type) bool {
	_, ok := t.(*ast.ArrayType)
	return ok
}

func isNumericType(t ast.Type) bool {
	if p, ok := t.(*ast.PrimitiveType); ok {
		return p.Name == "int" || p.Name == "float" || p.Name == "char"
	}
	return false
}

func isIntType(t ast.Type) bool {
	if p, ok := t.(*ast.PrimitiveType); ok {
		return p.Name == "int"
	}
	return false
}

func isBoolType(t ast.Type) bool {
	if p, ok := t.(*ast.PrimitiveType); ok {
		return p.Name == "bool"
	}
	return false
}

func (s *SemanticAnalyzer) findFunctionScope() *table.Symbol {
	for _, sym := range s.SymTable.Symbols {
		if p, ok := sym.Type.(*ast.PrimitiveType); ok && p.Name == "function" {
			if s.currentFunction == sym.Name {
				return &sym
			}
		}
	}
	return nil
}
