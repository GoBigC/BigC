package codegen

// Node is the common interface for all AST nodes
type Node interface {
    // Optional: Add methods like String() or Accept(visitor) if needed
}

// Program represents the root of the AST: program : declaration* EOF
type Program struct {
    Declarations []Declaration
}

// Declaration represents a variable or function declaration
// declaration : type arrayNotation? Identifier declarationRemainder
type Declaration struct {
    Type       string         // "int", "float", "bool", "char", "void"
    IsArray    bool           // True if arrayNotation present
    ArraySize  *int           // Size if array, nil if not array
    Name       string         // Identifier
    Remainder  DeclarationRemainder
}

// DeclarationRemainder is an interface for function or variable declarations
type DeclarationRemainder interface {
    IsFunction() bool
}

// FunctionDeclaration for function declarations
// declarationRemainder : '(' parameterList? ')' block
type FunctionDeclaration struct {
    Parameters []Parameter
    Body       Block
}

func (f *FunctionDeclaration) IsFunction() bool { return true }

// VariableDeclaration for variable declarations
// declarationRemainder : variableInitializer? ';'
type VariableDeclaration struct {
    Initializer Expression // nil if no initializer
}

func (v *VariableDeclaration) IsFunction() bool { return false }

// Parameter represents a function parameter
// parameter : type Identifier
type Parameter struct {
    Type string
    Name string
}

// Block represents a code block
// block : '{' blockItem* '}'
type Block struct {
    Items []BlockItem
}

// BlockItem can be a declaration or statement
type BlockItem interface {
    IsDeclaration() bool
}

// Implements BlockItem for Declaration
func (d *Declaration) IsDeclaration() bool { return true }

// Statement interface for all statement types
type Statement interface {
    Node
    IsStatement() bool
}

// Implements BlockItem for Statement
func (s *IfStatement) IsDeclaration() bool      { return false }
func (s *NonIfStatement) IsDeclaration() bool   { return false }
func (s *IfStatement) IsStatement() bool        { return true }
func (s *NonIfStatement) IsStatement() bool     { return true }

// IfStatement for if statements
// ifStatement : 'if' '(' expression ')' block elseClause?
type IfStatement struct {
    Condition  Expression
    Then       Block
    Else       *Block // nil if no elseClause; could also be *IfStatement for else-if
}

// NonIfStatement for other statements
// nonIfStatement : expression ';' | whileStatement | returnStatement
type NonIfStatement struct {
    ExprStmt   *Expression // nil if not expression statement
    While      *WhileStatement
    Return     *ReturnStatement
}

// WhileStatement for while loops
// whileStatement : 'while' '(' expression ')' block
type WhileStatement struct {
    Condition Expression
    Body      Block
}

// ReturnStatement for return statements
// returnStatement : 'return' expression ';'
type ReturnStatement struct {
    Value Expression
}

// Expression interface for all expressions
type Expression interface {
    Node
}

// AssignmentExpression for assignments
// assignmentExpression : logicalOrExpression assignmentRest?
type AssignmentExpression struct {
    Left  Expression
    Right *Expression // nil if no assignment
}

// LogicalOrExpression for || operations
// logicalOrExpression : logicalAndExpression logicalOrRest*
type LogicalOrExpression struct {
    Left  Expression
    Right []Expression // Multiple || operations
}

// LogicalAndExpression for && operations
// logicalAndExpression : equalityExpression logicalAndRest*
type LogicalAndExpression struct {
    Left  Expression
    Right []Expression
}

// EqualityExpression for ==, !=
// equalityExpression : comparisonExpression equalityRest*
type EqualityExpression struct {
    Left     Expression
    Operator []string    // "==" or "!=", one per rest
    Right    []Expression
}

// ComparisonExpression for <, <=, >, >=
// comparisonExpression : additionExpression comparisonRest*
type ComparisonExpression struct {
    Left     Expression
    Operator []string    // ">", "<", ">=", "<="
    Right    []Expression
}

// AdditionExpression for +, -
// additionExpression : multiplicationExpression additionExpressionRest*
type AdditionExpression struct {
    Left     Expression
    Operator []string    // "+" or "-"
    Right    []Expression
}

// MultiplicationExpression for *, /, %
// multiplicationExpression : unaryExpression multiplicationExpressionRest*
type MultiplicationExpression struct {
    Left     Expression
    Operator []string    // "*", "/", "%"
    Right    []Expression
}

// UnaryExpression for unary operators
// unaryExpression : postfixExpression | unaryOperator unaryExpression
type UnaryExpression struct {
    Operator string      // "++", "--", "!", "" if none
    Expr     Expression // Postfix or nested unary
}

// PostfixExpression for array access, function calls, postfix ++/--
// postfixExpression : primaryExpression (arrayAccess | functionCallArgs | increaseDecrease)?
type PostfixExpression struct {
    Base      Expression
    ArrayIdx  *Expression      // nil if not array access
    CallArgs  []Expression     // empty if not function call
    PostOp    string           // "++", "--", "" if none
}

// PrimaryExpression for basic values
// primaryExpression : Identifier | constant | '(' expression ')'
type PrimaryExpression struct {
    Identifier string     // "" if not variable
    Constant   *Constant  // nil if not constant
    ParenExpr  *Expression // nil if not parenthesized
}

// Constant for literal values
// constant : IntegerConstant | FloatingConstant | BooleanConstant | CharConstant
type Constant struct {
    IntValue    *int     // nil if not int
    FloatValue  *float64 // nil if not float
    BoolValue   *bool    // nil if not bool
    CharValue   *byte    // nil if not char
}