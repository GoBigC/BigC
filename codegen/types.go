// codegen/types.go
package codegen

func (cg *CodeGenerator) typeToDirective(typ string) string {
    switch typ {
    case "int":
        return ".word"
    case "float":
        return ".float" // Requires RV32F extension
    case "bool":
        return ".word"  // Treat as 1-byte value
    case "char":
        return ".byte"
    case "void":
        return "" // No storage, for functions
    default:
        panic("Unknown type: " + typ)
    }
}