package linters

import (
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPluginIOTyper(t *testing.T) {
	newPlugin, err := register.GetPlugin("iotyper")
	require.NoError(t, err)

	p, err := newPlugin(nil)
	require.NoError(t, err)

	analyzers, err := p.BuildAnalyzers()
	require.NoError(t, err)

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzers[0], "")
}
