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

func (eg *ExpressionGenerator) GenerateExpression(expr ast.Expression) string {
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

func (eg *ExpressionGenerator) GenerateIdentifier(expr *ast.Identifier) string {
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
        return valueRegister
    } else {
        valueRegister := rp.GetTmpRegister()
        cg.emit("    ld %s, 0(%s)", valueRegister, addressRegister)
        rp.ReleaseRegister(addressRegister)
        return valueRegister
    }
}

func (eg *ExpressionGenerator) GenerateBoolLiteral(expr *ast.BoolLiteral) string {
	return "unimplemented"
}

func (eg *ExpressionGenerator) GenerateCharLiteral(expr *ast.CharLiteral) string {
	return "unimplemented"
}

func (eg *ExpressionGenerator) GenerateIntegerLiteral(expr *ast.IntegerLiteral) string {
	// TODO: this is not handling the case where the integer is 
	// greater than the  immediate range -- need to cover that too
    reg := eg.CodeGen.Registers.GetTmpRegister()
    eg.CodeGen.emit("    li %s, %d", reg, expr.Value)
    return reg
}

func (eg *ExpressionGenerator) GenerateFloatLiteral(expr *ast.FloatLiteral) string {
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
	return valueRegister
}

func (eg *ExpressionGenerator) GenerateFunctionCallExpression(expr *ast.FunctionCallExpression) string {
    var funcName string
    cg := eg.CodeGen
    rp := cg.Registers

    if id, ok := expr.Function.(*ast.Identifier); ok {
        funcName = id.Name
    } else {
        panic("Function expression not supported")
    }
    
    if len(expr.Arguments) > 0 {
        argRegister := eg.GenerateExpression(expr.Arguments[0])
        
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
        return "fa0"
    }
    return "a0"
}

func (eg *ExpressionGenerator) GenerateArrayAccessExpression(e *ast.ArrayAccessExpression) string {
	panic("unimplemented")
}

func (eg *ExpressionGenerator) GenerateUnaryExpression(e *ast.UnaryExpression) string {
	cg := eg.CodeGen
    rp := cg.Registers
    
    switch e.Operator {
    case "!":
        operandReg := eg.GenerateExpression(e.Operand)
        resultReg := rp.GetTmpRegister()
        
        cg.emit("    seqz %s, %s", resultReg, operandReg)
        
        if operandReg != "a0" && operandReg != "fa0" {
            rp.ReleaseRegister(operandReg)
        }
        
        return resultReg
        
    case "-":
        operandReg := eg.GenerateExpression(e.Operand)
        resultReg := rp.GetTmpRegister()
        
        cg.emit("    neg %s, %s", resultReg, operandReg)
        
        if operandReg != "a0" && operandReg != "fa0" {
            rp.ReleaseRegister(operandReg)
        }
        
        return resultReg
        
    default:
        panic(fmt.Sprintf("Unsupported unary operator: %s", e.Operator))
    }
}

func (eg *ExpressionGenerator) GenerateBinaryExpression(expr *ast.BinaryExpression) string {
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
		return "No case should reach here, as everything should be handled in semantic analysis"
	}
}

func (eg *ExpressionGenerator) GenerateDivision(expr *ast.BinaryExpression) string {
	switch expr.Left.(type) {
	case *ast.IntegerLiteral:
		var leftInt int64 = expr.Left.(*ast.IntegerLiteral).Value
		var rightInt int64 = expr.Right.(*ast.IntegerLiteral).Value

		return eg.GenerateIntDivision(leftInt, rightInt)

	case *ast.FloatLiteral:
		var leftFloat float64 = expr.Left.(*ast.FloatLiteral).Value
		var rightFloat float64 = expr.Right.(*ast.FloatLiteral).Value

		return eg.GenerateFloatDivision(leftFloat, rightFloat)

	case *ast.Identifier:
		cg := eg.CodeGen
		var leftName string = expr.Left.(*ast.Identifier).Name
		var rightName string = expr.Right.(*ast.Identifier).Name

		var leftID string = leftName
		var rightID string = rightName

		leftSym, _ := cg.SymTable.Lookup(leftID)
		rightSym, _ := cg.SymTable.Lookup(rightID)

		switch leftSym.Type.(*ast.PrimitiveType).Name {
		case "int":
			return eg.GenerateIntDivision(leftSym.Value.(int64), rightSym.Value.(int64))
		case "float":
			return eg.GenerateFloatDivision(leftSym.Value.(float64), rightSym.Value.(float64))
		}

	}
	return "No case should reach here, as everything should be handled in semantic analysis"

}

