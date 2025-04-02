# BigC
A simple C and Java inspired language for CS212 Programming Languages Paradigm

## Build instruction 
1. Clone & Go into the project dir 
```
git clone https://github.com/GoBigC/BigC.git
cd BigC
```

2. Update Go module dependency
```
go mod tidy
```
(dont have to do this all the time, once in a while is okay)

3. Run sript `run.sh` in the project root to generate parser files through ANTLR
```
bash run.sh
```

4. Build and run Go module 
Stand at project root and run: 

```
go build . 
go run . test/sample.uia
```

After this, find the concrete syntax tree at `artifact/cst.txt`, and abstract syntax tree `at artifact/ast.txt`