package codegen

import (
	"BigCooker/pkg/syntax/ast"
)

type ExpressionGenerator struct {
	CodeGen *CodeGenerator
}

func NewExpressionGenerator(cg *CodeGenerator) *ExpressionGenerator {
	return &ExpressionGenerator{
		CodeGen: cg,
	}
}

func (epxrGen *ExpressionGenerator) GenerateExpression(expr ast.Expression) string {
	switch e := expr.(type) {
	case *ast.BinaryExpression:
		return epxrGen.GenerateBinaryExpression(e)
	case *ast.UnaryExpression:
		return epxrGen.GenerateUnaryExpression(e)
	case *ast.ArrayAccessExpression:
		return epxrGen.GenerateArrayAccessExpression(e)
	default:
		return ""
	}
}

func (epxrGen *ExpressionGenerator) GenerateArrayAccessExpression(e *ast.ArrayAccessExpression) string {
	panic("unimplemented")
}

func (epxrGen *ExpressionGenerator) GenerateUnaryExpression(e *ast.UnaryExpression) string {
	panic("unimplemented")
}

func (epxrGen *ExpressionGenerator) GenerateBinaryExpression(expr *ast.BinaryExpression) string {
	switch expr.Operator {
	case "+":
		return epxrGen.GenerateAddition(expr)
	case "-":
		return epxrGen.GenerateSubtraction(expr)
	case "*":
		return epxrGen.GenerateMultiplication(expr)
	case "/":
		return epxrGen.GenerateDivision(expr)
	default:
		return "No case should reach here, as everything should be handled in semantic analysis"
	}
}

func (epxrGen *ExpressionGenerator) GenerateDivision(expr *ast.BinaryExpression) string {
	panic("unimplemented")
}

func (epxrGen *ExpressionGenerator) GenerateMultiplication(expr *ast.BinaryExpression) string {
	panic("unimplemented")
}

func (epxrGen *ExpressionGenerator) GenerateSubtraction(expr *ast.BinaryExpression) string {
	panic("unimplemented")
}

func (epxrGen *ExpressionGenerator) GenerateAddition(expr *ast.BinaryExpression) string {

	/* Procedure:
	   Load left and right operands into registers
	   Check if able to use immediate values
	   If not, use temporary registers
	   Generate assembly code for addition
	   Store result in a register
	   return string of addition and register of result */

	// Only need to check for left, right operands must be same type after semantic analysis
	switch expr.Left.(type) {
	case *ast.IntegerLiteral:
		var leftInt int64 = expr.Left.(*ast.IntegerLiteral).Value
		var rightInt int64 = expr.Right.(*ast.IntegerLiteral).Value

		return epxrGen.GenerateIntAddition(leftInt, rightInt)

	case *ast.FloatLiteral:
		var leftFloat float64 = expr.Left.(*ast.FloatLiteral).Value
		var rightFloat float64 = expr.Right.(*ast.FloatLiteral).Value

		return epxrGen.GenerateFloatAddition(leftFloat, rightFloat)

	case *ast.Identifier:
		var leftName string = expr.Left.(*ast.Identifier).Name
		var rightName string = expr.Right.(*ast.Identifier).Name

		var leftID string = epxrGen.CodeGen.CurrentFunction + leftName
		var rightID string = epxrGen.CodeGen.CurrentFunction + rightName

		leftSym, _ := epxrGen.CodeGen.SymTable.Lookup(leftID)
		rightSym, _ := epxrGen.CodeGen.SymTable.Lookup(rightID)

		switch leftSym.Type.(*ast.PrimitiveType).Name {
		case "int":
			return epxrGen.GenerateIntAddition(leftSym.Value.(int64), rightSym.Value.(int64))
		case "float":
			return epxrGen.GenerateFloatAddition(leftSym.Value.(float64), rightSym.Value.(float64))
		}
	}
	return "No case should reach here, as everything should be handled in semantic analysis"
}

func (epxrGen *ExpressionGenerator) GenerateIntAddition(leftInt int64, rightInt int64) string {

	if isImmediateInt(leftInt) {
		if isImmediateInt(rightInt) {
			// Both are immediate integers
			reg := epxrGen.CodeGen.Registers.GetTmpRegister()
			epxrGen.CodeGen.emit("li %s, %d", reg, leftInt)
			epxrGen.CodeGen.emit("addi %s, %s, %d", reg, reg, rightInt)
			return reg
		} else {
			// Left is immediate, right is not
			// Load right operand into a register
			rightReg := epxrGen.CodeGen.Registers.GetTmpRegister()
			epxrGen.CodeGen.emit("li %s, %d", rightReg, rightInt)
			epxrGen.CodeGen.emit("addi %s, %s, %d", rightReg, rightReg, leftInt)
			return rightReg
		}

	} else {
		if isImmediateInt(rightInt) {
			// Left is not immediate, right is
			// Load left operand into a register
			leftReg := epxrGen.CodeGen.Registers.GetTmpRegister()
			epxrGen.CodeGen.emit("li %s, %d", leftReg, leftInt)
			epxrGen.CodeGen.emit("addi %s, %s, %d", leftReg, leftReg, rightInt)
			return leftReg
		} else {
			// Both are not immediate integers
			// Load both operands into registers
			leftReg := epxrGen.CodeGen.Registers.GetTmpRegister()
			rightReg := epxrGen.CodeGen.Registers.GetTmpRegister()
			epxrGen.CodeGen.emit("li %s, %d", leftReg, leftInt)
			epxrGen.CodeGen.emit("li %s, %d", rightReg, rightInt)
			epxrGen.CodeGen.emit("add %s, %s, %s", leftReg, leftReg, rightReg)
			return leftReg
		}
	}
}

func (epxrGen *ExpressionGenerator) GenerateFloatAddition(leftFloat float64, rightFloat float64) string {
	// Insert float data into the data section
	// Temporary names, will think of better names later
	epxrGen.CodeGen.insertData("double_1", ".double", leftFloat)
	epxrGen.CodeGen.insertData("double_2", ".double", rightFloat)

	// Load float values into registers
	leftReg := epxrGen.CodeGen.Registers.GetFloatTmpRegister()
	rightReg := epxrGen.CodeGen.Registers.GetFloatTmpRegister()
	// Load left float value
	epxrGen.CodeGen.emit("la %s, double_1", leftReg)
	epxrGen.CodeGen.emit("fld %s, 0(%s)", leftReg, leftReg)
	// Load right float value
	epxrGen.CodeGen.emit("la %s, double_2", rightReg)
	epxrGen.CodeGen.emit("fld %s, 0(%s)", rightReg, rightReg)
	// Perform addition
	epxrGen.CodeGen.emit("fadd.d %s, %s, %s", leftReg, leftReg, rightReg)
	// Should the result be stored inside a register or in the data section?
	// If the result is assigned to a variable, it should be stored in the data section
	// Else, it should be stored in a register
	// For now, we will store it in a register

	return leftReg

}

func isImmediateInt(value int64) bool {
	return value >= -2048 && value <= 2047
}
