package iotyper_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/CyberAgent/iotyper-lint"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)

	// Run tests for each test category
	testDirs := []string{
		"basic",       // Basic iota detection tests
		"expressions", // Tests for iota in expressions
		"typed",       // Tests with type specifications
		"nolint",      // Tests for nolint comment handling
		"edge_cases",  // Edge cases and non-iota constants
	}

	for _, dir := range testDirs {
		t.Run(dir, func(t *testing.T) {
			analysistest.Run(t, testdata, iotyper.Analyzer, dir)
		})
	}
}
