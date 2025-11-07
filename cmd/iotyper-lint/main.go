package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/CyberAgent/iotyper-lint"
)

func main() {
	unitchecker.Main(iotyper.Analyzer)
}
