package main

import (
	"github.com/tzmfreedom/go-sample/analysis2"
	"golang.org/x/tools/go/analysis/multichecker"
	//"golang.org/x/tools/go/analysis/passes/inspect"
)

func main() {
	multichecker.Main(
		analysis2.Analyzer,
	)
}
