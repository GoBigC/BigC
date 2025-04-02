#!/usr/bin/env bash 


GRAMMAR_FILE='grammar/BigC.g4'

# antlr4 -atn -Xforce-atn -Xlog -Dlanguage=Go -visitor -listener $GRAMMAR_FILE -o out/

antlr4 -Dlanguage=Go -Xlog -visitor $GRAMMAR_FILE -o pkg/syntax/parser/
