package codegen

import (
	"BigCooker/pkg/syntax/ast"
)

type BranchingGenerator struct {
    CodeGen     *CodeGenerator
}

func NewBranchingGenerator(cg *CodeGenerator) *BranchingGenerator {
	return &BranchingGenerator{
		CodeGen: cg,
	}
}

func (bg *BranchingGenerator) GenerateIfStatement(stmt ast.IfStatement) {}