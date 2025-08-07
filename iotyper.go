package linters

import (
	"go/ast"
	"go/token"
	"slices"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

const (
	name                = "iotyper"
	doc                 = "iotyper is lint to check for iota usage without type specification"
	iotaStr             = "iota"
	allRule             = "all"
	nolintCommentPrefix = "//nolint:"
)

type plugin struct{}

var _ register.LinterPlugin = new(plugin)

func init() {
	register.Plugin(name, New)
}

func New(_ any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	analyzer := &analysis.Analyzer{
		Name: name,
		Doc:  doc,
		Run:  p.run,
	}
	return []*analysis.Analyzer{analyzer}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

func (p *plugin) run(pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		fileName := pass.Fset.File(f.Pos()).Name()
		if strings.HasSuffix(fileName, "_test.go") {
			continue
		}

		ast.Inspect(f, func(node ast.Node) bool {
			gd, ok := node.(*ast.GenDecl)
			if !ok || gd.Tok != token.CONST {
				return true
			}

			for _, spec := range gd.Specs {
				vs, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}

				if vs.Comment != nil {
					shouldSkip := false
					for _, c := range vs.Comment.List {
						rules := parseNolintComment(c.Text)
						if hasIOTyperRule(rules) {
							shouldSkip = true
							break
						}
					}
					if shouldSkip {
						continue
					}
				}
				for _, value := range vs.Values {
					if id, ok := value.(*ast.Ident); ok && id.Name == iotaStr && vs.Type == nil {
						pass.Reportf(id.Pos(), "iota used without type specification")
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func parseNolintComment(comment string) []string {
	rules := make([]string, 0)
	if strings.HasPrefix(comment, nolintCommentPrefix) {
		rules = strings.Split(comment[len(nolintCommentPrefix):], ",")
		for i := range rules {
			rules[i] = strings.TrimSpace(rules[i])
		}
	}
	return rules
}

func hasIOTyperRule(rules []string) bool {
	return slices.Contains(rules, allRule) || slices.Contains(rules, name)
}
