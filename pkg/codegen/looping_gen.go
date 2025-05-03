package codegen

import (
	"BigCooker/pkg/syntax/ast"
	"fmt"
)

type LoopingGenerator struct {
	CodeGen           *CodeGenerator
	StartLabelCounter int
	EndLabelCounter   int
}

func NewLoopingGenerator(cg *CodeGenerator) *LoopingGenerator {
	return &LoopingGenerator{
		CodeGen:           cg,
		StartLabelCounter: 0,
		EndLabelCounter:   0,
	}
}

func (lg *LoopingGenerator) GenerateWhileStatement(stmt *ast.WhileStatement) {
	cg := lg.CodeGen
	rp := cg.Registers

	cg.emitComment("=== Begin while statement ===")
	startLabel := lg.NewStartLabel()
	endLabel := lg.NewEndLabel()

	cg.emit("%s:", startLabel) // Start of the loop
	cg.emitComment("Condition:")
	condReg, condType := cg.ExpressionGen.GenerateExpression(stmt.Condition)

	if condType.(*ast.PrimitiveType).Name != "bool" {
		panic("Condition must be a boolean type - this shoulld have been caught in semantic analysis")
	}

	cg.emit("    beq %s, x0, %s", condReg, endLabel) // If condition is false, jump to end
	rp.ReleaseRegister(condReg)

	cg.emitComment("=== While body ===")
	cg.BlockGen.GenerateBlock(stmt.Body)

	// Jump back to start
	cg.emitComment("=== Jump back to start ===")
	cg.emit("    j %s", startLabel)

	cg.emitComment("=== End of while statement ===")
	// Emit loop end label
	cg.emit("%s:", endLabel)

}

func (lg *LoopingGenerator) NewStartLabel() string {
	label := fmt.Sprintf("loop_%d_start", lg.StartLabelCounter)
	lg.StartLabelCounter++
	return label
}

func (lg *LoopingGenerator) NewEndLabel() string {
	label := fmt.Sprintf("loop_%d_end", lg.EndLabelCounter)
	lg.EndLabelCounter++
	return label
}
