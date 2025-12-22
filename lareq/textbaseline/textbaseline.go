package textbaseline

type TextBaseline = byte

const (
	Top TextBaseline = iota
	Hanging
	Middle
	Alphabetic
	Ideographic
	Bottom
)
