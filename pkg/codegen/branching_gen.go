package codegen

import (
	"BigCooker/pkg/syntax/ast"
)

type BranchingGenerator struct {
	CodeGen *CodeGenerator
}

func NewBranchingGenerator(cg *CodeGenerator) *BranchingGenerator {
	return &BranchingGenerator{
		CodeGen: cg,
	}
}

func (bg *BranchingGenerator) GenerateIfStatement(stmt *ast.IfStatement) {
	bg.CodeGen.emitComment("Begin if statement")

	// Skip condition evaluation if nil (for testing or placeholder statements)
	var condReg string
	if stmt.Condition != nil {
		// Generate code for the condition expression
		condReg = bg.CodeGen.ExpressionGen.GenerateExpression(stmt.Condition)
	} else {
		condReg = "x0" // Always false if no condition
	}

	elseLabel := bg.CodeGen.NewLabel()
	endLabel := bg.CodeGen.NewLabel()

	// Branch if condition is false
	bg.CodeGen.emit("beq %s, x0, %s", condReg, elseLabel)

	// Generate code for the then block
	bg.CodeGen.emitComment("Then block:")
	if stmt.ThenBlock != nil {
		bg.GenerateBlock(stmt.ThenBlock)
	}

	// Handle else block
	if stmt.ElseBlock != nil {
		bg.CodeGen.emit("j %s", endLabel)
		bg.CodeGen.emit("%s:", elseLabel)

		switch elseNode := stmt.ElseBlock.(type) {
		case *ast.Block:
			bg.GenerateBlock(elseNode)
		case *ast.IfStatement:
			bg.GenerateIfStatement(elseNode)
		default:
			bg.CodeGen.emitComment("Unsupported else block type: %T", stmt.ElseBlock)
		}
	} else {
		bg.CodeGen.emit("%s:", elseLabel)
	}

	bg.CodeGen.emit("%s:", endLabel)

	// Release the register used for condition
	if stmt.Condition != nil {
		bg.CodeGen.Registers.ReleaseRegister(condReg)
	}

	bg.CodeGen.emitComment("End if statement")
}

func (bg *BranchingGenerator) GenerateBlock(block *ast.Block) {
	for i, item := range block.Items {
		bg.CodeGen.emitComment("Statement #%d", i+1)
		bg.GenerateBlockItem(item)
	}
}

func (bg *BranchingGenerator) GenerateBlockItem(item ast.BlockItem) {
	switch stmt := item.(type) {
	case *ast.ExpressionStatement:
		bg.GenerateExpressionStatement(stmt)
	case *ast.VarDeclaration:
		bg.CodeGen.AssignmentGen.GenerateVarDeclaration(*stmt)
	case *ast.IfStatement:
		bg.GenerateIfStatement(stmt)
	// Add more statement types as needed
	default:
		bg.CodeGen.emitComment("Unsupported statement type: %T", item)
	}
}

func (bg *BranchingGenerator) GenerateExpressionStatement(stmt *ast.ExpressionStatement) {
	// Check if the expression is a function call
	if funcCall, ok := stmt.Expr.(*ast.FunctionCallExpression); ok {
		bg.GenerateFunctionCall(funcCall)
	} else {
		// Handle other types of expressions
		reg := bg.CodeGen.ExpressionGen.GenerateExpression(stmt.Expr)
		bg.CodeGen.Registers.ReleaseRegister(reg)
	}
}

func (bg *BranchingGenerator) GenerateFunctionCall(call *ast.FunctionCallExpression) {
	// Check if it's a builtin function
	if ident, ok := call.Function.(*ast.Identifier); ok {
		funcName := ident.Name

		// Handle arguments
		for i, arg := range call.Arguments {
			// Generate code to evaluate the argument
			argReg := bg.CodeGen.ExpressionGen.GenerateExpression(arg)

			// Move the argument to the appropriate argument register
			if isIntegerFunction(funcName) {
				bg.CodeGen.emit("mv a%d, %s", i, argReg)
			} else if isFloatFunction(funcName) {
				bg.CodeGen.emit("fmv.s fa%d, %s", i, argReg)
			}

			// Release the temporary register
			bg.CodeGen.Registers.ReleaseRegister(argReg)
		}

		// Call the function
		bg.CodeGen.emit("call %s", funcName)
	} else {
		// Handle function calls with non-identifier function expressions
		// This is for more complex cases like function pointers
		bg.CodeGen.emitComment("Function call with non-identifier function")
		funcReg := bg.CodeGen.ExpressionGen.GenerateExpression(call.Function)

		// Handle arguments...

		// Call through register
		bg.CodeGen.emit("jalr %s", funcReg)
		bg.CodeGen.Registers.ReleaseRegister(funcReg)
	}
}

// Helper function to determine if a function uses integer registers
func isIntegerFunction(name string) bool {
	// Add more as needed
	switch name {
	case "_printInt", "_printBool", "_printChar", "_printString":
		return true
	default:
		return true // Default to integer for unknown functions
	}
}

// Helper function to determine if a function uses float registers
func isFloatFunction(name string) bool {
	// Add more as needed
	switch name {
	case "_printFloat":
		return true
	default:
		return false
	}
}
