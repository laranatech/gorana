package lareq

import (
	"larana.tech/go/gorana/lareq/command"
	"larana.tech/go/gorana/lareq/style"
)

type Box struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type RenderCommand struct {
	Command command.CommandByte `json:"c"`
	Options any                 `json:"o"`
}

type RenderQueue struct {
	Commands []RenderCommand `json:"c"`
}

type Queue interface {
	GetCommands() *[]RenderCommand
	DedupeCtx() *[]RenderCommand
	Arc(opts ArcOptions)
	ArcTo(opts ArcToOptions)
	BeginPath()
	BezierQurveTo(opts BezierQurveToOptions)
	ClearRect(b Box)
	Clip()
	ClosePath()
	DrawImage(opts DrawImageOptions)
	DrawSprite(opts DrawSpriteOptions)
	Fill()
	FillText(opts TextOptions)
	LineTo(p Point)
	MoveTo(p Point)
	PasteBitmap(opts PasteBitmapOptions)
	QuadraticCurveTo(opts QuadraticCurveToOptions)
	Rect(b Box)
	Reset()
	ResetTransform()
	Restore()
	Rotate(angle float32)
	RoundedRect(opts RoundedRectOptions)
	Save()
	Scale(p Point)
	SetLineDash(opts []float32)
	SetTransform(matrix []float32)
	SetCtx(opts ...style.StyleOption)
	Stroke()
	StrokeText(opts TextOptions)
	Transform(matrix []float32)
	Translate(p Point)
	Ellipse(opts EllipseOpts)
	Tick()
}
