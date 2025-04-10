package table

import (
	"BigCooker/pkg/syntax/ast"
	"fmt"
	"os"
	"text/tabwriter"
)

type SymbolTable struct {
    Parent    *SymbolTable
    Symbols   map[string]Symbol
    ScopeType string
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
func NewSymbolTable(parent *SymbolTable, scopeType string) *SymbolTable {
    return &SymbolTable{
        Parent:    parent,
        Symbols:   make(map[string]Symbol),
        ScopeType: scopeType,
    }
}

func (symTable *SymbolTable) Define(name string, symbol Symbol) {
    symTable.Symbols[name] = symbol
}

func (symTable *SymbolTable) Lookup(name string) (Symbol, bool) {
    if sym, ok := symTable.Symbols[name]; ok {
        return sym, true
    }
    if symTable.Parent != nil {
        return symTable.Parent.Lookup(name)
    }
    return Symbol{}, false
}

func (symTable *SymbolTable) PrintTable() {
    // Collect all symbols from all scopes
    symbols := symTable.collectSymbols()

    // Set up tabwriter for aligned columns
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
    fmt.Fprintln(w, "Name\tType\tScope\tArraySize\tValue\tParameters\tReturnType")
    fmt.Fprintln(w, "----\t----\t-----\t---------\t-----\t----------\t----------")

    // Print each symbol as a row
    for _, sym := range symbols {
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

// collectSymbols gathers all symbols from the current scope and its parents.
func (symTable *SymbolTable) collectSymbols() []Symbol {
    var symbols []Symbol
    seen := make(map[string]bool) // Track symbols to avoid duplicates from parent scopes

    // Traverse from current scope up to global
    current := symTable
    for current != nil {
        for name, sym := range current.Symbols {
            if !seen[name] {
                symbols = append(symbols, sym)
                seen[name] = true
            }
        }
        current = current.Parent
    }
    return symbols
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