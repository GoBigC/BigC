.data
m: .dword 0
t: .dword 0
z: .dword 0
float_imm_0: .double 5.500000
y: .double 0.000000
x: .dword 0
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
    li t0, 10
    la t1, x
    sd t0, 0(t1)
	la t0, float_imm_0
	fld ft0, 0(t0)
    la t0, y
    fsd ft0, 0(t0)
    li t0, 1
    la t1, z
    sd t0, 0(t1)
    li t0, 2
    la t1, t
    sd t0, 0(t1)
    li t0, 0
    la t1, m
    sd t0, 0(t1)
    li t0, 1
    li t1, 2050
	add t2, t0, t1
    sd t2, -8(sp)
    ld t0, -8(sp)
    mv a0, t0
    jal _printInt
    li t0, 5
    li t1, 2
	sub t2, t0, t1
    sd t2, -16(sp)
    ld t0, -16(sp)
    mv a0, t0
    jal _printInt
    li t0, 3
    li t1, 4
	mul t2, t0, t1
    sd t2, -24(sp)
    ld t0, -24(sp)
    mv a0, t0
    jal _printInt
    li t0, 8
    li t1, 2
	div t2, t0, t1
    sd t2, -32(sp)
    ld t0, -32(sp)
    mv a0, t0
    jal _printInt
    li t0, 4
    sd t0, -40(sp)
    ld t0, -40(sp)
    mv a0, t0
    jal _printInt
    li t0, 5
    neg t1, t0
    sd t1, -48(sp)
    ld t0, -48(sp)
    mv a0, t0
    jal _printInt
    li t0, 13
    mv a0, t0
    jal _printInt
    li t0, 0
    mv a0, t0
	li a0, 0
	j _exit
