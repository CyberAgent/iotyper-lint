package testdata

const (
	WithoutTypeIota = iota // want "iota used without type specification"
)

const (
	WithTypeIota int = iota
	WithTypeIotaAnother
)

const (
	WithoutTypeIotaWithComment       = iota //nolint:iotyper
	WithoutTypeIotaWithCommentAll    = iota //nolint:all
	WithoutTypeIotaWithCommentUnused = iota //nolint:unused // want "iota used without type specification"
)
