package ast

// [Explain] interface is used for objects that 
// would be extend over time (a Node satisfy this, 
// because it could be branch node or leaf node)
// The Node especially need to be an interface, 
// because it is very polymorphic -- all AST nodes 
// will be a specific kind of Node 
type Node interface {
	isNode()
}

// [Explain] struct is used if the object needs to
// hold data (BaseNode satisfy this because it should
// contain line & col information for bug reports)
type BaseNode struct {
	Line 		int 
	Column 		int 
	EndLine 	int
	EndColumn	int
}

// [Design Pattern] Sealed interface
// https://www.baeldung.com/java-sealed-classes-interfaces
// the existence of this method prevents accidental 
// implementations (have to explicitly implement this method
// to qualify as a Declaration AST node)
// Having this extra method helps with that, while not adding 
// any additional cost -- the method is empty, it does not use 
// any computation resource, cost nothing :)
func (b *BaseNode) isNode(){} 

// ----- Program structure -----
// [Explain] struct is also used when the object is 
// rather simple/concrete and would not need to be  
// extended over time (Program satisfy this because a 
// program is a list of declarations, no polymorphs here)
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

type ArrayType struct{
	BaseType
	ElementType	Type
	Size		Expression
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

type FunctionCallExpression struct {
	BaseExpression 
	Function 	Expression 
	Arguments	[]Expression
}

type BinaryExpression struct { // likely unused :)
	// contains + - * / 
	BaseExpression 
	Left 		Expression 
	Operator	string 
	Right 		Expression 
}

type UnaryExpression struct { 
	// contains ! (logical not)
	BaseExpression 
	Operator 	string 
	Operand		Expression 
}

type ArrayAccessExpression struct {
	BaseExpression 
	Array 	Expression // to support matrix defn, where mat[i][j] then mat[i] is an expression, not just a var
	Index	Expression
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