func (eg *ExpressionGenerator) GenerateIntDivision(leftInt int64, rightInt int64) string {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg := rp.GetTmpRegister()
	rightReg := rp.GetTmpRegister()

	cg.emit("	li %s, %d", leftReg, leftInt)
	cg.emit("	li %s, %d", rightReg, rightInt)
	cg.emit("	div a0, %s, %s", leftReg, rightReg)
	// The result will be a 128 bit integer, but for now we will just return the lower 64 bits
	// Meaning we will ignore overflow, very C-like

	return "a0"
}

func (eg *ExpressionGenerator) GenerateFloatDivision(leftFloat float64, rightFloat float64) string {
	cg := eg.CodeGen
	rp := cg.Registers

	cg.insertData("double_1", ".double", leftFloat)
	cg.insertData("double_2", ".double", rightFloat)

	leftAddressReg := rp.GetTmpRegister()
	rightAddressReg := rp.GetTmpRegister()
	// Load float values into registers
	leftReg := rp.GetFloatTmpRegister()
	rightReg := rp.GetFloatTmpRegister()
	// Load left float value
	cg.emit("	la %s, double_1", leftAddressReg)
	cg.emit("	fld %s, 0(%s)", leftReg, leftAddressReg)
	// Load right float value
	cg.emit("	la %s, double_2", rightAddressReg)
	cg.emit("	fld %s, 0(%s)", rightReg, rightAddressReg)
	// Perform subtraction
	cg.emit("	fdiv.d fa0, %s, %s", leftReg, rightReg)

	return "fa0"
}

func (eg *ExpressionGenerator) GenerateMultiplication(expr *ast.BinaryExpression) string {
	switch expr.Left.(type) {
	case *ast.IntegerLiteral:
		var leftInt int64 = expr.Left.(*ast.IntegerLiteral).Value
		var rightInt int64 = expr.Right.(*ast.IntegerLiteral).Value

		return eg.GenerateIntMultiplication(leftInt, rightInt)

	case *ast.FloatLiteral:
		var leftFloat float64 = expr.Left.(*ast.FloatLiteral).Value
		var rightFloat float64 = expr.Right.(*ast.FloatLiteral).Value

		return eg.GenerateFloatMultiplication(leftFloat, rightFloat)

	case *ast.Identifier:
		cg := eg.CodeGen

		var leftName string = expr.Left.(*ast.Identifier).Name
		var rightName string = expr.Right.(*ast.Identifier).Name

		var leftID string = leftName
		var rightID string = rightName

		leftSym, _ := cg.SymTable.Lookup(leftID)
		rightSym, _ := cg.SymTable.Lookup(rightID)

		switch leftSym.Type.(*ast.PrimitiveType).Name {
		case "int":
			return eg.GenerateIntMultiplication(leftSym.Value.(int64), rightSym.Value.(int64))
		case "float":
			return eg.GenerateFloatMultiplication(leftSym.Value.(float64), rightSym.Value.(float64))
		}
	}
	return "No case should reach here, as everything should be handled in semantic analysis"
}

func (eg *ExpressionGenerator) GenerateIntMultiplication(leftInt int64, rightInt int64) string {
	cg := eg.CodeGen
	rp := cg.Registers
	
	leftReg := rp.GetTmpRegister()
	rightReg := rp.GetTmpRegister()

	cg.emit("	li %s, %d", leftReg, leftInt)
	cg.emit("	li %s, %d", rightReg, rightInt)
	cg.emit("	mul a0, %s, %s", leftReg, rightReg)
	// The result will be a 128 bit integer, but for now we will just return the lower 64 bits
	// Meaning we will ignore overflow, very C-like

	return "a0"
}

