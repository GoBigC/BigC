.text
main:
# If at 1 value is in immediate range [-2048, 2047]
    li t0, 30
    li t1, 50   # Load immediate value 10 into t1
    #div a0, t1, t0
    rem a0, t1, t0
# Print result
    li a7, 1
    ecall
