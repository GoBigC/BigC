This directory contains `ast.go`, here is the process to achieve it: 

1. Read the grammar (`BigC.g4`), determine which nodes are necessary to build the AST 
- Not all symbols in the grammar file qualify to be an AST node (in which case we'd have a concrete syntax tree like `artifact/cst.txt`)
- AST node types include those that are important to the program struture: statement, expression, type, block, terminals, etc. 
- AST node types do not include: operation precedence, delimiters like ';', precedence grouping symbols like '(' or ')'

The ast.go file contains all the node types we considered necessary to build the AST

2. Determine if the node should be a `struct` or an `interface`
- A struct when: simple, concrete, no need for extending, need to hold data, need to be embedded into other structs 
- An interface when: polymorphic, needs extending to subclasses/subinterfaces

Specific example of when to use struct vs interface is explained as comments in `ast.go` 

3. Model the grammar 

Write the structs/interfaces so that it models the structure of the grammar, adding any fields to hold metadata if need be (for example, `Line` and `Column` number is considered metadata about the token).

If you have ever written Entity class to wrap around a database in Java Spring, this is a parallel of that. The structs/interfaces in `ast.go` is a wrapper around the grammar so that we can populate its fields with data later. 

4. Rinse and repeat until covered all rules 

Once again, the work of this is quite technical and formulaic. Remember, the devil is in the details. 