func (eg *ExpressionGenerator) GenerateFloatMultiplication(leftFloat float64, rightFloat float64) string {
	cg := eg.CodeGen
	rp := cg.Registers
	
	cg.insertData("double_1", ".double", leftFloat)
	cg.insertData("double_2", ".double", rightFloat)

	leftAddressReg := rp.GetTmpRegister()
	rightAddressReg := rp.GetTmpRegister()
	// Load float values into registers
	leftReg := rp.GetFloatTmpRegister()
	rightReg := rp.GetFloatTmpRegister()
	// Load left float value
	cg.emit("	la %s, double_1", leftAddressReg)
	cg.emit("	fld %s, 0(%s)", leftReg, leftAddressReg)
	// Load right float value
	cg.emit("	la %s, double_2", rightAddressReg)
	cg.emit("	fld %s, 0(%s)", rightReg, rightAddressReg)
	// Perform subtraction
	cg.emit("	fmul.d fa0, %s, %s", leftReg, rightReg)

	return "fa0"
}

func (eg *ExpressionGenerator) GenerateSubtraction(expr *ast.BinaryExpression) string {
	switch expr.Left.(type) {
	case *ast.IntegerLiteral:
		var leftInt int64 = expr.Left.(*ast.IntegerLiteral).Value
		var rightInt int64 = expr.Right.(*ast.IntegerLiteral).Value

		return eg.GenerateIntSubtraction(leftInt, rightInt)

	case *ast.FloatLiteral:
		var leftFloat float64 = expr.Left.(*ast.FloatLiteral).Value
		var rightFloat float64 = expr.Right.(*ast.FloatLiteral).Value

		return eg.GenerateFloatSubtraction(leftFloat, rightFloat)

	case *ast.Identifier:
		cg := eg.CodeGen

		var leftName string = expr.Left.(*ast.Identifier).Name
		var rightName string = expr.Right.(*ast.Identifier).Name

		var leftID string = leftName
		var rightID string = rightName

		leftSym, _ := cg.SymTable.Lookup(leftID)
		rightSym, _ := cg.SymTable.Lookup(rightID)

		switch leftSym.Type.(*ast.PrimitiveType).Name {
		case "int":
			return eg.GenerateIntSubtraction(leftSym.Value.(int64), rightSym.Value.(int64))
		case "float":
			return eg.GenerateFloatSubtraction(leftSym.Value.(float64), rightSym.Value.(float64))
		}
	}
	return "No case should reach here, as everything should be handled in semantic analysis"
}

func (eg *ExpressionGenerator) GenerateIntSubtraction(leftInt int64, rightInt int64) string {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg := rp.GetTmpRegister()
	rightReg := rp.GetTmpRegister()

	cg.emit("	li %s, %d", leftReg, leftInt)
	cg.emit("	li %s, %d", rightReg, rightInt)
	cg.emit("	sub a0, %s, %s", leftReg, rightReg)

	return "a0"
}

func (eg *ExpressionGenerator) GenerateFloatSubtraction(leftFloat float64, rightFloat float64) string {
	cg := eg.CodeGen
	rp := cg.Registers

	cg.insertData("double_1", ".double", leftFloat)
	cg.insertData("double_2", ".double", rightFloat)

	leftAddressReg := rp.GetTmpRegister()
	rightAddressReg := rp.GetTmpRegister()
	// Load float values into registers
	leftReg := rp.GetFloatTmpRegister()
	rightReg := rp.GetFloatTmpRegister()
	// Load left float value
	cg.emit("	la %s, double_1", leftAddressReg)
	cg.emit("	fld %s, 0(%s)", leftReg, leftAddressReg)
	// Load right float value
	cg.emit("	la %s, double_2", rightAddressReg)
	cg.emit("	fld %s, 0(%s)", rightReg, rightAddressReg)
	// Perform subtraction
	cg.emit("	fsub.d fa0, %s, %s", leftReg, rightReg)

	return "fa0"

}

