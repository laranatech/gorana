package command

type CommandByte = byte

const (
	Arc CommandByte = iota
	ArcTo
	BeginPath
	BezierQurveTo
	ClearRect
	Clip
	ClosePath
	DrawImage
	DrawSprite
	Fill
	FillText
	LineTo
	MoveTo
	PasteBitmap
	QuadraticCurveTo
	Rect
	Reset
	ResetTransform
	Restore
	Rotate
	RoundedRect
	Save
	Scale
	SetLineDash
	SetTransform
	SetCtx
	Stroke
	StrokeText
	Transform
	Translate
	Ellipse
	Tick
)
