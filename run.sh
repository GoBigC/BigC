#!/usr/bin/env bash 

if [ -n $SOURCE_FILE ]; then
    SOURCE_FILE='test/semantic_tests.uia'
else 
    SOURCE_FILE="$1"
fi 

go build && go run . $SOURCE_FILE