func (eg *ExpressionGenerator) GenerateAddition(expr *ast.BinaryExpression) string {

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

		return eg.GenerateIntAddition(leftInt, rightInt)

	case *ast.FloatLiteral:
		var leftFloat float64 = expr.Left.(*ast.FloatLiteral).Value
		var rightFloat float64 = expr.Right.(*ast.FloatLiteral).Value

		return eg.GenerateFloatAddition(leftFloat, rightFloat)

	case *ast.Identifier:
		cg := eg.CodeGen

		var leftName string = expr.Left.(*ast.Identifier).Name
		var rightName string = expr.Right.(*ast.Identifier).Name

		var leftID string = leftName
		var rightID string = rightName

		leftSym, _ := cg.SymTable.Lookup(leftID)
		rightSym, _ := cg.SymTable.Lookup(rightID)

		switch leftSym.Type.(*ast.PrimitiveType).Name {
		case "int":
			return eg.GenerateIntAddition(leftSym.Value.(int64), rightSym.Value.(int64))
		case "float":
			return eg.GenerateFloatAddition(leftSym.Value.(float64), rightSym.Value.(float64))
		}
	}
	return "No case should reach here, as everything should be handled in semantic analysis"
}

func (eg *ExpressionGenerator) GenerateIntAddition(leftInt int64, rightInt int64) string {
	cg := eg.CodeGen
	rp := cg.Registers

	if isImmediateInt(leftInt) {
		if isImmediateInt(rightInt) {
			// Both are immediate integers
			reg := rp.GetTmpRegister()
			cg.emit("	li %s, %d", reg, leftInt)
			cg.emit("	addi a0, %s, %d", reg, rightInt)
			// return reg
		} else {
			// Left is immediate, right is not
			// Load right operand into a register
			rightReg := rp.GetTmpRegister()
			cg.emit("	li %s, %d", rightReg, rightInt)
			cg.emit("	addi a0, %s, %d", rightReg, leftInt)
			// return rightReg
		}

	} else {
		if isImmediateInt(rightInt) {
			// Left is not immediate, right is
			// Load left operand into a register
			leftReg := rp.GetTmpRegister()
			cg.emit("	li %s, %d", leftReg, leftInt)
			cg.emit("	addi a0, %s, %d", leftReg, rightInt)
			// return leftReg
		} else {
			// Both are not immediate integers
			// Load both operands into registers
			leftReg := rp.GetTmpRegister()
			rightReg := rp.GetTmpRegister()
			cg.emit("	li %s, %d", leftReg, leftInt)
			cg.emit("	li %s, %d", rightReg, rightInt)
			cg.emit("	add a0, %s, %s", leftReg, rightReg)
			// return leftReg
		}
	}
	return "a0"
}

func (eg *ExpressionGenerator) GenerateFloatAddition(leftFloat float64, rightFloat float64) string {
	cg := eg.CodeGen
	rp := cg.Registers

	// Insert float data into the data section
	// Temporary names, will think of better names later
	cg.insertData("double_1", ".double", leftFloat)
	cg.insertData("double_2", ".double", rightFloat)

	leftAddressReg := rp.GetTmpRegister()
	rightAddressReg := rp.GetTmpRegister()
	// Load float values into registers
	leftReg := rp.GetFloatTmpRegister()
	rightReg := rp.GetFloatTmpRegister()
	// Load left float value
	cg.emit("	la %s, double_1", leftAddressReg)
	cg.emit("	fld %s, 0(%s)", leftReg, leftAddressReg)
	// Load right float value
	cg.emit("	la %s, double_2", rightAddressReg)
	cg.emit("	fld %s, 0(%s)", rightReg, rightAddressReg)
	// Perform addition
	cg.emit("	fadd.d fa0, %s, %s", leftReg, rightReg)
	// Should the result be stored inside a register or in the data section?
	// If the result is assigned to a variable, it should be stored in the data section
	// Else, it should be stored in a register
	// For now, we will store it in a register

	return "fa0"

}

func isImmediateInt(value int64) bool {
	return value >= -2048 && value <= 2047
}


func (eg *ExpressionGenerator) GenerateGreaterThan(expr *ast.BinaryExpression) string {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg := eg.GenerateExpression(expr.Left)
    rightReg := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()
	cg.emit("    slt %s, %s, %s", resultReg, rightReg, leftReg)

	if leftReg != "a0" && leftReg != "fa0" {
        rp.ReleaseRegister(leftReg)
    }
    if rightReg != "a0" && rightReg != "fa0" {
        rp.ReleaseRegister(rightReg)
    }

    return resultReg
}

