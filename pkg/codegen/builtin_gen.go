package codegen 

func (cg *CodeGenerator) GenerateAllBuiltinFunctions(){
	cg.Generate_exit()

	cg.Generate_printInt()
	cg.Generate_printFloat()
	cg.Generate_printChar()
	cg.Generate_printBool()
	cg.Generate_printString()
}

const (
	PRTINT = 1 
	PRTFLT = 2
	PRTDBL = 3 
	PRTSTR = 4
	PRTCHR = 11 

	RDINT = 5
	RDFLT = 6 
	RDDBL = 7
	RDSTR = 8 
	RDCHR = 12

	EXIT = 10 
)

func (cg *CodeGenerator) generateBuiltinPrologue() {
	cg.emit("	addi sp, sp, -16")
	cg.emit("	sd ra, 8(sp)")
	cg.emit("	sd s0, 0(sp)")
	cg.emit("	addi s0, sp, 16")
}

func (cg *CodeGenerator) generateBuiltinEpilogue() {
	cg.emit("	ld ra, 8(sp)")
	cg.emit("	ld s0, 0(sp)")
	cg.emit("	addi sp, sp, 16")
	cg.emit("	ret")
}

func (cg *CodeGenerator) generatePrintNewline() {
	cg.emit("	li a0, 10")
	cg.emit("	li a7, %d", PRTCHR)
	cg.emit("	ecall")	
}

func (cg *CodeGenerator) Generate_exit() {
	cg.emit("_exit:")
	cg.emit("	li a7, %d", EXIT)
	cg.emit("	ecall")
}

func (cg *CodeGenerator) Generate_printInt() {
	/*
	Print integer representation of whatever is in `a0`
	*/
	cg.emit("_printInt:")
	cg.generateBuiltinPrologue()

	cg.emit("	li a7, %d", PRTINT)
	cg.emit("	ecall")

	cg.generatePrintNewline()
	cg.generateBuiltinEpilogue()
}

func (cg *CodeGenerator) Generate_printFloat() {
	/*
	Print float64 representation of whatever is in `fa0`
	*/
	cg.emit("_printFloat:")
	cg.generateBuiltinPrologue()

	cg.emit("	li a7, %d", PRTFLT)
	cg.emit("	ecall")

	cg.generatePrintNewline()
	cg.generateBuiltinEpilogue()
}

func (cg *CodeGenerator) Generate_printChar() {
	/*
	Print char representation of whatever is in `a0`
	*/
	cg.emit("_printChar:")
	cg.generateBuiltinPrologue()

	cg.emit("	li a7, %d", PRTCHR)
	cg.emit("	ecall")

	cg.generatePrintNewline()
	cg.generateBuiltinEpilogue()
}

func (cg *CodeGenerator) Generate_printBool() {
	cg.emit("_printBool:")
	cg.emit("	j printInt")
} 

// i think its just this but not very sure 
func (cg *CodeGenerator) Generate_printString() {
	/*
	Print the string whose address is stored in `a0`
	*/
    cg.emit("_printString:")
	cg.generateBuiltinPrologue()

	cg.emit("	li a7, %d", PRTSTR)
	cg.emit("	ecall")

	cg.generateBuiltinEpilogue()
}

// not sure how to read lol 

// func (cg *CodeGenerator) Generate_readInt() {}
// func (cg *CodeGenerator) Generate_readFloat() {}
// func (cg *CodeGenerator) Generate_readChar() {}
// func (cg *CodeGenerator) Generate_readBool() {}
// func (cg *CodeGenerator) Generate_readString() {}