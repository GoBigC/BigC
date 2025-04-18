package codegen 

import (
    "strings"
    "BigCooker/pkg/semantic/table"
    "BigCooker/pkg/syntax/ast"
)

type CodeGenerator struct {
    Program     *ast.Program        // ast root
    SymTable    *table.SymbolTable  // symbol table
    AsmOut      *strings.Builder    // assembly string output

    // program state tracking 
    CurrentFunction string 
    Labels          int
    Registers       *RegisterPool
    StackSize       int
    VarStackOffset  map[string]int

    //
    ExpressionGen   *ExpressionGenerator 
    AssignmentGen   *AssignmentGenerator
    BranchingGen    *BranchingGenerator
    FunctionGen     *FunctionGenerator
    LoopingGen      *LoopingGenerator
}

type RegisterPool struct {
    TmpRegs     []string
    ArgRegs     []string
    SavedRegs   []string
    SpecialRegs []string
    InUse       map[string] bool
}