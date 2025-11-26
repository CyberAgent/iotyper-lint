package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/CyberAgent/iotyper-lint"
)

func main() {
	singlechecker.Main(iotyper.Analyzer)
}
