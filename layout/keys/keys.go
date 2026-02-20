package keys

type ArgumentKey byte

const (
	PaddingArg ArgumentKey = iota
	GapArg
	SizeArg
	DirectionArg
	AlignmentArg
	ChildrenArg
	IdArg
)

type SizeKey byte

const (
	GrowSize SizeKey = iota
	FitSize
	FixSize
	MinSize
	MaxSize
)

type AlignmentKey byte

const (
	Start AlignmentKey = iota
	Center
	End
)

type DirectionKey byte

const (
	Row DirectionKey = iota
	Column
	Stack
)
