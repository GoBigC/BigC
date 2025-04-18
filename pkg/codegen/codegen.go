package codegen

import (
	"BigCooker/pkg/semantic/table"
	"BigCooker/pkg/syntax/ast"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type RegisterPool struct {
    TmpRegs     []string // t0-t6
    ArgRegs     []string // a0-a7
    SavedRegs   []string // s1-s11
    Reserved    []string // zero, ra, sp, gp, tp, s0/fp

    InUse       map[string]bool
}

func NewRegisterPool() *RegisterPool {
    tmps := []string{"t0", "t1", "t2", "t3", "t4", "t5", "t6"}
    args := []string{"a0", "a1", "a2", "a3", "a4", "a5", "a6", "a7"}
    saves := []string{"s1", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "s10", "s11"}
    rsv := []string{"zero", "ra", "sp", "gp", "tp", "s0", "fp"}

    return &RegisterPool{
        TmpRegs:    tmps, 
        ArgRegs:    args, 
        SavedRegs:  saves, 
        Reserved:   rsv,
        InUse:      make(map[string]bool),
    }
}

/*
// Closer to a standard allocate register function 
// But seemingly too advanced for our usecase 
// For now, just use GetXRegister() to allocate 
// register from a specific type (temp, arg, saved, etc.)
// I keep this here for future references 

func (rp *RegisterPool) GetRegister(preferCalleeSaved bool) string {
    primaryPool := rp.TmpRegs
    secondaryPool := rp.ArgRegs
    tertiaryPool := rp.SavedRegs

    if preferCalleeSaved { // if need registers that span between functions
        primaryPool := rp.SavedRegs
        secondaryPool := rp.TmpRegs
        tertiaryPool := rp.ArgRegs
    }
    // ...
}
*/

func (rp *RegisterPool) GetTmpRegister() string {
    return rp.AllocateRegisterFallback([][]string{
        rp.TmpRegs,
        rp.ArgRegs[2:], 
        rp.ArgRegs[:2], // a0-a1 usually for return value, try to not use them
        rp.SavedRegs,   // last resort
    })
}

func (rp *RegisterPool) GetArgRegister(index int) string {
    if index >=0 && index < len(rp.ArgRegs) {
        reg := rp.ArgRegs[index]
        if !rp.InUse[reg]{
            rp.InUse[reg] = true 
            return reg
        }
    }

    return rp.GetTmpRegister()
}

func (rp *RegisterPool) GetSavedRegister() string {
    return rp.AllocateRegisterFallback([][]string{
        rp.SavedRegs, 
        rp.TmpRegs, 
        rp.ArgRegs[2:], 
        rp.ArgRegs[:2], 
    })
}

func (rp *RegisterPool) AllocateRegisterFallback(regGroups [][]string) string {
    for _, group := range regGroups {
        for _, reg := range group {
            if !rp.InUse[reg] && !contains(rp.Reserved, reg){
                rp.InUse[reg] = true 
                return reg 
            }
        }
    }

    panic("No register available. Don't know what do.")
}

func contains(arr []string, item string) bool {
    for _, i := range arr {
        if i == item {
            return true
        }
    }
    return false 
}

func (rp *RegisterPool) ReleaseRegister(reg string) {
    rp.InUse[reg] = false 
}

func (rp *RegisterPool) GetAllUsedRegisters() [] string {
	var used []string
	for _, reg := range rp.SavedRegs {
		if rp.InUse[reg] {
			used = append(used, reg)
		}
	}
	return used
}


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

func NewCodeGenerator(program *ast.Program, symTable *table.SymbolTable) *CodeGenerator {
    cg := &CodeGenerator{
        Program:        program, 
        SymTable:       symTable, 
        Labels:         0, 
        Registers:      NewRegisterPool(), 
        VarStackOffset: make(map[string]int),
    }

    cg.ExpressionGen = NewExpressionGenerator(cg)
    cg.AssignmentGen = NewAssignmentGenerator(cg)
	cg.BranchingGen  = NewBranchingGenerator(cg)
	cg.FunctionGen   = NewFunctionGenerator(cg)
	cg.LoopingGen    = NewLoopingGenerator(cg)

    return cg
}

func (cg *CodeGenerator) emit(format string, args ...interface{}) {
    instruction := fmt.Sprintf(format, args...)
    cg.AsmOut.WriteString(instruction + "\n")
}

func (cg *CodeGenerator) emitComment(format string, args ...interface{}) {
	comment := fmt.Sprintf("# "+format, args...)
	cg.AsmOut.WriteString(comment + "\n")
}

func (cg *CodeGenerator) Generate() error {
    outFile := "asm.asm"

    cg.emit(".text")
    cg.emit(".globl main") // first function is main

    for _, decl := range cg.Program.Declarations {
        cg.generateDeclaration(decl)
    }

    err := os.MkdirAll(filepath.Dir(outFile), 0777)
    if err != nil {
        return fmt.Errorf("Cannot create output file: %w", err)
    }

    err = os.WriteFile(outFile, []byte(cg.AsmOut.String()), 0777)
    if err != nil {
        return fmt.Errorf("Failed to write assembly to file: %w", err)
    }

    return nil
}