package semantic

import (
	"BigCooker/pkg/semantic/table"
	"BigCooker/pkg/syntax/ast"
	"fmt"
	"math"
)

type SemanticAnalyzer struct {
    SymTable     *table.SymbolTable
    errors []string
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
    return &SemanticAnalyzer{
        SymTable:     table.NewSymbolTable(nil, "global"),
        errors: []string{},
    }
}

func (analyzer *SemanticAnalyzer) Error(line int, msg string) {
    analyzer.errors = append(analyzer.errors, fmt.Sprintf("Line %d: %s", line, msg))
}

// 2 pass analyzer:
// First pass collect all symbols
// Second pass check for usage and type checking
// Allows for declarations after function call
func (analyzer *SemanticAnalyzer) Analyze(program *ast.Program) []string {
    for _, decl := range program.Declarations {
        analyzer.collectDeclaration(decl)
    }

    for _, decl := range program.Declarations {
        analyzer.analyzeDeclaration(decl)
    }

    fmt.Println("Symbol Table Dump:")
    analyzer.SymTable.PrintTable()
    return analyzer.errors
}

// First pass to add all symbols to table first
func (analyzer *SemanticAnalyzer) collectDeclaration(declr ast.Declaration) {
    lastLine := math.MaxInt // Global default
    switch d := declr.(type) {
    case *ast.VarDeclaration:
        if analyzer.SymTable.ScopeType != "global" {
            lastLine = d.EndLine 
        }
        analyzer.SymTable.Define(d.Name, table.Symbol{
            Name:      d.Name,
            Type:      d.Type,
            Scope:     table.ScopeInfo{ValidFirstLine: d.Line, ValidLastLine: lastLine}, // Global
            ArraySize: getArraySize(d.Type),
        })
    case *ast.FunctionDeclaration:
        if analyzer.SymTable.ScopeType != "global" {
            lastLine = d.EndLine 
        }
        analyzer.SymTable.Define(d.Name, table.Symbol{
            Name:       d.Name,
            Type:       &ast.PrimitiveType{Name: "function"},
            Scope:      table.ScopeInfo{ValidFirstLine: d.Line, ValidLastLine: lastLine}, // Requires EndLine
            Parameters: d.Parameters,
            ReturnType: d.ReturnType,
        })
    }
}

// Now we analyze -> can declare and use anywhere within scope
func (analyzer *SemanticAnalyzer) analyzeDeclaration(declr ast.Declaration) {
    switch d := declr.(type) {
    case *ast.VarDeclaration:
        if d.Initializer != nil {
            initType := analyzer.checkExpression(d.Initializer)
            if !typesMatch(d.Type, initType) {
                analyzer.Error(d.Line, fmt.Sprintf("type mismatch in initializer: expected %s, got %s", typeString(d.Type), typeString(initType)))
            }
            if val, ok := getConstantValue(d.Initializer); ok {
                sym := analyzer.SymTable.Symbols[d.Name]
                sym.Value = val
                analyzer.SymTable.Symbols[d.Name] = sym
            }
        }
    case *ast.FunctionDeclaration:
        analyzer.SymTable = table.NewSymbolTable(analyzer.SymTable, "function")
        for _, param := range d.Parameters {
            analyzer.SymTable.Define(param.Name, table.Symbol{
                Name:  param.Name,
                Type:  param.Type,
                Scope: table.ScopeInfo{ValidFirstLine: d.Line, ValidLastLine: d.Body.EndLine},
            })
        }
        analyzer.checkBlock(d.Body, d.Body.EndLine)
        analyzer.SymTable = analyzer.SymTable.Parent
    }
}

// --- Visitors ---
func (analyzer *SemanticAnalyzer) visitDeclaration(decl ast.Declaration) {
    switch d := decl.(type) {
    case *ast.VarDeclaration:
        analyzer.checkVarDeclaration(d)
    case *ast.FunctionDeclaration:
        analyzer.checkFunctionDeclaration(d)
    }
}

