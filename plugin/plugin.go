package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/CyberAgent/iotyper-lint"
)

func init() {
	register.Plugin(iotyper.Analyzer.Name, New)
}

type PluginTmpl struct{}

func New(_ any) (register.LinterPlugin, error) {
	return &PluginTmpl{}, nil
}

func (p *PluginTmpl) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{iotyper.Analyzer}, nil
}

func (p *PluginTmpl) GetLoadMode() string {
	return register.LoadModeSyntax
}
