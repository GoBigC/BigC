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
	cg := eg.CodeGen
	rp := cg.Registers

	reg := rp.GetTmpRegister()
	if expr.Value {
		cg.emit("    li %s, 1", reg)
	} else {
		cg.emit("    li %s, 0", reg)
	}
	return reg, &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateCharLiteral(expr *ast.CharLiteral) (string, ast.Type) {
	reg := eg.CodeGen.Registers.GetTmpRegister()
	eg.CodeGen.emit("    li %s, %d", reg, expr.Value)
	return reg, &ast.PrimitiveType{Name: "char"}
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

/*
Have: 
	a: .space N <-- correctly allocating N consecutive bytes in memory, stores
					address of first element in a
Have to do next to access a[i]: 
	1. Load base address of array a
	2. Calculate offset: i*elem_size
	3. Pointer arithmetic: Add offset to base --> get the location of the i-th element
	4. Load value from / Write value to that calculated address 
*/

func (eg *ExpressionGenerator) GenerateArrayAccessExpression(e *ast.ArrayAccessExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	elemAddrRegister, indexRegister, elemType := eg.CalculateArrayElementAddress(e.Array, e.Index)

	// 4. load value from element address 
	var resultRegister string 
	isFloat := false
    if primType, ok := elemType.(*ast.PrimitiveType); ok {
        isFloat = primType.Name == "float"
    }

	if isFloat {
		resultRegister = rp.GetFloatTmpRegister()
		cg.emit("	fld %s, 0(%s)", resultRegister, elemAddrRegister)
	} else {
		resultRegister = rp.GetTmpRegister()
		cg.emit("	ld %s, 0(%s)", resultRegister, elemAddrRegister)
	}

	if elemAddrRegister != "a0" && elemAddrRegister != "fa0" {
		rp.ReleaseRegister(elemAddrRegister)
	}
	if indexRegister != "a0" && indexRegister != "fa0" {
		rp.ReleaseRegister(indexRegister)
	}

	return resultRegister, elemType
}

func (eg *ExpressionGenerator) CalculateArrayElementAddress(arrExpr ast.Expression, indexExpr ast.Expression) (string, string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	var arrayName string
	var elemType ast.Type 

	if id, ok := arrExpr.(*ast.Identifier); ok {
		// this is to extract 'a' from 'a[i]' 
		arrayName = id.Name 
	}

	symbol, found := cg.SymTable.Lookup("main." + arrayName) // find local first
	if !found {
		symbol, found = cg.SymTable.Lookup(arrayName) // try global
	}

	if found {
		if arrayType, ok := symbol.Type.(*ast.ArrayType); ok {
			elemType = arrayType.ElementType
		}
	} else {
		panic("Dont know what to do with array expression--maybe too complex?")
	}

	// 1. get array base reg
	baseAddrRegister := rp.GetTmpRegister()
	cg.emit("	la %s, %s", baseAddrRegister, arrayName)

	// 2. get index value --> calculate offset bytes by mul 8 
	indexRegister, _ := eg.GenerateExpression(indexExpr)
	offsetValueRegister := rp.GetTmpRegister()
	cg.emit("	li %s, 8", offsetValueRegister)
	cg.emit("	mul %s, %s, %s", offsetValueRegister, indexRegister, offsetValueRegister)

	// 3. pointer arithmetic: a[i] = a + i*8
	elemAddrRegister := rp.GetTmpRegister()
	cg.emit("	add %s, %s, %s", elemAddrRegister, offsetValueRegister, baseAddrRegister)

	if baseAddrRegister != "a0" && baseAddrRegister != "fa0" {
		rp.ReleaseRegister(baseAddrRegister)
	}
	if offsetValueRegister != "a0" && offsetValueRegister != "fa0" {
		rp.ReleaseRegister(offsetValueRegister)
	}
	
	return elemAddrRegister, indexRegister, elemType
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
	if expr.Operator == "=" {
		if arrayAccess, ok := expr.Left.(*ast.ArrayAccessExpression); ok {
			eg.CodeGen.AssignmentGen.GenerateArrayAssignment(arrayAccess, expr.Right)
			return "a0", &ast.PrimitiveType{Name: "void"} // assignment dont return value
		}
	}
	
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
	case "&&":
		return eg.GenerateLogicalAnd(expr)
	case "||":
		return eg.GenerateLogicalOr(expr)
	default:
		panic("GenerateBinaryExpression - No case should reach here, as everything should be handled in semantic analysis")
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
		releaseRegAfterUse(*rp, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "int"}
	case "float":
		resultReg = rp.GetFloatTmpRegister()
		cg.emit("	fdiv.d %s, %s, %s", resultReg, leftReg, rightReg)
		releaseRegAfterUse(*rp, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "float"}
	}


	panic("GenerateDivision - No case should reach here, as everything should be handled in semantic analysis")

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
		releaseRegAfterUse(*rp, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "int"}
	case "float":
		resultReg = rp.GetFloatTmpRegister()
		cg.emit("	fmul.d %s, %s, %s", resultReg, leftReg, rightReg)
		releaseRegAfterUse(*rp, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "float"}
	}

	panic("GenerateMultiplication - No case should reach here, as everything should be handled in semantic analysis")
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
		releaseRegAfterUse(*rp, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "int"}
	case "float":
		resultReg = rp.GetFloatTmpRegister()
		cg.emit("	fsub.d %s, %s, %s", resultReg, leftReg, rightReg)
		releaseRegAfterUse(*rp, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "float"}
	}


	panic("GenerateSubtraction - No case should reach here, as everything should be handled in semantic analysis")
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
		releaseRegAfterUse(*rp, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "int"}
	case "float":
		resultReg = rp.GetFloatTmpRegister()
		cg.emit("	fadd.d %s, %s, %s", resultReg, leftReg, rightReg)
		releaseRegAfterUse(*rp, leftReg, rightReg)
		return resultReg, &ast.PrimitiveType{Name: "float"}
	}


	panic("GenerateAddition - No case should reach here, as everything should be handled in semantic analysis")
}

