package codegen

import (
	"BigCooker/pkg/syntax/ast"
)

type FunctionGenerator struct {
    CodeGen     *CodeGenerator
}

func NewFunctionGenerator(cg *CodeGenerator) *FunctionGenerator {
	return &FunctionGenerator{
		CodeGen: cg,
	}
}

func (fg *FunctionGenerator) GenerateFunctionDeclaration(funcDecl ast.FunctionDeclaration) {
	fg.CodeGen.CurrentFunction = funcDecl.Name
	
}