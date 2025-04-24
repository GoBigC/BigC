package codegen

import (
	"BigCooker/pkg/syntax/ast"
)


type AssignmentGenerator struct {
    CodeGen     *CodeGenerator 
}

func NewAssignmentGenerator(cg *CodeGenerator) *AssignmentGenerator {
	return &AssignmentGenerator{
		CodeGen: cg,
	}
}

func (ag *AssignmentGenerator) GenerateVarDeclaration(vardecl ast.VarDeclaration) {
	// code here -- this function was created to bypass the go errors
}