func (eg *ExpressionGenerator) GenerateGreaterThan(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()
	cg.emit("    slt %s, %s, %s", resultReg, rightReg, leftReg)

	releaseRegAfterUse(*rp, leftReg, rightReg)

	return resultReg, &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateLessThan(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()

	cg.emit("    slt %s, %s, %s", resultReg, leftReg, rightReg)

	releaseRegAfterUse(*rp, leftReg, rightReg)

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

	releaseRegAfterUse(*rp, leftReg, rightReg)
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

	releaseRegAfterUse(*rp, leftReg, rightReg)
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

	releaseRegAfterUse(*rp, leftReg, rightReg)

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

	releaseRegAfterUse(*rp, leftReg, rightReg)

	return resultReg, &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateLogicalOr(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()

	cg.emit("    or %s, %s, %s", resultReg, leftReg, rightReg)

	releaseRegAfterUse(*rp, leftReg, rightReg)

	return resultReg, &ast.PrimitiveType{Name: "bool"}
}

func (eg *ExpressionGenerator) GenerateLogicalAnd(expr *ast.BinaryExpression) (string, ast.Type) {
	cg := eg.CodeGen
	rp := cg.Registers

	leftReg, _ := eg.GenerateExpression(expr.Left)
	rightReg, _ := eg.GenerateExpression(expr.Right)

	resultReg := rp.GetTmpRegister()

	cg.emit("    and %s, %s, %s", resultReg, leftReg, rightReg)

	releaseRegAfterUse(*rp, leftReg, rightReg)

	return resultReg, &ast.PrimitiveType{Name: "bool"}
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

func releaseRegAfterUse(rp RegisterPool, leftReg, rightReg string) {
	if leftReg != "a0" && leftReg != "fa0" {
		rp.ReleaseRegister(leftReg)
	}
	if rightReg != "a0" && rightReg != "fa0" {
		rp.ReleaseRegister(rightReg)
	}
}