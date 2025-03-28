package ast

type Node interface {
	isNode()
}

type BaseNode struct {
	Line int 
	Column int 
}

func (b *BaseNode) isNode(){} // required interface implementation 

// ----- Program structure -----
type Program struct {
	BaseNode
	Declarations []Declaration
}

// ----- Declaration -----
type Declaration interface {
	Node
	isDeclaration()
}

type BaseDeclaration struct {
	BaseNode
}

func (b *BaseDeclaration) isDeclaration() {}

type VarDeclaration struct {
	BaseDeclaration 
	Name		string 
	Type		Type 
	Initializer	Expression // nullable
}

type FunctionDeclaration struct {
	BaseDeclaration 
	Name 		string 
	ReturnType 	Type 
	Parameters	[]Parameter 
	Body 		*Block
}

type Parameter struct {
	BaseNode 
	Name string 
	Type Type 
}

// ----- Types -----
type Type interface {
	Node 
	isType()
}

type BaseType struct {
	BaseNode 
}

func (b *BaseType) isType() {}

type PrimitiveType struct {
	BaseType
	Name string
}

// ----- Statements -----
type Statement interface {
	Node 
	isStatement()
}

type BaseStatement struct {
	BaseNode 
}

func (b *BaseStatement) isStatement() {}

type Block struct {
	BaseStatement
	Items 	[]BlockItem
}

type BlockItem interface {
	Node 
	isBlockItem()
}

func (b *BaseDeclaration) isBlockItem() {}
func (b *BaseStatement) isBlockItem() {}

type IfStatement struct {
	BaseStatement
	Condition Expression
	ThenBlock *Block 
	ElseBlock Node // nullable
}

type WhileStatement struct {
	BaseStatement
	Condition 	Expression 
	Body 		*Block
}

type ReturnStatement struct {
	BaseStatement 
	Value Expression
}

type ExpressionStatement struct {
	BaseStatement 
	Expr Expression
}

// ----- Expressions ----- 
type Expression interface {
	Node 
	isExpression()
}

type BaseExpression struct {
	BaseNode 
}

func (b *BaseExpression) isExpression() {}

type BinaryExpression struct {
	// contains + - * / 
	BaseExpression 
	Left 		Expression 
	Operator	string 
	Right 		Expression 
}

type UnaryExpression struct { 
	// contains ! ++ -- 
	BaseExpression 
	Operator 	string 
	Operand		Expression 
	IsPrefix 	bool 	// true if prefix, false if postfix 
}

type FunctionCallExpression struct {
	BaseExpression 
	Function 	Expression 
	Arguments	[]Expression
}

// ----- Identifiers & Literals ----- 
type Identifier struct {
	BaseExpression
	Name string 
}

type IntegerLiteral struct { 
	BaseExpression
	Value int64
}

type FloatLiteral struct {
	BaseExpression
	Value float64
}

type BoolLiteral struct {
	BaseExpression 
	Value bool 
}

type CharLiteral struct {
	BaseExpression 
	Value rune // goofy af name 
}