func (analyzer *SemanticAnalyzer) visitStatement(stmt ast.Statement, blockEndLine int) {
    switch st := stmt.(type) {
    case *ast.Block:
        analyzer.checkBlock(st, blockEndLine)
    case *ast.IfStatement:
        analyzer.checkIfStatement(st, blockEndLine)
    case *ast.WhileStatement:
        analyzer.checkWhileStatement(st, blockEndLine)
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

// --- SemanticAnalyzerntic Checks ---
func (analyzer *SemanticAnalyzer) checkVarDeclaration(varDeclr *ast.VarDeclaration) {
    lastLine := math.MaxInt // Global default
    if analyzer.SymTable.ScopeType != "global" {
        lastLine = varDeclr.EndLine 
    }

    var size int64 = getArraySize(varDeclr.Type)
    if size < 0 && isArray(varDeclr.Type) {
        analyzer.Error(varDeclr.Line, "array size must be a positive constant")
    }
    var sym table.Symbol = table.Symbol{
        Name:      varDeclr.Name,
        Type:      varDeclr.Type,
        Scope:     table.ScopeInfo{ValidFirstLine: 1, ValidLastLine: lastLine}, // Updated later
        ArraySize: size,
    }
    if varDeclr.Initializer != nil {
        var initType ast.Type = analyzer.checkExpression(varDeclr.Initializer)
        if !typesMatch(varDeclr.Type, initType) {
            analyzer.Error(varDeclr.Line, fmt.Sprintf("type mismatch in initializer: expected %s, got %s", typeString(varDeclr.Type), typeString(initType)))
        }
        if val, ok := getConstantValue(varDeclr.Initializer); ok {
            sym.Value = val
        }
    }
    analyzer.SymTable.Define(varDeclr.Name, sym)
}

func (analyzer *SemanticAnalyzer) checkFunctionDeclaration(funcDeclr *ast.FunctionDeclaration) {
    sym := table.Symbol{
        Name:       funcDeclr.Name,
        Type:       &ast.PrimitiveType{Name: "function"},
        Scope:      table.ScopeInfo{ValidFirstLine: 1, ValidLastLine: math.MaxInt}, // No nested functions allowed, so all functions are global
        Parameters: funcDeclr.Parameters,
        ReturnType: funcDeclr.ReturnType,
    }
    analyzer.SymTable.Define(funcDeclr.Name, sym)
    analyzer.SymTable = table.NewSymbolTable(analyzer.SymTable, "function")
    for _, param := range funcDeclr.Parameters {
        analyzer.SymTable.Define(param.Name, table.Symbol{
            Name:  param.Name,
            Type:  param.Type,
            Scope: table.ScopeInfo{ValidFirstLine: funcDeclr.Line, ValidLastLine: funcDeclr.Body.EndLine},
        })
    }
    analyzer.checkBlock(funcDeclr.Body, funcDeclr.Body.EndLine) // TODO: This is block start line
    fmt.Printf("Table in function scope at line %d \n", funcDeclr.Line)
    analyzer.SymTable.PrintTable()
    analyzer.SymTable = analyzer.SymTable.Parent
}

func (analyzer *SemanticAnalyzer) checkBlock(block *ast.Block, blockEndLine int) {
    analyzer.SymTable = table.NewSymbolTable(analyzer.SymTable, "block")
    for _, item := range block.Items {
        switch it := item.(type) {
        case ast.Declaration:
            switch d := it.(type) {
            case *ast.VarDeclaration:
                analyzer.visitDeclaration(d)
                if sym, ok := analyzer.SymTable.Symbols[d.Name]; ok {
                    sym.Scope.ValidLastLine = blockEndLine
                    analyzer.SymTable.Symbols[d.Name] = sym
                }
            case *ast.FunctionDeclaration:
                // Nested function declarations aren't valid in BigC, so error
                analyzer.Error(d.Line, "nested function declarations are not allowed")
            }
        case ast.Statement:
            analyzer.visitStatement(it, blockEndLine)
        }
    }
    fmt.Printf("Table in block scope at line %d \n", block.Line)
    analyzer.SymTable.PrintTable()
    analyzer.SymTable = analyzer.SymTable.Parent
}

func (analyzer *SemanticAnalyzer) checkIfStatement(ifStmt *ast.IfStatement, blockEndLine int) {
    condType := analyzer.checkExpression(ifStmt.Condition)
    if !isBoolType(condType) {
        analyzer.Error(ifStmt.Line, "if condition must be boolean")
    }
    analyzer.checkBlock(ifStmt.ThenBlock, blockEndLine)
    if ifStmt.ElseBlock != nil {
        if elseBlock, ok := ifStmt.ElseBlock.(*ast.Block); ok {
            analyzer.checkBlock(elseBlock, blockEndLine)
        } else if elseIf, ok := ifStmt.ElseBlock.(*ast.IfStatement); ok {
            analyzer.checkIfStatement(elseIf, blockEndLine)
        }
    }
}

func (analyzer *SemanticAnalyzer) checkWhileStatement(w *ast.WhileStatement, blockEndLine int) {
    condType := analyzer.checkExpression(w.Condition)
    if !isBoolType(condType) {
        analyzer.Error(w.Line, "while condition must be boolean")
    }
    analyzer.checkBlock(w.Body, blockEndLine)
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
        }

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

func (analyzer *SemanticAnalyzer) checkArrayAccessExpression(arrAcxessExpr *ast.ArrayAccessExpression) ast.Type {
    var arrayType ast.Type = analyzer.checkExpression(arrAcxessExpr.Array)
    if arr, ok := arrayType.(*ast.ArrayType); ok {
        if lit, ok := arrAcxessExpr.Index.(*ast.IntegerLiteral); ok {
            var size int64 = getArraySize(arr)
            if lit.Value < 0 || (size >= 0 && lit.Value >= size) {
                analyzer.Error(arrAcxessExpr.Line, "index out of bounds")
            }
        }
        return arr.ElementType
    }
    analyzer.Error(arrAcxessExpr.Line, "array access on non-array type")
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
        return fmt.Sprintf("%s[%d]", typeString(a.ElementType), getArraySize(a))
    }
    return "unknown"
}

func getArraySize(astTypeInterface ast.Type) int64 {
    if arr, ok := astTypeInterface.(*ast.ArrayType); ok {
        if lit, ok := arr.Size.(*ast.IntegerLiteral); ok {
            return int64(lit.Value)
        }
        return -1 // Non-constant size
    }
    return 0
}

func getConstantValue(expr ast.Expression) (any, bool) {
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
    scope := s.SymTable
    for scope != nil {
        for _, sym := range scope.Symbols {
            if p, ok := sym.Type.(*ast.PrimitiveType); ok && p.Name == "function" {
                return &sym
            }
        }
        scope = scope.Parent
    }
    return nil
}

