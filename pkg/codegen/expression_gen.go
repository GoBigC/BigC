package codegen

import (
	"BigCooker/pkg/syntax/ast"
	"fmt"
	"strings"
)

type ExpressionGenerator struct {
	CodeGen *CodeGenerator
}

func NewExpressionGenerator(cg *CodeGenerator) *ExpressionGenerator {
	return &ExpressionGenerator{
		CodeGen: cg,
	}
}

func (eg *ExpressionGenerator) GenerateExpression(expr ast.Expression) (string, ast.Type) {
	switch e := expr.(type) {
	case *ast.BinaryExpression:
		return eg.GenerateBinaryExpression(e)
	case *ast.UnaryExpression:
		return eg.GenerateUnaryExpression(e)
	case *ast.ArrayAccessExpression:
		return eg.GenerateArrayAccessExpression(e)
	case *ast.FunctionCallExpression:
		return eg.GenerateFunctionCallExpression(e)
	case *ast.IntegerLiteral:
		return eg.GenerateIntegerLiteral(e)
	case *ast.FloatLiteral:
		return eg.GenerateFloatLiteral(e)
	case *ast.BoolLiteral:
		return eg.GenerateBoolLiteral(e)
	case *ast.CharLiteral:
		return eg.GenerateCharLiteral(e)
	case *ast.Identifier:
		return eg.GenerateIdentifier(e)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
}

func (eg *ExpressionGenerator) GenerateIdentifier(expr *ast.Identifier) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	addressRegister := rp.GetTmpRegister()
	cg.emit("    la %s, %s", addressRegister, expr.Name)

	isFloatVar := false
	symbol, found := cg.SymTable.Lookup(expr.Name)
	if !found {
		symbol, found = cg.SymTable.Lookup("main." + expr.Name)
	}

	if found {
		if primitiveType, ok := symbol.Type.(*ast.PrimitiveType); ok {
			isFloatVar = primitiveType.Name == "float"
		}
	}

	if isFloatVar {
		valueRegister := rp.GetFloatTmpRegister()
		cg.emit("    fld %s, 0(%s)", valueRegister, addressRegister)
		rp.ReleaseRegister(addressRegister)
		return valueRegister, symbol.Type
	} else {
		valueRegister := rp.GetTmpRegister()
		cg.emit("    ld %s, 0(%s)", valueRegister, addressRegister)
		rp.ReleaseRegister(addressRegister)
		return valueRegister, symbol.Type
	}
}

func (eg *ExpressionGenerator) GenerateBoolLiteral(expr *ast.BoolLiteral) (string, ast.Type) {
	return "unimplemented", &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateCharLiteral(expr *ast.CharLiteral) (string, ast.Type) {
	return "unimplemented", &ast.PrimitiveType{Name: "char"}
}

func (eg *ExpressionGenerator) GenerateIntegerLiteral(expr *ast.IntegerLiteral) (string, ast.Type) {
	// TODO: this is not handling the case where the integer is
	// greater than the  immediate range -- need to cover that too
	reg := eg.CodeGen.Registers.GetTmpRegister()
	eg.CodeGen.emit("    li %s, %d", reg, expr.Value)
	return reg, &ast.PrimitiveType{Name: "int"}
}

func (eg *ExpressionGenerator) GenerateFloatLiteral(expr *ast.FloatLiteral) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers
	label := fmt.Sprintf("float_imm_%d", cg.Labels)
	cg.Labels++
	cg.insertData(label, ".double", expr.Value)

	addressRegister := rp.GetTmpRegister()
	valueRegister := rp.GetFloatTmpRegister()

	cg.emit("	la %s, %s", addressRegister, label)
	cg.emit("	fld %s, 0(%s)", valueRegister, addressRegister)

	rp.ReleaseRegister(addressRegister)
	return valueRegister, &ast.PrimitiveType{Name: "float"}
}

