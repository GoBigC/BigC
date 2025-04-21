.data
x: .word 10
y: .word 8000
.text
main:
# If not immediate (in range [-2048, 2047]
    la t0, x      # t0 = address(x)
    lw t0, 0(t0)	# t0 = x
    la t1, y      # t1 = address(y)
    lw t1, 0(t1)	# t1 = y
    add t1, t1, t0   # t1 = t0 + 200

# Print result
    li a7, 1
    mv a0, t1
    ecall