package linejoin

type LineJoin = byte

const (
	Round LineJoin = iota
	Bevel
	Miter
)
