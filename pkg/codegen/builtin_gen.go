package codegen

func (cg *CodeGenerator) GenerateAllBuiltinFunctions() {
	cg.Generate_exit()

	cg.Generate_printInt()
	cg.Generate_printFloat()
	cg.Generate_printChar()
	cg.Generate_printBool()
	cg.Generate_printString()

	cg.Generate_readInt()
	cg.Generate_readFloat()
	cg.Generate_readChar()
	cg.Generate_readBool()
	cg.Generate_readString()
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

	cg.emit("	li a7, %d", PRTDBL)
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
	cg.emit("	j _printInt")
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
// Generate_readInt reads an integer from input and returns it in a0
func (cg *CodeGenerator) Generate_readInt() {
	cg.emit("_readInt:")
	cg.generateBuiltinPrologue()

	cg.emit("    li a7, %d", RDINT) // Syscall 5: Read integer
	cg.emit("    ecall")            // Result in a0

	cg.generateBuiltinEpilogue()
}

// Generate_readFloat reads a double-precision float from input and returns it in fa0
func (cg *CodeGenerator) Generate_readFloat() {
	cg.emit("_readFloat:")
	cg.generateBuiltinPrologue()

	cg.emit("    li a7, %d", RDDBL) // Syscall 7: Read double
	cg.emit("    ecall")            // Result in fa0

	cg.generateBuiltinEpilogue()
}

// Generate_readChar reads a single character from input and returns it in a0
func (cg *CodeGenerator) Generate_readChar() {
	cg.emit("_readChar:")
	cg.generateBuiltinPrologue()

	cg.emit("    li a7, %d", RDCHR) // Syscall 12: Read character
	cg.emit("    ecall")            // Result in a0

	cg.generateBuiltinEpilogue()
}

// Generate_readBool reads an integer (0 or 1) from input and returns it in a0
func (cg *CodeGenerator) Generate_readBool() {
	cg.emit("_readBool:")
	cg.emit("    j _readInt") // Reuse _readInt, as bool is treated as 0/1
}

// Generate_readString reads a string into a buffer (address in a0) and returns the number of bytes read in a0
func (cg *CodeGenerator) Generate_readString() {
	cg.emit("_readString:")
	cg.generateBuiltinPrologue()

	// a0 contains the buffer address (caller-provided)
	// a1 should contain the maximum buffer size; assume caller sets it
	cg.emit("    li a7, %d", RDSTR) // Syscall 8: Read string
	cg.emit("    ecall")            // Reads into buffer at a0, returns bytes read in a0

	cg.generateBuiltinEpilogue()
}
