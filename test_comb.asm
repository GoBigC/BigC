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
printInt:
# function prologue
	addi sp, sp, -80
	sd ra, 72(sp)
	sd s0, 64(sp)
	addi s0, sp, 80
# setup parameters
# param x in register a0
# local var addInt at offset 16(sp)
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
# local var addInt at offset 16(sp)
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
# local var x at offset 24(sp)
# local var addInt at offset 32(sp)
# function body
li t0, 2050
addi a0, t0, 1
	sd a0, 32(sp)
	ld t1, 32(sp)
    mv a0, t1
	jal printInt
li t1, 0
addi a0, t1, 0
# function epilogue
	ld ra, 88(sp)
	ld s0, 80(sp)
	addi sp, sp, 96
	j _exit
