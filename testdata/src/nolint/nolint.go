package nolint

// Basic nolint comments
const (
	DisabledByIotyper = iota //nolint:iotyper // No warning: explicitly disabled
	DisabledByAll     = iota //nolint:all     // No warning: all linters disabled
	NotDisabled       = iota //nolint:unused   // want "iota used without type specification"
)

// Multiple nolint rules with various formats
const (
	MultipleRulesComma = iota //nolint:iotyper,unused     // No warning: iotyper disabled
	MultipleRulesSpace = iota //nolint: iotyper , unused  // No warning: handles spaces
	WrongLinter        = iota //nolint:govet,unused       // want "iota used without type specification"
)
