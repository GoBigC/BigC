.text
main:
# If both values are in immediate range [-2048, 2047]
    li t0, 100     
    li t1, 201
    sub t1, t1, t0   # t1 = t1 - t0

# Print result
    li a7, 1
    mv a0, t1
    ecall
