.data
a: .space 120
y: .double 5.500000
x: .dword 15
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
	li a7, 3
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
	la t0, a
    li t1, 0
	li t2, 8
	mul t2, t1, t2
	add t3, t2, t0
    li t0, 69
	sd t0, 0(t3)
	la t0, a
    li t1, 0
	li t2, 8
	mul t2, t1, t2
	add t3, t2, t0
	ld t0, 0(t3)
    mv a0, t0
    jal _printInt
    li t0, 0
    mv a0, t0
	li a0, 0
	j _exit
