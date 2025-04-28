package codegen

import (
	"BigCooker/pkg/syntax/ast"
	"fmt"
)

type FunctionGenerator struct {
	CodeGen *CodeGenerator
}

func NewFunctionGenerator(cg *CodeGenerator) *FunctionGenerator {
	return &FunctionGenerator{
		CodeGen: cg,
	}
}

func (fg *FunctionGenerator) calculateFrameSize(funcName string) int {
	fs := 16 // minimum 8(ra) + 8(s0)
	table := fg.CodeGen.SymTable

	funcSymbol, found := table.Lookup(funcName)
	if !found {
		panic(fmt.Sprintf("Function %s not found in symbol table", funcName))
	}

	n_param := len(funcSymbol.Parameters)
	if n_param > 8 {
		// we have 8 registers a0-a7 to store up to 8 params
		// but if we need more, we need to allocate additional stack space
		// the stack is in memory (not cpu like registers)
		// --> only allocate additional stack space if needed
		fs += (n_param - 8) * 8
	}

	funcScope := funcSymbol.Scope
	for _, symbol := range table.Symbols {
		if symbol.Scope.ValidFirstLine >= funcScope.ValidFirstLine &&
			symbol.Scope.ValidLastLine <= funcScope.ValidLastLine &&
			symbol.Name != funcName {
			// we counted this before so skip this symbol
			isParam := false
			for _, param := range funcSymbol.Parameters {
				if param.Name == symbol.Name {
					isParam = true
					break
				}
			}

			if !isParam {
				var size int

				if symbol.ArraySize > 0 {
					elemSize := 8
					size = int(symbol.ArraySize) * elemSize
				} else {
					// if not an array then its just 8
					size = 8
				}

				fs += size
			}
		}
	}

	// in case the framesize is not multiple of 16 --> round it up to nearest
	// see this: https://riscv.org/wp-content/uploads/2024/12/riscv-calling.pdf
	// tl;dr: "In the standard RISC-V calling convention, the stack grows downward
	// and the stack pointer is always kept 16-byte aligned."
	rem := fs % 16
	if rem != 0 {
		fs += 16 - (rem % 16)
	}

	return fs
}

func typeString(astType ast.Type) string {
	if p, ok := astType.(*ast.PrimitiveType); ok {
		return p.Name
	}
	if a, ok := astType.(*ast.ArrayType); ok {
		size := -1
		if lit, ok := a.Size.(*ast.IntegerLiteral); ok {
			size = int(lit.Value)
		}
		return fmt.Sprintf("%s[%d]", typeString(a.ElementType), size)
	}
	return "unknown"
}

