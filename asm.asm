.data
literalInt: .dword 4
divInt: .dword 0
mulInt: .dword 0
subInt: .dword 0
addInt: .dword 0
x: .dword 10
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
	li t0, 2050
	addi a0, t0, 1
	la t1, addInt
	sd a0, 0(t1)
    la t1, addInt
    ld t2, 0(t1)
    mv a0, t2
    jal _printInt
	li t1, 5
	li t2, 2
	sub a0, t1, t2
	la t3, subInt
	sd a0, 0(t3)
    la t3, subInt
    ld t4, 0(t3)
    mv a0, t4
    jal _printInt
	li t3, 3
	li t4, 4
	mul a0, t3, t4
	la t5, mulInt
	sd a0, 0(t5)
    la t5, mulInt
    ld t6, 0(t5)
    mv a0, t6
    jal _printInt
	li t5, 8
	li t6, 2
	div a0, t5, t6
	la a2, divInt
	sd a0, 0(a2)
    la a2, divInt
    ld a3, 0(a2)
    mv a0, a3
    jal _printInt
    la a2, literalInt
    ld a3, 0(a2)
    mv a0, a3
    jal _printInt
    li a2, 0
    mv a0, a2
	li a0, 0
	j _exit
