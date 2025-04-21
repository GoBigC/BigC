.data
x: .word 8000 # Store the larger value inside the word
.text
main:
# If at 1 value is in immediate range [-2048, 2047]
    la t0, x     
    lw t0, 0(t0) # t0 = 200 
    li t1, 10   # Load immediate value 10 into t1
    sub t1, t1, t0   # t1 = t0 + 10

# Print result
    li a7, 1
    mv a0, t1
    ecall