func (eg *ExpressionGenerator) GenerateFunctionCallExpression(expr *ast.FunctionCallExpression) (string, ast.Type) {
	var funcName string
	cg := eg.CodeGen
	rp := cg.Registers

	if id, ok := expr.Function.(*ast.Identifier); ok {
		funcName = id.Name
	} else {
		panic("Function expression not supported")
	}

	if len(expr.Arguments) > 0 {
		argRegister, _ := eg.GenerateExpression(expr.Arguments[0])

		switch funcName {
		case "_printFloat":
			if argRegister != "fa0" && strings.HasPrefix(argRegister, "f") {
				cg.emit("    fmv.d fa0, %s", argRegister)
			}
		case "_printInt", "_printChar", "_printBool":
			if argRegister != "a0" {
				cg.emit("    mv a0, %s", argRegister)
			}
		case "_printString":
			if argRegister != "a0" {
				cg.emit("    mv a0, %s", argRegister)
			}
		}

		if argRegister != "a0" && argRegister != "fa0" {
			rp.ReleaseRegister(argRegister)
		}
	}

	cg.emit("    jal %s", funcName)

	if funcName == "_printFloat" {
		return "fa0", nil
	}
	return "a0", nil
}

func (eg *ExpressionGenerator) GenerateArrayAccessExpression(e *ast.ArrayAccessExpression) (string, ast.Type) {
	panic("unimplemented")
}

func (eg *ExpressionGenerator) GenerateUnaryExpression(e *ast.UnaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	var resultReg string

	switch e.Operator {
	case "!":
		operandReg, _ := eg.GenerateExpression(e.Operand)
		resultReg = rp.GetTmpRegister()

		cg.emit("    seqz %s, %s", resultReg, operandReg)

		if operandReg != "a0" && operandReg != "fa0" {
			rp.ReleaseRegister(operandReg)
		}

		return resultReg, &ast.PrimitiveType{Name: "bool"}

	case "-":
		operandReg, operandType := eg.GenerateExpression(e.Operand)

		if primitiveType, ok := operandType.(*ast.PrimitiveType); ok {
			switch primitiveType.Name {
			case "int":
				resultReg = rp.GetTmpRegister()
				cg.emit("    neg %s, %s", resultReg, operandReg)
			case "float":
				resultReg = rp.GetFloatTmpRegister()
				cg.emit("    fneg.d %s, %s", resultReg, operandReg)
			default:
				panic(fmt.Sprintf("Unsupported type for unary minus: %s", primitiveType.Name))
			}
		}
		if operandReg != "a0" && operandReg != "fa0" {
			rp.ReleaseRegister(operandReg)
		}
		return resultReg, operandType
	default:
		panic(fmt.Sprintf("Unsupported unary operator: %s", e.Operator))
	}
}

func (eg *ExpressionGenerator) GenerateBinaryExpression(expr *ast.BinaryExpression) (string, ast.Type) {
	switch expr.Operator {
	case "+":
		return eg.GenerateAddition(expr)
	case "-":
		return eg.GenerateSubtraction(expr)
	case "*":
		return eg.GenerateMultiplication(expr)
	case "/":
		return eg.GenerateDivision(expr)
	case "==":
		return eg.GenerateEquality(expr)
	case "!=":
		return eg.GenerateInequality(expr)
	case "<":
		return eg.GenerateLessThan(expr)
	case "<=":
		return eg.GenerateLessThanOrEqual(expr)
	case ">":
		return eg.GenerateGreaterThan(expr)
	case ">=":
		return eg.GenerateGreaterThanOrEqual(expr)
	// case "&&":
	// 	return eg.GenerateLogicalAnd(expr)
	// case "||":
	// 	return eg.GenerateLogicalOr(expr)
	default:
		return "No case should reach here, as everything should be handled in semantic analysis", nil
	}
}

func (eg *ExpressionGenerator) GenerateDivision(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, leftType := eg.GenerateExpression(expr.Left)
	rightReg, rightType := eg.GenerateExpression(expr.Right)

	var resultReg string

	resultType := determineResultType(leftType, rightType)

	switch resultType.(*ast.PrimitiveType).Name {
	case "int":
		resultReg = rp.GetTmpRegister()
		cg.emit("	div %s, %s, %s", resultReg, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "int"}
	case "float":
		resultReg = rp.GetFloatTmpRegister()
		cg.emit("	fdiv.d %s, %s, %s", resultReg, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "float"}
	}

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}

	return "GenerateDivision - No case should reach here, as everything should be handled in semantic analysis", nil

}