func (eg *ExpressionGenerator) GenerateLessThan(expr *ast.BinaryExpression) string {
    cg := eg.CodeGen
    rp := cg.Registers

    leftReg := eg.GenerateExpression(expr.Left)
    rightReg := eg.GenerateExpression(expr.Right)

    resultReg := rp.GetTmpRegister()

    cg.emit("    slt %s, %s, %s", resultReg, leftReg, rightReg)

    if leftReg != "a0" && leftReg != "fa0" {
        rp.ReleaseRegister(leftReg)
    }
    if rightReg != "a0" && rightReg != "fa0" {
        rp.ReleaseRegister(rightReg)
    }

    return resultReg
}

func (eg *ExpressionGenerator) GenerateGreaterThanOrEqual(expr *ast.BinaryExpression) string {
    cg := eg.CodeGen
    rp := cg.Registers

    leftReg := eg.GenerateExpression(expr.Left)
    rightReg := eg.GenerateExpression(expr.Right)

    resultReg := rp.GetTmpRegister()
    tempReg := rp.GetTmpRegister()

    cg.emit("    slt %s, %s, %s", tempReg, leftReg, rightReg)
    cg.emit("    xori %s, %s, 1", resultReg, tempReg)  // Invert the result

    if leftReg != "a0" && leftReg != "fa0" {
        rp.ReleaseRegister(leftReg)
    }
    if rightReg != "a0" && rightReg != "fa0" {
        rp.ReleaseRegister(rightReg)
    }
    rp.ReleaseRegister(tempReg)

    return resultReg
}

func (eg *ExpressionGenerator) GenerateLessThanOrEqual(expr *ast.BinaryExpression) string {
    cg := eg.CodeGen
    rp := cg.Registers

    leftReg := eg.GenerateExpression(expr.Left)
    rightReg := eg.GenerateExpression(expr.Right)

    resultReg := rp.GetTmpRegister()
    tempReg := rp.GetTmpRegister()

    cg.emit("    slt %s, %s, %s", tempReg, rightReg, leftReg)
    cg.emit("    xori %s, %s, 1", resultReg, tempReg)  // Invert the result

   	if leftReg != "a0" && leftReg != "fa0" {
        rp.ReleaseRegister(leftReg)
    }
    if rightReg != "a0" && rightReg != "fa0" {
        rp.ReleaseRegister(rightReg)
    }
    rp.ReleaseRegister(tempReg)

    return resultReg
}

func (eg *ExpressionGenerator) GenerateEquality(expr *ast.BinaryExpression) string {
    cg := eg.CodeGen
    rp := cg.Registers

    leftReg := eg.GenerateExpression(expr.Left)
    rightReg := eg.GenerateExpression(expr.Right)

    resultReg := rp.GetTmpRegister()

    cg.emit("    sub %s, %s, %s", resultReg, leftReg, rightReg)
    cg.emit("    seqz %s, %s", resultReg, resultReg)  // Set to 1 if equal to zero

    if leftReg != "a0" && leftReg != "fa0" {
        rp.ReleaseRegister(leftReg)
    }
    if rightReg != "a0" && rightReg != "fa0" {
        rp.ReleaseRegister(rightReg)
    }

    return resultReg
}

func (eg *ExpressionGenerator) GenerateInequality(expr *ast.BinaryExpression) string {
    cg := eg.CodeGen
    rp := cg.Registers

    leftReg := eg.GenerateExpression(expr.Left)
    rightReg := eg.GenerateExpression(expr.Right)

    resultReg := rp.GetTmpRegister()

    cg.emit("    sub %s, %s, %s", resultReg, leftReg, rightReg)
    cg.emit("    snez %s, %s", resultReg, resultReg)  // Set to 1 if not equal to zero

    if leftReg != "a0" && leftReg != "fa0" {
        rp.ReleaseRegister(leftReg)
    }
    if rightReg != "a0" && rightReg != "fa0" {
        rp.ReleaseRegister(rightReg)
    }

    return resultReg
}
