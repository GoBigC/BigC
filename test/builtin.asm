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

# ----- begin hardcode function -----
main: 
    li a0, 99      # Value to print
    li a7, 1       # Print integer
    ecall
    
    li a0, 10      # Newline
    li a7, 11      # Print character
    ecall
	# Function prologue
    addi sp, sp, -16
    sd ra, 8(sp)
    sd s0, 0(sp)
    addi s0, sp, 16
    
    # Load 42 into a0
    li a0, 42
    # Call _printInt
    jal _printInt
    
    # Load 6.9 into fa0
    li t0, 0x40dc0000  # IEEE 754 representation of 6.9 (approximation)
    fmv.w.x fa0, t0
    # Call _printFloat
    jal _printFloat
    
    # Function epilogue
    ld ra, 8(sp)
    ld s0, 0(sp)
    addi sp, sp, 16

# ----- end hardcode function -----
j _exit
