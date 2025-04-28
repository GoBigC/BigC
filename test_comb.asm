.data
double_2: .double 6.900000
double_1: .double 4.200000
.text
j main
_exit:
	li a7, 10
	ecall
_printInt:
	addi sp, sp, -16
	sd ra, 8(sp)
	sd s0, 0(sp)
	addi s0, sp, 16
	li a7, 1
	ecall
	li a0, 10
	li a7, 11
	ecall
	ld ra, 8(sp)
	ld s0, 0(sp)
	addi sp, sp, 16
	ret
_printFloat:
	addi sp, sp, -16
	sd ra, 8(sp)
	sd s0, 0(sp)
	addi s0, sp, 16
	li a7, 2
	ecall
	li a0, 10
	li a7, 11
	ecall
	ld ra, 8(sp)
	ld s0, 0(sp)
	addi sp, sp, 16
	ret
_printChar:
	addi sp, sp, -16
	sd ra, 8(sp)
	sd s0, 0(sp)
	addi s0, sp, 16
	li a7, 11
	ecall
	li a0, 10
	li a7, 11
	ecall
	ld ra, 8(sp)
	ld s0, 0(sp)
	addi sp, sp, 16
	ret
_printBool:
	j _printInt
_printString:
	addi sp, sp, -16
	sd ra, 8(sp)
	sd s0, 0(sp)
	addi s0, sp, 16
	li a7, 4
	ecall
	ld ra, 8(sp)
	ld s0, 0(sp)
	addi sp, sp, 16
	ret
printInt:
# function prologue
	addi sp, sp, -80
	sd ra, 72(sp)
	sd s0, 64(sp)
	addi s0, sp, 80
# setup parameters
# param x in register a0
# local var addFloat at offset 16(sp)
# function body
	jal _printInt
    li t0, 0
    mv a0, t0
# function epilogue
	ld ra, 72(sp)
	ld s0, 64(sp)
	addi sp, sp, 80
	ret
printFloat:
# function prologue
	addi sp, sp, -80
	sd ra, 72(sp)
	sd s0, 64(sp)
	addi s0, sp, 80
# setup parameters
# param x in register a0
# local var addFloat at offset 16(sp)
# function body
	jal _printFloat
# function epilogue
	ld ra, 72(sp)
	ld s0, 64(sp)
	addi sp, sp, 80
	ret
main:
# function prologue
	addi sp, sp, -96
	sd ra, 88(sp)
	sd s0, 80(sp)
	addi s0, sp, 96
# setup parameters
# local var x at offset 16(sp)
# local var addFloat at offset 24(sp)
# local var x at offset 32(sp)
# function body
la ft0, double_1
fld ft0, 0(ft0)
la ft1, double_2
fld ft1, 0(ft1)
fadd.d fa0, ft0, ft1
	fs fa0, 24(sp)
	fld ft2, 24(sp)
    fmv.d fa0, ft2
	jal printFloat
li t0, 0
addi a0, t0, 0
# function epilogue
	ld ra, 88(sp)
	ld s0, 80(sp)
	addi sp, sp, 96
	j _exit
