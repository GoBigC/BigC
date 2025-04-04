// codegen/registers.go
package codegen

type RegisterManager struct {
    used map[string]bool
}

func NewRegisterManager() *RegisterManager {
    return &RegisterManager{
        used: make(map[string]bool),
    }
}

func (rm *RegisterManager) Allocate() string {
    for i := 0; i <= 6; i++ {
        reg := "t" + string(rune('0'+i))
        if !rm.used[reg] {
            rm.used[reg] = true
            return reg
        }
    }
    panic("Out of registers") // Add stack spilling later
}

func (rm *RegisterManager) Free(reg string) {
    delete(rm.used, reg)
}