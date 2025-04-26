.text
main:
# If at 1 value is in immediate range [-2048, 2047]
    li t0, 4294967296
    li t1, 8589934592   # Load immediate value 10 into t1
    mul a0, t1, t0

# Print result
    li a7, 1
    ecall
