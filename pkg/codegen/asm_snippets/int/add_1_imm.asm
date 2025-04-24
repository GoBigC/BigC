.text
main:
# If at 1 value is in immediate range [-2048, 2047]
    li t0, 200     # Load immediate value 200 into t0
    addi t0, t0, 10   # t1 = t0 + 10

# Print result
    li a7, 1
    mv a0, t0
    ecall
