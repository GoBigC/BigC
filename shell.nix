{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go 
    gotools 
    gopls 
    delve # Go debugger 

    antlr4
    jdk17

    rars # riscv simulator 

    graphviz # draw .dot files
  ];

  shellHook = ''
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$PATH 
    mkdir -p $GOPATH/src $GOPATH/bin $GOPATH/pkgs

    if [ ! -d "$GOPATH/pkg/mod/github.com/antlr/antlr4" ]; then
      echo "ANTLR4 Go runtime not found. Installing..." 
      go install github.com/antlr/antlr4/runtime/Go/antlr/v4@latest
    else
      echo "ANTLR4 Go runtime ready" 
    fi

    echo $(go version)

    export CLASSPATH=".:${pkgs.antlr4}/share/java/antlr-4.13.0-complete.jar:$CLASSPATH"
    alias antlr4='java -jar ${pkgs.antlr4}/share/java/antlr-4.13.0-complete.jar'
    alias grun='java org.antlr.v4.gui.TestRig'

    zsh 
  '';
}
