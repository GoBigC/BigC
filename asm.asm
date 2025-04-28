.data
addFloat: .double 6.900000
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
main:
    la t0, addFloat
# ld from GenerateIdentifier, case !isFloatVar
    ld t1, 0(t0)
    jal _printFloat
	li t0, 0
	addi a0, t0, 0
	li a0, 0
	j _exit
