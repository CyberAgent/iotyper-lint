package main

import (
	"golang.org/x/tools/go/analysis"

	"github.com/CyberAgent/iotyper-lint"
)

var AnalyzerPlugin analyzerPlugin

type analyzerPlugin struct{}

func (analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		iotyper.Analyzer,
	}
}
