grammar BigC;

// ---------- PARSER RULES ----------
program
    : declaration* EOF
    ;

declaration
    : type Identifier arrayNotation? declarationRemainder
    ;

arrayNotation   
    : '[' expression ']'
    ;

type
    : 'int'
    | 'float'
    | 'bool'
    | 'char'
    | 'void' 
    ;

declarationRemainder 
    : '(' parameterList? ')' block 
    | variableInitializer? ';'
    ;

parameterList
    : parameter (',' parameter)*
    ;

parameter
    : type Identifier arrayNotation?
    ;

block 
    : '{' blockItem* '}'
    ;

blockItem
    : declaration
    | statement
    ;

statement
    : ifStatement
    | nonIfStatement
    ;

ifStatement
    : 'if' '(' expression ')' block elseClause?
    ;

elseClause
    : 'else' (block | ifStatement)
    ;

nonIfStatement 
    : expression ';'
    | whileStatement
    | returnStatement
    ;

whileStatement
    : 'while' '(' expression ')' block 
    ;

returnStatement 
    : 'return' expression ';'
    ;

// Expression precedence (from highest to lowest):
// 1. Primary expressions (constants, variables, parenthesized)
// 2. Postfix operations (arr[i], fn(), x++, x--) -- only support the first 2
// 3. Unary operations (!x, ++x, --x) -- only support the first
// 4. Multiplicative (*, /) 
// 5. Additive (+, -)
// 6. Comparison (<, <=, >=, >)
// 7. Equality (==, !=)
// 8. Logical AND (&&)
// 9. Logical OR (||)
// 10. Assignment (=)

expression 
    : assignmentExpression 
    ;

assignmentExpression
    : logicalOrExpression assignmentRest?
    ;

assignmentRest
    : '=' assignmentExpression
    ;

variableInitializer
    : '=' expression
    ;

logicalOrExpression
    : logicalAndExpression logicalOrRest*
    ;

logicalOrRest
    : '||' logicalAndExpression
    ;

logicalAndExpression
    : equalityExpression logicalAndRest*
    ;

logicalAndRest
    : '&&' equalityExpression
    ;

equalityExpression
    : comparisonExpression equalityRest*
    ;

equalityRest
    : equalityOperator comparisonExpression
    ;

equalityOperator 
    : '=='
    | '!='
    ;

comparisonExpression 
    : additionExpression comparisonRest*
    ;

comparisonRest
    : comparisonOperator additionExpression
    ;

comparisonOperator
    : '>'
    | '<'
    | '>='
    | '<='
    ;

additionExpression
    : multiplicationExpression additionExpressionRest*
    ;

additionExpressionRest
    : addSubtractOperator multiplicationExpression
    ;

addSubtractOperator 
    : '+'
    | '-'
    ;

multiplicationExpression
    : unaryExpression multiplicationExpressionRest*
    ;

multiplicationExpressionRest
    : multDivModOperator unaryExpression
    ;

multDivModOperator
    : '*'
    | '/'
    ;

unaryExpression 
    : postfixExpression 
    | unaryOperator unaryExpression 
    ;

unaryOperator
    : '!'
    | '-'
    ;

postfixExpression 
    : primaryExpression (arrayAccess | functionCallArgs)?
    ;

arrayAccess 
    : '[' expression ']'
    ;

functionCallArgs
    : '(' argList? ')'
    ;

argList 
    : assignmentExpression (',' assignmentExpression)*
    ;

primaryExpression 
    : Identifier 
    | constant 
    | '(' expression ')'
    ;

constant
    : IntegerConstant 
    | FloatingConstant 
    | BooleanConstant 
    | CharConstant 
    ;

// ---------- LEXER RULES ----------
BooleanConstant: 'true' | 'false';
CharConstant: '\'' . '\'';
Identifier: [a-zA-Z_][a-zA-Z0-9_]*;

IntegerConstant: [0-9]+;
FloatingConstant: [0-9]+ '.' [0-9]+;

WS: [ \t\r\n]+ -> skip;
COMMENT: '//' ~[\r\n]* -> skip;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;