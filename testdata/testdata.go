// Package testdata contains test cases for the iotyper linter.
package testdata

// Basic iota usage without type (should trigger warning)
const (
	BasicWithoutType = iota // want "iota used without type specification"
)

// Single-line const declarations
const SingleLineWithoutType = iota // want "iota used without type specification"
const SingleLineWithType int = iota // OK: has type

// Multiple constants with inherited iota
const (
	FirstInGroup  = iota // want "iota used without type specification"
	SecondInGroup        // No warning: inherits iota value but doesn't use iota directly
	ThirdInGroup         // No warning: same as above
)

// iota in expressions
const (
	IotaPlusOne  = iota + 1  // want "iota used without type specification"
	IotaShifted  = 1 << iota // want "iota used without type specification"
	IotaMultiple = iota * 2  // want "iota used without type specification"
)

// iota with explicit type (should NOT trigger warning)
const (
	WithTypeInt      int = iota // OK: explicit int type
	WithTypeIntAgain int = iota // OK: explicit int type
)

// Mixed type specifications in same block
const (
	MixedWithType      int = iota // OK: has type
	MixedWithoutType       = iota // want "iota used without type specification"
	MixedAgainWithType int = iota // OK: has type
)

// nolint comments
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

// Non-iota constants (should NOT trigger warnings)
const (
	PlainNumber = 42
	PlainString = "hello"
	PlainBool   = true
)
