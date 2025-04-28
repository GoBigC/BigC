.data
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
print1:
# function prologue
	addi sp, sp, -80
	sd ra, 72(sp)
	sd s0, 64(sp)
	addi s0, sp, 80
# setup parameters
# param x in register a0
# function body --- hardcode 
	jal _printInt
# function epilogue
	ld ra, 72(sp)
	ld s0, 64(sp)
	addi sp, sp, 80
	ret
print2:
# function prologue
	addi sp, sp, -80
	sd ra, 72(sp)
	sd s0, 64(sp)
	addi s0, sp, 80
# setup parameters
# param x in register a0
# function body --- hardcode 
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
# local var x at offset 24(sp)
# function body --- hardcode 
	li a0, 42
	jal print1

	li t0, 0x40DC0000
	fmv.w.x fa0, t0
	jal print2
# function epilogue
	ld ra, 88(sp)
	ld s0, 80(sp)
	addi sp, sp, 96
	# ret # problem in this line: maybe it doesnt have a place to return? 
j _exit
