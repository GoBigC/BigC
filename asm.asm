.data
total: .double 0.000000
float_imm_18: .double 15.250000
float_imm_17: .double 20.750000
float_imm_16: .double 10.500000
prices: .space 24
float_imm_15: .double 6.699000
float_imm_14: .double 3.141200
c: .space 160
b: .space 128
ff: .dword 0
n1: .dword 0
n: .dword 0
k: .dword 0
z: .dword 4
a: .space 120
comparisonNEQ: .dword 0
comparisonEQ: .dword 0
comparisonGT: .dword 0
comparisonLT: .dword 0
notOperation: .dword 0
orOperation: .dword 0
andOperation: .dword 0
falseValue: .dword 0
trueValue: .dword 1
literalChar: .dword 65
float_imm_13: .double 3.141590
float_imm_12: .double 2.200000
mixDivFloat: .double 0.000000
float_imm_11: .double 1.100000
mixSubFloat: .double 0.000000
float_imm_10: .double 2.000000
mixMulFloat: .double 0.000000
float_imm_9: .double 2.300000
mixedAddFloat: .double 0.000000
float_imm_8: .double 3.140000
minusFloat: .double 0.000000
literalFloat: .double 6.280000
float_imm_7: .double 3.000000
float_imm_6: .double 15.000000
divFloat: .double 0.000000
float_imm_5: .double 3.500000
float_imm_4: .double 2.000000
mulFloat: .double 0.000000
float_imm_3: .double 4.200000
float_imm_2: .double 10.500000
subFloat: .double 0.000000
float_imm_1: .double 2.500000
float_imm_0: .double 3.140000
addFloat: .double 0.000000
mixDivInt: .dword 0
mixSubInt: .dword 0
mixMulInt: .dword 0
mixedAddInt: .dword 0
minusInt: .dword 0
literalInt: .dword 4
divInt: .dword 0
mulInt: .dword 0
subInt: .dword 0
addInt: .dword 0
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
    li t0, 1
    li t1, 2050
	add t2, t0, t1
	la t0, addInt
	sd t2, 0(t0)
    la t0, addInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    li t0, 5
    li t1, 2
	sub t2, t0, t1
	la t0, subInt
	sd t2, 0(t0)
    la t0, subInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    li t0, 3
    li t1, 4
	mul t2, t0, t1
	la t0, mulInt
	sd t2, 0(t0)
    la t0, mulInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    li t0, 8
    li t1, 2
	div t2, t0, t1
	la t0, divInt
	sd t2, 0(t0)
    la t0, divInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    la t0, literalInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    li t0, 5
    neg t1, t0
	la t0, minusInt
	sd t1, 0(t0)
    la t0, minusInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    la t0, x
    ld t1, 0(t0)
    li t0, 12
	add t2, t1, t0
	la t0, mixedAddInt
	sd t2, 0(t0)
    la t0, mixedAddInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    la t0, x
    ld t1, 0(t0)
    li t0, 1
	mul t2, t1, t0
	la t0, mixMulInt
	sd t2, 0(t0)
    la t0, mixMulInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    la t0, x
    ld t1, 0(t0)
    li t0, 1
	sub t2, t1, t0
	la t0, mixSubInt
	sd t2, 0(t0)
    la t0, mixSubInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    la t0, x
    ld t1, 0(t0)
    li t0, 2
	div t2, t1, t0
	la t0, mixDivInt
	sd t2, 0(t0)
    la t0, mixDivInt
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    li t0, 13
    mv a0, t0
    jal _printInt
	la t0, float_imm_0
	fld ft0, 0(t0)
	la t0, float_imm_1
	fld ft1, 0(t0)
	fadd.d ft2, ft0, ft1
	la t0, addFloat
	fsd ft2, 0(t0)
    la t0, addFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
	la t0, float_imm_2
	fld ft0, 0(t0)
	la t0, float_imm_3
	fld ft1, 0(t0)
	fsub.d ft2, ft0, ft1
	la t0, subFloat
	fsd ft2, 0(t0)
    la t0, subFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
	la t0, float_imm_4
	fld ft0, 0(t0)
	la t0, float_imm_5
	fld ft1, 0(t0)
	fmul.d ft2, ft0, ft1
	la t0, mulFloat
	fsd ft2, 0(t0)
    la t0, mulFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
	la t0, float_imm_6
	fld ft0, 0(t0)
	la t0, float_imm_7
	fld ft1, 0(t0)
	fdiv.d ft2, ft0, ft1
	la t0, divFloat
	fsd ft2, 0(t0)
    la t0, divFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
    la t0, literalFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
	la t0, float_imm_8
	fld ft0, 0(t0)
    fneg.d ft1, ft0
	la t0, minusFloat
	fsd ft1, 0(t0)
    la t0, minusFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
    la t0, y
    fld ft0, 0(t0)
	la t0, float_imm_9
	fld ft1, 0(t0)
	fadd.d ft2, ft0, ft1
	la t0, mixedAddFloat
	fsd ft2, 0(t0)
    la t0, mixedAddFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
    la t0, y
    fld ft0, 0(t0)
	la t0, float_imm_10
	fld ft1, 0(t0)
	fmul.d ft2, ft0, ft1
	la t0, mixMulFloat
	fsd ft2, 0(t0)
    la t0, mixMulFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
    la t0, y
    fld ft0, 0(t0)
	la t0, float_imm_11
	fld ft1, 0(t0)
	fsub.d ft2, ft0, ft1
	la t0, mixSubFloat
	fsd ft2, 0(t0)
    la t0, mixSubFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
    la t0, y
    fld ft0, 0(t0)
	la t0, float_imm_12
	fld ft1, 0(t0)
	fdiv.d ft2, ft0, ft1
	la t0, mixDivFloat
	fsd ft2, 0(t0)
    la t0, mixDivFloat
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
	la t0, float_imm_13
	fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
    la t0, literalChar
    ld t1, 0(t0)
    mv a0, t1
    jal _printChar
    li t0, 66
    mv a0, t0
    jal _printChar
    la t0, trueValue
    ld t1, 0(t0)
    mv a0, t1
    jal _printBool
    la t0, falseValue
    ld t1, 0(t0)
    mv a0, t1
    jal _printBool
    li t0, 1
    li t1, 0
    and t2, t0, t1
	la t0, andOperation
	sd t2, 0(t0)
    la t0, andOperation
    ld t1, 0(t0)
    mv a0, t1
    jal _printBool
    li t0, 1
    li t1, 0
    or t2, t0, t1
	la t0, orOperation
	sd t2, 0(t0)
    la t0, orOperation
    ld t1, 0(t0)
    mv a0, t1
    jal _printBool
    la t0, trueValue
    ld t1, 0(t0)
    seqz t0, t1
	la t1, notOperation
	sd t0, 0(t1)
    la t0, notOperation
    ld t1, 0(t0)
    mv a0, t1
    jal _printBool
    li t0, 5
    li t1, 10
    slt t2, t0, t1
	la t0, comparisonLT
	sd t2, 0(t0)
    la t0, comparisonLT
    ld t1, 0(t0)
    mv a0, t1
    jal _printBool
    li t0, 20
    li t1, 15
    slt t2, t1, t0
	la t0, comparisonGT
	sd t2, 0(t0)
    la t0, comparisonGT
    ld t1, 0(t0)
    mv a0, t1
    jal _printBool
    li t0, 7
    li t1, 7
    sub t2, t0, t1
    seqz t2, t2
	la t0, comparisonEQ
	sd t2, 0(t0)
    la t0, comparisonEQ
    ld t1, 0(t0)
    mv a0, t1
    jal _printBool
    li t0, 8
    li t1, 9
    sub t2, t0, t1
    snez t2, t2
	la t0, comparisonNEQ
	sd t2, 0(t0)
    la t0, comparisonNEQ
    ld t1, 0(t0)
    mv a0, t1
    jal _printBool
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
	la t0, a
    la t1, k
    ld t2, 0(t1)
	li t1, 8
	mul t1, t2, t1
	add t3, t0, t1
	ld t0, 0(t3)
    mv a0, t0
    jal _printInt
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
	la t0, a
    la t1, k
    ld t2, 0(t1)
	li t1, 8
	mul t1, t2, t1
	add t3, t0, t1
	ld t0, 0(t3)
    la t1, n1
    sd t0, 0(t1)
    la t0, n1
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
    li t0, 1
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
    li t0, 1
	la t1, a
    li t2, 1
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
	la t1, a
    li t2, 1
	li t3, 8
	mul t3, t2, t3
	add t4, t1, t3
	ld t1, 0(t4)
	add t2, t0, t1
	la t0, ff
	sd t2, 0(t0)
    la t0, n
    ld t1, 0(t0)
    mv a0, t1
    jal _printInt
	la t0, float_imm_14
	fld ft0, 0(t0)
	la t0, c
    li t1, 0
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fsd ft0, 0(t3)
	la t0, float_imm_15
	fld ft0, 0(t0)
	la t0, c
    li t1, 1
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fsd ft0, 0(t3)
	la t0, c
    li t1, 1
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fld ft0, 0(t3)
    fmv.d fa0, ft0
    jal _printFloat
	la t0, c
    li t1, 0
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fld ft0, 0(t3)
    fmv.d fa0, ft0
    jal _printFloat
	la t0, float_imm_16
	fld ft0, 0(t0)
	la t0, prices
    li t1, 0
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fsd ft0, 0(t3)
	la t0, float_imm_17
	fld ft0, 0(t0)
	la t0, prices
    li t1, 1
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fsd ft0, 0(t3)
	la t0, float_imm_18
	fld ft0, 0(t0)
	la t0, prices
    li t1, 2
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fsd ft0, 0(t3)
	la t0, prices
    li t1, 0
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fld ft0, 0(t3)
    fmv.d fa0, ft0
    jal _printFloat
	la t0, prices
    li t1, 1
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fld ft0, 0(t3)
    fmv.d fa0, ft0
    jal _printFloat
	la t0, prices
    li t1, 0
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fld ft0, 0(t3)
	la t0, prices
    li t1, 1
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fld ft1, 0(t3)
	fadd.d ft2, ft0, ft1
	la t0, prices
    li t1, 2
	li t2, 8
	mul t2, t1, t2
	add t3, t0, t2
	fld ft0, 0(t3)
	fadd.d ft1, ft2, ft0
	la t0, total
	fsd ft1, 0(t0)
    la t0, total
    fld ft0, 0(t0)
    fmv.d fa0, ft0
    jal _printFloat
    li t0, 0
    mv a0, t0
	li a0, 0
	j _exit
