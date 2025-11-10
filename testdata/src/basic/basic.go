package basic

// Basic iota without type
const (
	BasicWithoutType = iota // want "iota used without type specification"
)

// Single-line const declarations
const SingleLineWithoutType = iota  // want "iota used without type specification"
const SingleLineWithType int = iota // OK: has type

// Multiple constants with inherited iota
const (
	FirstInGroup  = iota // want "iota used without type specification"
	SecondInGroup        // No warning: inherits iota value but doesn't use iota directly
	ThirdInGroup         // No warning: same as above
)
