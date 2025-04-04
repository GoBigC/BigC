// codegen/codegen.go
package codegen

import (
	"strings"
)

// CodeGenerator holds state for generating RISC-V assembly
type CodeGenerator struct {
    output      strings.Builder
    regManager  *RegisterManager // Defined in registers.go
    labelCount  int
    variables   map[string]string // Maps variable names to types
}

// NewCodeGenerator initializes a new code generator
func NewCodeGenerator() *CodeGenerator {
    return &CodeGenerator{
        regManager: NewRegisterManager(),
        variables:  make(map[string]string),
        labelCount: 0,
    }
}

// Generate produces RISC-V assembly from an AST Program
func (cg *CodeGenerator) Generate(program *Program) string {
    // Data section: Declare variables
    cg.output.WriteString(".data\n")
    cg.collectVariables(program)
    for varName, varType := range cg.variables {
        cg.output.WriteString("    " + varName + ": " + cg.typeToDirective(varType) + " 0\n")
    }

    // Text section: Code
    cg.output.WriteString(".text\n    main:\n")
    for _, decl := range program.Declarations {
        cg.generateDeclaration(&decl)
    }
    cg.output.WriteString("    li a7, 10\n    ecall\n") // Exit syscall

    return cg.output.String()
}

// collectVariables gathers all variable declarations
func (cg *CodeGenerator) collectVariables(program *Program) {
    for _, decl := range program.Declarations {
        cg.variables[decl.Name] = decl.Type
        // Handle arrays or function parameters later
    }
}