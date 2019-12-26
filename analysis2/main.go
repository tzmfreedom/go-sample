package analysis2

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "unusedresult",
	Doc:  "check for unused results of calls to some functions",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			switch n.(type) {
			case *ast.File:
				if f.Name.String() == "analysis2" {
					pass.Reportf(f.Pos(), "ERROR: %s", f.Name.String())
				}
			}
			return true
		})
	}
	return nil, nil
}
