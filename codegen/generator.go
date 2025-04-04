// codegen/generator.go
package codegen

import (
    "fmt"
)

func (cg *CodeGenerator) generateDeclaration(decl *Declaration) {
    switch remainder := decl.Remainder.(type) {
    case *VariableDeclaration:
        if remainder.Initializer != nil {
            reg := cg.generateExpression(remainder.Initializer)
            cg.output.WriteString(fmt.Sprintf("    sw %s, %s\n", reg, decl.Name))
            cg.regManager.Free(reg)
        }
    case *FunctionDeclaration:
        // Handle functions later; for now, assume main is the only function
        cg.generateBlock(remainder.Body)
    }
}

func (cg *CodeGenerator) generateBlock(block Block) {
    for _, item := range block.Items {
        switch node := item.(type) {
        case *Declaration:
            cg.generateDeclaration(node)
        case *IfStatement:
            cg.generateIfStatement(node)
        case *NonIfStatement:
            cg.generateNonIfStatement(node)
        }
    }
}

func (cg *CodeGenerator) generateIfStatement(stmt *IfStatement) {
    condReg := cg.generateExpression(stmt.Condition)
    endLabel := cg.nextLabel()
    cg.output.WriteString(fmt.Sprintf("    beq %s, zero, %s\n", condReg, endLabel))
    cg.regManager.Free(condReg)
    cg.generateBlock(stmt.Then)
    if stmt.Else != nil {
        elseLabel := cg.nextLabel()
        cg.output.WriteString(fmt.Sprintf("    j %s\n", elseLabel))
        cg.output.WriteString(fmt.Sprintf("%s:\n", endLabel))
        cg.generateBlock(*stmt.Else)
        cg.output.WriteString(fmt.Sprintf("%s:\n", elseLabel))
    } else {
        cg.output.WriteString(fmt.Sprintf("%s:\n", endLabel))
    }
}

func (cg *CodeGenerator) generateNonIfStatement(stmt *NonIfStatement) {
    if stmt.ExprStmt != nil {
        reg := cg.generateExpression(*stmt.ExprStmt)
        cg.regManager.Free(reg) // Expression evaluated but not stored
    } else if stmt.While != nil {
        cg.generateWhileStatement(stmt.While)
    } else if stmt.Return != nil {
        reg := cg.generateExpression(stmt.Return.Value)
        cg.output.WriteString(fmt.Sprintf("    mv a0, %s\n", reg)) // Return value in a0
        cg.regManager.Free(reg)
    }
}

func (cg *CodeGenerator) generateWhileStatement(stmt *WhileStatement) {
    startLabel := cg.nextLabel()
    endLabel := cg.nextLabel()
    cg.output.WriteString(fmt.Sprintf("%s:\n", startLabel))
    condReg := cg.generateExpression(stmt.Condition)
    cg.output.WriteString(fmt.Sprintf("    beq %s, zero, %s\n", condReg, endLabel))
    cg.regManager.Free(condReg)
    cg.generateBlock(stmt.Body)
    cg.output.WriteString(fmt.Sprintf("    j %s\n", startLabel))
    cg.output.WriteString(fmt.Sprintf("%s:\n", endLabel))
}

func (cg *CodeGenerator) generateExpression(expr Expression) string {
    switch e := expr.(type) {
    case *AssignmentExpression:
        return cg.generateAssignmentExpression(e)
    case *AdditionExpression:
        return cg.generateAdditionExpression(e)
    case *PrimaryExpression:
        return cg.generatePrimaryExpression(e)
    // Add other expression types as needed
    default:
        panic(fmt.Sprintf("Unsupported expression type: %T", e))
    }
}

func (cg *CodeGenerator) generateAssignmentExpression(expr *AssignmentExpression) string {
    leftReg := cg.generateExpression(expr.Left)
    if expr.Right != nil {
        rightReg := cg.generateExpression(*expr.Right)
        cg.output.WriteString(fmt.Sprintf("    sw %s, 0(%s)\n", rightReg, leftReg)) // Store to memory
        cg.regManager.Free(rightReg)
    }
    return leftReg
}

func (cg *CodeGenerator) generateAdditionExpression(expr *AdditionExpression) string {
    resultReg := cg.generateExpression(expr.Left)
    for i, right := range expr.Right {
        rightReg := cg.generateExpression(right)
        op := expr.Operator[i]
        tempReg := cg.regManager.Allocate()
        if op == "+" {
            cg.output.WriteString(fmt.Sprintf("    add %s, %s, %s\n", tempReg, resultReg, rightReg))
        } else if op == "-" {
            cg.output.WriteString(fmt.Sprintf("    sub %s, %s, %s\n", tempReg, resultReg, rightReg))
        }
        cg.regManager.Free(resultReg)
        cg.regManager.Free(rightReg)
        resultReg = tempReg
    }
    return resultReg
}

func (cg *CodeGenerator) generatePrimaryExpression(expr *PrimaryExpression) string {
    reg := cg.regManager.Allocate()
    if expr.Identifier != "" {
        cg.output.WriteString(fmt.Sprintf("    la %s, %s\n", reg, expr.Identifier))
    } else if expr.Constant != nil {
        if expr.Constant.IntValue != nil {
            cg.output.WriteString(fmt.Sprintf("    li %s, %d\n", reg, *expr.Constant.IntValue))
        }
    } else if expr.ParenExpr != nil {
        return cg.generateExpression(*expr.ParenExpr)
    }
    return reg
}

func (cg *CodeGenerator) nextLabel() string {
    label := fmt.Sprintf("L%d", cg.labelCount)
    cg.labelCount++
    return label
}