package expressions

// iota in expressions
const (
	IotaPlusOne  = iota + 1  // want "iota used without type specification"
	IotaShifted  = 1 << iota // want "iota used without type specification"
	IotaMultiple = iota * 2  // want "iota used without type specification"
)
