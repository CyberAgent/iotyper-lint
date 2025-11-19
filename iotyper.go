package iotyper

import (
	"go/ast"
	"go/token"
	"slices"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const (
	name                = "iotyper"
	doc                 = "check for iota usage without type specification"
	iotaStr             = "iota"
	allRule             = "all"
	nolintCommentPrefix = "//nolint:"
)

var Analyzer = &analysis.Analyzer{
	Name: name,
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func init() {
	register.Plugin(Analyzer.Name, New)
}

type PluginTmpl struct{}

func New(_ any) (register.LinterPlugin, error) {
	return &PluginTmpl{}, nil
}

func (p *PluginTmpl) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{Analyzer}, nil
}

func (p *PluginTmpl) GetLoadMode() string {
	return register.LoadModeSyntax
}

// run performs the actual linting logic for the analyzer.
func run(pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		// Get file information to check if it's a test file
		filePos := pass.Fset.File(f.Pos())
		if filePos == nil {
			continue
		}
		fileName := filePos.Name()
		// Skip test files as they often use iota for test case numbering
		if strings.HasSuffix(fileName, "_test.go") {
			continue
		}

		// Traverse the AST to find const declarations
		ast.Inspect(f, func(node ast.Node) bool {
			// GenDecl represents general declarations: import, const, type, or var
			// We're only interested in const declarations that might use iota
			gd, ok := node.(*ast.GenDecl)
			if !ok || gd.Tok != token.CONST {
				// Not a general declaration or not a const declaration, continue traversing
				return true
			}

			// Process each constant specification in the declaration
			// e.g., in "const (A = iota; B = 2)", processes A and B separately
			for _, spec := range gd.Specs {
				vs, ok := spec.(*ast.ValueSpec)
				if !ok {
					// Not a value specification, skip
					continue
				}

				// Check for nolint comments that disable this linter
				// Supports: //nolint:iotyper or //nolint:all
				if vs.Comment != nil {
					shouldSkip := false
					for _, c := range vs.Comment.List {
						rules := parseNolintComment(c.Text)
						if hasIotyperRule(rules) {
							shouldSkip = true
							break
						}
					}
					if shouldSkip {
						continue
					}
				}

				// Skip if type is explicitly specified
				// e.g., "const x MyType = iota" is OK
				if vs.Type != nil {
					continue
				}

				// Check if any of the values contains iota without type specification
				// e.g., "const x = iota" or "const x = iota + 1" triggers a warning
				for _, value := range vs.Values {
					// Check if the value contains iota
					if containsIota(value) {
						// Find the position of iota for error reporting
						iotaPos := findIotaPos(value)
						if iotaPos != token.NoPos {
							pass.Reportf(iotaPos, "iota used without type specification")
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

// parseNolintComment extracts linter rules from a nolint comment.
// Example: "//nolint:iotyper,govet" returns ["iotyper", "govet"]
func parseNolintComment(comment string) []string {
	if !strings.HasPrefix(comment, nolintCommentPrefix) {
		return nil
	}

	// Extract the rules part (after //nolint:)
	rulesStr := comment[len(nolintCommentPrefix):]

	// Handle additional comments after the rules (e.g., "//nolint:iotyper // reason")
	if idx := strings.Index(rulesStr, " //"); idx != -1 {
		rulesStr = rulesStr[:idx]
	}

	// Split by comma and trim spaces
	rules := strings.Split(rulesStr, ",")
	for i := range rules {
		rules[i] = strings.TrimSpace(rules[i])
	}
	return rules
}

// hasIotyperRule checks if the nolint comment disables this linter.
// Returns true if the comment contains "iotyper" or "all".
func hasIotyperRule(rules []string) bool {
	return slices.Contains(rules, allRule) || slices.Contains(rules, name)
}

// containsIota recursively checks if an expression contains iota
func containsIota(expr ast.Expr) bool {
	var found bool
	ast.Inspect(expr, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == iotaStr {
			found = true
			return false
		}
		return true
	})
	return found
}

// findIotaPos finds the position of the first iota in an expression
func findIotaPos(expr ast.Expr) token.Pos {
	var pos token.Pos
	ast.Inspect(expr, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == iotaStr {
			pos = id.Pos()
			return false
		}
		return true
	})
	return pos
}
