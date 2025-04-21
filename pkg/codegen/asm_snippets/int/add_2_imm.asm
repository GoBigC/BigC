.text
main:
# If both values are in immediate range [-2048, 2047]
    li t0, 100     
    addi t1, t0, 11   # t1 = t0 + 11

# Print result
    li a7, 1
    mv a0, t1
    ecall
