package codegen

import (
	// "BigCooker/pkg/syntax/ast"
)

type LoopingGenerator struct {
    CodeGen     *CodeGenerator
}

func NewLoopingGenerator(cg *CodeGenerator) *LoopingGenerator {
	return &LoopingGenerator{
		CodeGen: cg,
	}
}
