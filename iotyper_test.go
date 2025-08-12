package linters

import (
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestPluginIotyper tests the iotyper analyzer for detecting iota usage without type specification.
func TestPluginIotyper(t *testing.T) {
	// Register the plugin explicitly in test
	register.Plugin("iotyper", New)

	// Test plugin registration
	newPlugin, err := register.GetPlugin("iotyper")
	require.NoError(t, err, "Failed to get plugin")
	require.NotNil(t, newPlugin, "Plugin should not be nil")

	// Test plugin initialization
	p, err := newPlugin(nil)
	require.NoError(t, err, "Failed to create plugin instance")
	require.NotNil(t, p, "Plugin instance should not be nil")

	// Test analyzer creation
	analyzers, err := p.BuildAnalyzers()
	require.NoError(t, err, "Failed to build analyzers")
	require.Len(t, analyzers, 1, "Should have exactly one analyzer")
	require.Equal(t, "iotyper", analyzers[0].Name, "Analyzer name should be 'iotyper'")

	// Run analyzer tests with testdata
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzers[0])
}
