package table

import (
	"BigCooker/pkg/syntax/ast"
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

type SymbolTable struct {
    Symbols   map[string]Symbol
}

type Symbol struct {
    Name       string
    Type       ast.Type // Use ast.Type for richer type info (e.g., ArrayType.ElementType)
    Scope      ScopeInfo
    ArraySize  int64
    Value      any
    Parameters []ast.Parameter // For functions
    ReturnType ast.Type        // For functions
}

type ScopeInfo struct {
    ValidFirstLine int
    ValidLastLine  int
}

// Add helper methods
func NewSymbolTable() *SymbolTable {
    table := &SymbolTable{
        Symbols:   make(map[string]Symbol),
    }
    table.RegisterBuiltinFunctions()
    return table
}

func (symTable *SymbolTable) Define(name string, symbol Symbol) {
    symTable.Symbols[name] = symbol
}

func (symTable *SymbolTable) Lookup(name string) (Symbol, bool) {
    if sym, ok := symTable.Symbols[name]; ok {
        return sym, true
    }
    return Symbol{}, false
}


func (symTable *SymbolTable) PrintTable() {

    // Set up tabwriter for aligned columns
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
    fmt.Fprintln(w, "Name\tType\tScope\tArraySize\tValue\tParameters\tReturnType")
    fmt.Fprintln(w, "----\t----\t-----\t---------\t-----\t----------\t----------")

    // Print each symbol as a row
    for _, sym := range symTable.Symbols {
        name := sym.Name
        typ := typeString(sym.Type)
        scope := fmt.Sprintf("%d-%d", sym.Scope.ValidFirstLine, sym.Scope.ValidLastLine)
        arraySize := "-"
        if sym.ArraySize > 0 {
            arraySize = fmt.Sprintf("%d", sym.ArraySize)
        }
        value := "-"
        if sym.Value != nil {
            value = fmt.Sprintf("%v", sym.Value)
        }
        params := "-"
        if len(sym.Parameters) > 0 {
            params = "["
            for i, p := range sym.Parameters {
                if i > 0 {
                    params += ", "
                }
                params += fmt.Sprintf("%s: %s", p.Name, typeString(p.Type))
            }
            params += "]"
        }
        retType := "-"
        if sym.ReturnType != nil {
            retType = typeString(sym.ReturnType)
        }
        fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", name, typ, scope, arraySize, value, params, retType)
    }
    fmt.Fprintln(w, "-----------------------------------------------------------------")

    w.Flush()
    fmt.Println()
}

// typeString converts an ast.Type to a string for printing.
func typeString(astType ast.Type) string {
    if p, ok := astType.(*ast.PrimitiveType); ok {
        return p.Name
    }
    if a, ok := astType.(*ast.ArrayType); ok {
        size := -1
        if lit, ok := a.Size.(*ast.IntegerLiteral); ok {
            size = int(lit.Value)
        }
        return fmt.Sprintf("%s[%d]", typeString(a.ElementType), size)
    }
    return "unknown"
}

func (symTable *SymbolTable) RegisterBuiltinFunctions(){
    lastLine := math.MaxInt

    intType := &ast.PrimitiveType{Name: "int"}
    floatType := &ast.PrimitiveType{Name: "float"}
    charType := &ast.PrimitiveType{Name: "char"}
    boolType := &ast.PrimitiveType{Name: "bool"}
    
    // TODO: register functions for reads{int, float, char, bool, string} && printString, readString

    symTable.Define("_printInt", Symbol{
        Name:       "_printInt",
        Type:       &ast.PrimitiveType{Name: "function"},
        Scope:      ScopeInfo{ValidFirstLine: 1, ValidLastLine: lastLine}, 
        Parameters: []ast.Parameter{
            {Name: "value", Type: intType},
        },
        ReturnType: nil,
    })
    
    symTable.Define("_printFloat", Symbol{
        Name:       "_printFloat",
        Type:       &ast.PrimitiveType{Name: "function"},
        Scope:      ScopeInfo{ValidFirstLine: 1, ValidLastLine: lastLine}, 
        Parameters: []ast.Parameter{
            {Name: "value", Type: floatType},
        },
        ReturnType: nil,
    })
    
    symTable.Define("_printChar", Symbol{
        Name:       "_printChar",
        Type:       &ast.PrimitiveType{Name: "function"},
        Scope:      ScopeInfo{ValidFirstLine: 1, ValidLastLine: lastLine}, 
        Parameters: []ast.Parameter{
            {Name: "c", Type: charType},
        },
        ReturnType: nil,
    })
    
    symTable.Define("_printBool", Symbol{
        Name:       "_printBool",
        Type:       &ast.PrimitiveType{Name: "function"},
        Scope:      ScopeInfo{ValidFirstLine: 1, ValidLastLine: lastLine}, 
        Parameters: []ast.Parameter{
            {Name: "value", Type: boolType},
        },
        ReturnType: nil,
    })
    
    symTable.Define("_exit", Symbol{
        Name:       "_exit",
        Type:       &ast.PrimitiveType{Name: "function"},
        Scope:      ScopeInfo{ValidFirstLine: 1, ValidLastLine: lastLine}, 
        Parameters: []ast.Parameter{
            {Name: "status", Type: intType},
        },
        ReturnType: nil,
    })
}