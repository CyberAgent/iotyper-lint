package typed

// iota with explicit type
const (
	WithTypeInt      int = iota // OK: explicit int type
	WithTypeIntAgain int = iota // OK: explicit int type
)

// Mixed type specifications
const (
	MixedWithType      int = iota // OK: has type
	MixedWithoutType       = iota // want "iota used without type specification"
	MixedAgainWithType int = iota // OK: has type
)
