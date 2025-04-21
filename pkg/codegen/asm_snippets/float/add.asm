.data
float.1: .double 6.9  # we assume all floats used in BigC to be float64
float.2: .double 4.2

.text
main:
    la t0, float.1     
    fld ft0, 0(t0) # t0 = a
    
    la t1, float.2
    fld ft1, 0(t1) # t1 = b
    
    fadd.d ft1, ft1, ft0   # t1 = t1 + t0

# Print result
    li a7, 3
    fmv.d fa0, ft1
    ecall