func (fg *FunctionGenerator) GenerateFunctionDeclaration(funcDecl ast.FunctionDeclaration) {
	/*
		Convention:
		- registers have roles as described in CodeGenerator struct
		- stack frame size: framesize represents the total size of function's stack frame.
		because we are in riscv 64-bit mode:
			- 8 bytes per saved registers
			- 8 bytes per local variables
			- --> total should be a multiple of 16 bytes for proper alignment

		Prologue:
		- allocate stack space for local var
		- saved calle-save registers (agreement between functions: callee-saved registers s0-s11 must
		have the same values when a function returns as when it was called. we backup the values in
		these registers so that we can use it and restore values before returning)
		- set value of frame pointer s0/fp to access local var (stable, reliable reference point to access
		local var and params)

		Epilogue:
		- restore saved register in reverse order (to the values we backup in prologue)
		- restore frame pointer (because s0/fp itself is a callee-saved register)
		- deallocate stack frame
		- returns to the caller (`ret` instruction returns to address stored in `ra` register)
	*/
	cg := fg.CodeGen
	cg.CurrentFunction = funcDecl.Name
	funcName := funcDecl.Name
	fs := fg.calculateFrameSize(funcName)

	// function label
	cg.emit("%s:", funcName)

	// function prologue
	cg.emitComment("function prologue")
	cg.emit("	addi sp, sp, -%d", fs) // allocate stack space
	cg.emit("	sd ra, %d(sp)", fs-8)  // save return address (ra)
	cg.emit("	sd s0, %d(sp)", fs-16) // save frame pointer
	cg.emit("	addi s0, sp, %d", fs)  // set frame pointer to this function's scope

	// function parameters & local var
	cg.emitComment("setup parameters")
	funcSymbol, _ := cg.SymTable.Lookup(funcName)
	params := funcSymbol.Parameters

	for i, param := range params {
		if i < 8 { // first 8 in a0-a7
			cg.VarStackOffset[param.Name] = -1 // -1 means register (not stack)
			// ...do logic to load the param in a0-a7...
			cg.emitComment("param %s in register a%d", param.Name, i)
		} else { // needs stack
			// first memory is to backup ra
			// second memory is to backup s0
			offset := 16 + 8*(i-8)
			cg.VarStackOffset[param.Name] = offset
			cg.emitComment("param %s at offset %d(sp)", param.Name, offset)
		}
	}

	localVarOffset := 16
	if len(params) > 8 {
		localVarOffset += (len(params) - 8) * 8
	}

	funcScope := funcSymbol.Scope
	for _, symbol := range cg.SymTable.Symbols {
		if (symbol.Name != funcName) &&
			(typeString(symbol.Type) != "function") &&
			(symbol.Scope.ValidFirstLine >= funcScope.ValidFirstLine) &&
			(symbol.Scope.ValidLastLine <= funcScope.ValidLastLine) {
			isParam := false
			for _, param := range params {
				if param.Name == symbol.Name {
					isParam = true
					break
				}
			}

			if !isParam {
				var size int
				if /*symbol.ArraySize != nil &&*/ symbol.ArraySize > 0 {
					size = int(symbol.ArraySize) * 8
				} else {
					size = 8
				}

				cg.VarStackOffset[symbol.Name] = localVarOffset
				cg.emitComment("local var %s at offset %d(sp)", symbol.Name, localVarOffset)
				localVarOffset += size
			}
		}
	}

	// function body
	cg.emitComment("function body")
	if funcDecl.Body != nil {
		for _, stmt := range funcDecl.Body.Items {
			fg.GenerateBlockItem(stmt)
		}
	}

	// function epilogue
	cg.emitComment("function epilogue")
	cg.emit("	ld ra, %d(sp)", fs-8)  // restore ra (of caller func)
	cg.emit("	ld s0, %d(sp)", fs-16) // restore fp
	cg.emit("	addi sp, sp, %d", fs)  // deallocate stack frame pointer

	if funcName == "main" {
		cg.emit("	j _exit") // Special case for main
	} else {
		cg.emit("	ret") // return to caller func, whose address is in `ra`
	}
}

func (fg *FunctionGenerator) GenerateReturnStatement(stmt ast.ReturnStatement) {
	cg := fg.CodeGen

	if stmt.Value != nil {
		resultRegister := cg.ExpressionGen.GenerateExpression(stmt.Value)
		returnType := fg.getFunctionReturnType(cg.CurrentFunction)

		if resultRegister == "" {
			if returnType == "float" {
				resultRegister = "fa0"
			} else {
				resultRegister = "a0"
			}
		}

		if returnType == "float" && resultRegister != "fa0" {
			fg.CodeGen.emit("	fmv.d fa0, %s", resultRegister)
		} else if returnType != "float" && resultRegister != "a0" {
			fg.CodeGen.emit("	mv a0, %s", resultRegister)
		}

		if resultRegister != "a0" && resultRegister != "fa0" {
			fg.CodeGen.Registers.ReleaseRegister(resultRegister)
		}
	}
}

func (fg *FunctionGenerator) getFunctionReturnType(funcName string) string {
	if symbol, found := fg.CodeGen.SymTable.Lookup(funcName); found {
		if symbol.ReturnType != nil {
			if primitiveType, ok := symbol.ReturnType.(*ast.PrimitiveType); ok {
				return primitiveType.Name
			}
		}
	}
	return ""
}

func (fg *FunctionGenerator) GenerateBlockItem(stmt ast.BlockItem) {
	cg := fg.CodeGen
	switch s := stmt.(type) {
	case *ast.ExpressionStatement:
		cg.ExpressionGen.GenerateExpression(s.Expr)
	case *ast.ReturnStatement:
		fg.GenerateReturnStatement(*s)
	case *ast.IfStatement:
		cg.BranchingGen.GenerateIfStatement(s)
	case *ast.WhileStatement:
		cg.LoopingGen.GenerateWhileStatement(*s)
	case *ast.VarDeclaration:
		// it makes sense why this func does not take
		// argument but be careful this may cause problem
		cg.AssignmentGen.GenerateVarDeclaration(*s) // watch out
	// case *ast.Block:
	//     fg.GenerateBlock(*s)
	case *ast.FunctionDeclaration:
		panic("nested function declaration not allowed")
	default:
		panic(fmt.Sprintf("unknown statement type in function body: %T", stmt))
	}
}

// func (fg *FunctionGenerator) GenerateBlock(block ast.Block) {
//     for _, item := range block.Items {
//         fg.GenerateBlockItem(item)
//     }
// }
