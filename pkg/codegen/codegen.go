package codegen

import (
	"BigCooker/pkg/semantic/table"
	"BigCooker/pkg/syntax/ast"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type RegisterPool struct {
	TmpRegs        []string // t0-t6
	ArgRegs        []string // a0-a7
	SavedRegs      []string // s1-s11
	Reserved       []string // zero, ra, sp, gp, tp, s0/fp
	FloatTmpRegs   []string // ft0-ft11
	FloatArgRegs   []string // fa0-fa7
	FloatSavedRegs []string // fs1-fs11

	InUse map[string]bool
}

func NewRegisterPool() *RegisterPool {
	tmps := []string{"t0", "t1", "t2", "t3", "t4", "t5", "t6"}
	args := []string{"a0", "a1", "a2", "a3", "a4", "a5", "a6", "a7"}
	saves := []string{"s1", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "s10", "s11"}
	rsv := []string{"zero", "ra", "sp", "gp", "tp", "s0", "fp"}

	floatTmps := []string{"ft0", "ft1", "ft2", "ft3", "ft4", "ft5", "ft6", "ft7", "ft8", "ft9", "ft10", "ft11"}
	floatArgs := []string{"fa0", "fa1", "fa2", "fa3", "fa4", "fa5", "fa6", "fa7"}
	floatSaves := []string{"fs1", "fs2", "fs3", "fs4", "fs5", "fs6", "fs7", "fs8", "fs9", "fs10", "fs11"}

	return &RegisterPool{
		TmpRegs:        tmps,
		ArgRegs:        args,
		SavedRegs:      saves,
		Reserved:       rsv,
		FloatTmpRegs:   floatTmps,
		FloatArgRegs:   floatArgs,
		FloatSavedRegs: floatSaves,

		InUse: make(map[string]bool),
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
	if index >= 0 && index < len(rp.ArgRegs) {
		reg := rp.ArgRegs[index]
		if !rp.InUse[reg] {
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

func (rp *RegisterPool) GetFloatTmpRegister() string {
	return rp.AllocateRegisterFallback([][]string{
		rp.FloatTmpRegs,
		rp.FloatArgRegs[2:],
		rp.FloatArgRegs[:2],
		rp.FloatSavedRegs,
	})
}

func (rp *RegisterPool) GetFloatArgRegister(index int) string {
	if index >= 0 && index < len(rp.FloatArgRegs) {
		reg := rp.FloatArgRegs[index]
		if !rp.InUse[reg] {
			rp.InUse[reg] = true
			return reg
		}
	}

	return rp.GetFloatTmpRegister()
}

func (rp *RegisterPool) GetFloatSavedRegister() string {
	return rp.AllocateRegisterFallback([][]string{
		rp.FloatSavedRegs,
		rp.FloatTmpRegs,
		rp.FloatArgRegs[2:],
		rp.FloatArgRegs[:2],
	})
}

func (rp *RegisterPool) AllocateRegisterFallback(regGroups [][]string) string {
	for _, group := range regGroups {
		for _, reg := range group {
			if !rp.InUse[reg] && !contains(rp.Reserved, reg) {
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

func (rp *RegisterPool) GetAllUsedRegisters() []string {
	var used []string
	for _, reg := range rp.SavedRegs {
		if rp.InUse[reg] {
			used = append(used, reg)
		}
	}
	return used
}

type CodeGenerator struct {
	Program  *ast.Program       // ast root
	SymTable *table.SymbolTable // symbol table
	AsmOut   *strings.Builder   // assembly string output

	// program state tracking
	Labels         int
	Registers      *RegisterPool
	StackSize      int64
	VarStackOffset map[string]int64

	//
	ExpressionGen *ExpressionGenerator
	AssignmentGen *AssignmentGenerator
	BranchingGen  *BranchingGenerator
	LoopingGen    *LoopingGenerator
	BlockGen      *BlockGenerator
}

func NewCodeGenerator(program *ast.Program, symTable *table.SymbolTable) *CodeGenerator {
	cg := &CodeGenerator{
		Program:        program,
		SymTable:       symTable,
		AsmOut:         &strings.Builder{},
		Labels:         0,
		Registers:      NewRegisterPool(),
		VarStackOffset: make(map[string]int64),
	}

	cg.ExpressionGen = NewExpressionGenerator(cg)
	cg.AssignmentGen = NewAssignmentGenerator(cg)
	cg.BranchingGen = NewBranchingGenerator(cg)
	cg.LoopingGen = NewLoopingGenerator(cg)
	cg.BlockGen = NewBlockGenerator(cg)

	return cg
}

func (cg *CodeGenerator) AllocateStack(symbol table.Symbol) int64 {
	var size int64
	// Check if the variable is an array
    if _, isArray := symbol.Type.(*ast.ArrayType); isArray {
        // Array: 8 bytes per element * array size
        size = 8 * symbol.ArraySize
    } else {
        // Scalar types: int, float, bool, char
        size = 8
    }

    cg.StackSize += size
    offset := cg.StackSize - size // Offset from the adjusted sp
    cg.VarStackOffset[symbol.Name] = offset
    return offset
}

func (cg *CodeGenerator) ResetStack() {
	cg.VarStackOffset = make(map[string]int64)
	cg.StackSize = 0
}

func (cg *CodeGenerator) GetStackOffset(symbolName string) int64 {
	if offset, ok := cg.VarStackOffset[symbolName]; ok {
		return offset
	}
	return 0 // Indicates global or error
}

func (cg *CodeGenerator) emit(format string, args ...interface{}) {
	instruction := fmt.Sprintf(format, args...)
	cg.AsmOut.WriteString(instruction + "\n")
}

func (cg *CodeGenerator) emitComment(format string, args ...interface{}) {
	comment := fmt.Sprintf("# "+format, args...)
	cg.AsmOut.WriteString(comment + "\n")
}

// Insert data into the ".data" section of the assembly output
// Highly inefficient, copies and creates a new string every time
// We need it to work first, then we can optimize it later
func (cg *CodeGenerator) insertData(label string, dataType string, value any) error {
	currentContent := cg.AsmOut.String()

	// use regexp to match more accurately than .Contains()
	pattern := fmt.Sprintf(`(?m)^%s:[\s]`, regexp.QuoteMeta(label))
	alreadyExists, err := regexp.MatchString(pattern, currentContent)
	if err != nil {
		fmt.Printf("WARNING: Error checking for existing label %s: %v\n", label, err)
	}

	if alreadyExists {
		fmt.Printf("DEBUG: Skipping insertion of %s, already exists\n", label)
		return nil
	}

	dataMarker := ".data\n"
	pos := strings.Index(currentContent, dataMarker)
	if pos == -1 {
		return fmt.Errorf("marker %q not found", dataMarker)
	}

	insertPos := pos + len(dataMarker)

	var newLabel string
	switch v := value.(type) {
	case int, int64, int32, int16, int8:
		newLabel = fmt.Sprintf("%s: %s %d\n", label, dataType, v)
	case float64, float32:
		newLabel = fmt.Sprintf("%s: %s %.6f\n", label, dataType, v)
	case string:
		if dataType == ".asciz" || dataType == ".ascii" {
			newLabel = fmt.Sprintf("%s: %s \"%s\"\n", label, dataType, v)
		} else {
			newLabel = fmt.Sprintf("%s: %s %s\n", label, dataType, v)
		}
	case []byte:
		var bytes strings.Builder
		for i, b := range v {
			if i > 0 {
				bytes.WriteString(", ")
			}
			bytes.WriteString(fmt.Sprintf("0x%02x", b))
		}
		newLabel = fmt.Sprintf("%s: %s %s\n", label, dataType, bytes.String())
	case bool:
		intValue := 0
		if v {
			intValue = 1
		}
		newLabel = fmt.Sprintf("%s: %s %d\n", label, dataType, intValue)
	default:
		newLabel = fmt.Sprintf("%s: %s %v\n", label, dataType, v)
	}

	fmt.Printf("DEBUG: Inserting %s: %s into .data section\n", label, dataType)

	// Create a new strings.Builder to hold the updated content
	var newBuilder strings.Builder
	// Write content before insertion point
	newBuilder.WriteString(currentContent[:insertPos])
	// Write the new label
	newBuilder.WriteString(newLabel)
	// Write content after insertion point
	newBuilder.WriteString(currentContent[insertPos:])

	// Reset the original builder and write the updated content
	cg.AsmOut.Reset()
	cg.AsmOut.WriteString(newBuilder.String())

	return nil
}

func (cg *CodeGenerator) GenerateProgram(outFile string) error { //renamed Generate()
	cg.emit(".data")

	cg.emit(".text")
	cg.emit("j main") // first function is main
	cg.GenerateAllBuiltinFunctions()
	cg.emit("main:")

	// Calculate stack size for main
	var stackSize int64
	for _, decl := range cg.Program.Declarations {
		if funcDecl, ok := decl.(*ast.FunctionDeclaration); ok && funcDecl.Name == "main" {
			stackSize = cg.CalculateStackSize(funcDecl.Body)
		}
	}

	// Prologue: Allocate stack frame
	if stackSize > 0 {
		cg.emitComment("=== main Prologue ===")
		cg.emit("    addi sp, sp, -%d", stackSize)
	}

	for _, decl := range cg.Program.Declarations {
		if funcDecl, ok := decl.(*ast.FunctionDeclaration); ok {
			if funcDecl.Name == "main" {
				for _, stmt := range funcDecl.Body.Items {
					cg.GenerateStatement(stmt)
				}
			}
		} else {
			cg.GenerateDeclaration(decl)
		}
	}

	if stackSize > 0 {
		cg.emitComment("=== main Epilogue ===")
		cg.emit("    addi sp, sp, %d", stackSize)
	}

	cg.emit("	li a0, 0")
	cg.emit("	j _exit")

	err := os.MkdirAll(filepath.Dir(outFile), 0777)
	if err != nil {
		return fmt.Errorf("cannot create output file: %w", err)
	}

	err = os.WriteFile(outFile, []byte(cg.AsmOut.String()), 0777)
	if err != nil {
		return fmt.Errorf("failed to write assembly to file: %w", err)
	}
	return nil
}

func (cg *CodeGenerator) GenerateDeclaration(decl ast.Declaration) {
	switch d := decl.(type) {
	case *ast.VarDeclaration:
		cg.AssignmentGen.GenerateVarDeclaration(d)
	// add more cases as we generate
	default:
		panic(fmt.Sprintf("Cannot generate code for unknown declaration type: %T", decl))
	}
}

func (cg *CodeGenerator) GenerateStatement(item ast.BlockItem) {
	switch stmt := item.(type) {
	case *ast.ExpressionStatement:
		cg.ExpressionGen.GenerateExpression(stmt.Expr)
	case *ast.ReturnStatement:
		if stmt.Value != nil {
			reg, _ := cg.ExpressionGen.GenerateExpression(stmt.Value)
			if reg != "a0" {
				cg.emit("    mv a0, %s", reg)
				if reg != "a0" && reg != "fa0" {
					cg.Registers.ReleaseRegister(reg)
				}
			}
		}
	case *ast.IfStatement:
		cg.BranchingGen.GenerateIfStatement(stmt)
	case *ast.WhileStatement:
		cg.LoopingGen.GenerateWhileStatement(stmt)
	case *ast.VarDeclaration:
		cg.AssignmentGen.GenerateVarDeclaration(stmt)
	case *ast.Block:
		for _, blockItem := range stmt.Items {
			cg.GenerateStatement(blockItem)
		}
	default:
		panic(fmt.Sprintf("unknown statement type: %T", stmt))
	}
}

func (cg *CodeGenerator) CalculateStackSize(block *ast.Block) int64 {
	var size int64
	for _, item := range block.Items {
		switch stmt := item.(type) {
		case *ast.VarDeclaration:
			if _, ok := stmt.Type.(*ast.ArrayType); ok {
				// Array: 8 bytes per element * array size
				symbol, found := cg.SymTable.Lookup("main." + stmt.Name)
				if !found {
					symbol, _ = cg.SymTable.Lookup(stmt.Name)
				}
				arraySize := symbol.ArraySize // From symbol table
				size += 8 * arraySize
			} else {
				// Scalar types: int, float, bool, char
				size += 8
			}
		case *ast.Block:
			size += cg.CalculateStackSize(stmt)
		}
	}
	return size
}