func (eg *ExpressionGenerator) GenerateMultiplication(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, leftType := eg.GenerateExpression(expr.Left)
	rightReg, rightType := eg.GenerateExpression(expr.Right)

	var resultReg string

	resultType := determineResultType(leftType, rightType)

	switch resultType.(*ast.PrimitiveType).Name {
	case "int":
		resultReg = rp.GetTmpRegister()
		cg.emit("	mul %s, %s, %s", resultReg, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "int"}
	case "float":
		resultReg = rp.GetFloatTmpRegister()
		cg.emit("	fmul.d %s, %s, %s", resultReg, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "float"}
	}

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}
	return "GenerateMultiplication - No case should reach here, as everything should be handled in semantic analysis", nil
}

func (eg *ExpressionGenerator) GenerateSubtraction(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, leftType := eg.GenerateExpression(expr.Left)
	rightReg, rightType := eg.GenerateExpression(expr.Right)

	var resultReg string

	resultType := determineResultType(leftType, rightType)

	switch resultType.(*ast.PrimitiveType).Name {
	case "int":
		resultReg = rp.GetTmpRegister()
		cg.emit("	sub %s, %s, %s", resultReg, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "int"}
	case "float":
		resultReg = rp.GetFloatTmpRegister()
		cg.emit("	fsub.d %s, %s, %s", resultReg, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "float"}
	}

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}
	return "GenerateSubtraction - No case should reach here, as everything should be handled in semantic analysis", nil
}

func (eg *ExpressionGenerator) GenerateAddition(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, leftType := eg.GenerateExpression(expr.Left)
	rightReg, rightType := eg.GenerateExpression(expr.Right)

	var resultReg string

	resultType := determineResultType(leftType, rightType)

	switch resultType.(*ast.PrimitiveType).Name {
	case "int":
		resultReg = rp.GetTmpRegister()
		cg.emit("	add %s, %s, %s", resultReg, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "int"}
	case "float":
		resultReg = rp.GetFloatTmpRegister()
		cg.emit("	fadd.d %s, %s, %s", resultReg, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "float"}
	}

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}

	return "GenerateAddition - No case should reach here, as everything should be handled in semantic analysis", nil
}

func (eg *ExpressionGenerator) GenerateGreaterThan(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()
	cg.emit("    slt %s, %s, %s", resultReg, rightReg, leftReg)

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}

	return resultReg, &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateLessThan(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()

	cg.emit("    slt %s, %s, %s", resultReg, leftReg, rightReg)

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}

	return resultReg, &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateGreaterThanOrEqual(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()
	tempReg := rp.GetTmpRegister()

	cg.emit("    slt %s, %s, %s", tempReg, leftReg, rightReg)
	cg.emit("    xori %s, %s, 1", resultReg, tempReg) // Invert the result

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}
	rp.ReleaseRegister(tempReg)

	return resultReg, &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateLessThanOrEqual(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()
	tempReg := rp.GetTmpRegister()

	cg.emit("    slt %s, %s, %s", tempReg, rightReg, leftReg)
	cg.emit("    xori %s, %s, 1", resultReg, tempReg) // Invert the result

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}
	rp.ReleaseRegister(tempReg)

	return resultReg, &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateEquality(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()

	cg.emit("    sub %s, %s, %s", resultReg, leftReg, rightReg)
	cg.emit("    seqz %s, %s", resultReg, resultReg) // Set to 1 if equal to zero

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}

	return resultReg, &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateInequality(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()

	cg.emit("    sub %s, %s, %s", resultReg, leftReg, rightReg)
	cg.emit("    snez %s, %s", resultReg, resultReg) // Set to 1 if not equal to zero

	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}

	return resultReg, &ast.PrimitiveType{Name: "bool"}
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

func determineResultType(left, right ast.Type) ast.Type {
	if left == nil || right == nil {
		return nil
	}
	leftPrim, leftIsPrim := left.(*ast.PrimitiveType)
	rightPrim, rightIsPrim := right.(*ast.PrimitiveType)
	if !leftIsPrim || !rightIsPrim {
		return nil
	}
	if leftPrim.Name == "float" || rightPrim.Name == "float" {
		return &ast.PrimitiveType{Name: "float"}
	}
	if leftPrim.Name == "int" && rightPrim.Name == "int" {
		return &ast.PrimitiveType{Name: "int"}
	}
	return nil
}
