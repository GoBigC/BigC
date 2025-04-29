.data
float_imm_0: .double 3.141590
literalFloat: .double 6.280000
divFloat: .double 0.000000
mulFloat: .double 0.000000
subFloat: .double 0.000000
addFloat: .double 0.000000
double_2: .double 2.500000
double_1: .double 3.140000
y: .double 5.500000
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
	la t0, double_1
	fld ft0, 0(t0)
	la t1, double_2
	fld ft1, 0(t1)
	fadd.d fa0, ft0, ft1
	la t2, addFloat
	fsd fa0, 0(t2)
    la t2, addFloat
    fld ft2, 0(t2)
    fmv.d fa0, ft2
    jal _printFloat
	la t2, double_1
	fld ft2, 0(t2)
	la t3, double_2
	fld ft3, 0(t3)
	fsub.d fa0, ft2, ft3
	la t4, subFloat
	fsd fa0, 0(t4)
    la t4, subFloat
    fld ft4, 0(t4)
    fmv.d fa0, ft4
    jal _printFloat
	la t4, double_1
	fld ft4, 0(t4)
	la t5, double_2
	fld ft5, 0(t5)
	fmul.d fa0, ft4, ft5
	la t6, mulFloat
	fsd fa0, 0(t6)
    la t6, mulFloat
    fld ft6, 0(t6)
    fmv.d fa0, ft6
    jal _printFloat
	la t6, double_1
	fld ft6, 0(t6)
	la a2, double_2
	fld ft7, 0(a2)
	fdiv.d fa0, ft6, ft7
	la a3, divFloat
	fsd fa0, 0(a3)
    la a3, divFloat
    fld ft8, 0(a3)
    fmv.d fa0, ft8
    jal _printFloat
    la a3, literalFloat
    fld ft8, 0(a3)
    fmv.d fa0, ft8
    jal _printFloat
	la a3, float_imm_0
	fld ft8, 0(a3)
    fmv.d fa0, ft8
    jal _printFloat
    li a3, 0
    mv a0, a3
	li a0, 0
	j _exit
