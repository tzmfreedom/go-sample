package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"github.com/k0kubun/pp"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, os.Args[1], nil, 0)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(ast.Node) bool {
		return true
	})
	debug(ast)
}

func debug(args ...interface{}) {
	pp.Println(args...)
}