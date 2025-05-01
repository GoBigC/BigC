.data
n1: .dword 0
n: .dword 0
k: .dword 0
z: .dword 4
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
    li t0, 69
	la t1, a
    li t2, 0
	li t3, 8
	mul t3, t2, t3
	add t4, t1, t3
	sd t0, 0(t4)
	la t0, a
    li t1, 0
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	ld t0, 0(t3)
    mv a0, t0
    jal _printInt
    la t0, x
    ld t1, 0(t0)
    la t0, z
    ld t2, 0(t0)
	sub t0, t1, t2
	la t1, k
	sd t0, 0(t1)
    li t0, 420
	la t1, a
    la t2, k
    ld t3, 0(t2)
	li t2, 8
	mul t2, t3, t2
	add t4, t1, t2
	sd t0, 0(t4)
	la t0, a
    la t1, k
    ld t2, 0(t1)
	li t1, 8
	mul t1, t2, t1
	add t3, t0, t1
	ld t0, 0(t3)
	la t1, n
	sd t0, 0(t1)
    la t0, n
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    li t0, 425
	la t1, a
    la t2, k
    ld t3, 0(t2)
	li t2, 8
	mul t2, t3, t2
	add t4, t1, t2
	sd t0, 0(t4)
    la t0, n
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    li t0, 0
    mv a0, t0
	li a0, 0
	j _exit
