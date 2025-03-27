#!/usr/bin/env bash 

if [ -n $GRAMMAR_FILE ]; then
    GRAMMAR_FILE='BigC.g4'
else 
    GRAMMAR_FILE="$1"
fi 

# antlr4 -atn -Xforce-atn -Xlog -Dlanguage=Go -visitor -listener $GRAMMAR_FILE -o out/

antlr4 -Dlanguage=Go -visitor $GRAMMAR_FILE -o parser/
