package textalign

type TextAlign = byte

const (
	Left TextAlign = iota
	Right
	Center
	Start
	End
)
