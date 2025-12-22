package lareq

import "github.com/laranatech/gorana/lareq/command"

type ArcToOptions struct {
	A Point   `json:"a"`
	B Point   `json:"b"`
	R float64 `json:"r"`
}

func (q *RenderQueue) ArcTo(opts ArcToOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.ArcTo,
		Options: opts,
	})
}

type BezierQurveToOptions struct {
	C1 Point `json:"c1"`
	C2 Point `json:"c2"`
	P  Point `json:"p"`
}

func (q *RenderQueue) BezierQurveTo(opts BezierQurveToOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.BezierQurveTo,
		Options: opts,
	})
}

func (q *RenderQueue) MoveTo(p Point) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.MoveTo,
		Options: p,
	})
}

func (q *RenderQueue) LineTo(p Point) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.LineTo,
		Options: p,
	})
}

type QuadraticCurveToOptions struct {
	C Point `json:"c"`
	P Point `json:"p"`
}

func (q *RenderQueue) QuadraticCurveTo(opts QuadraticCurveToOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.QuadraticCurveTo,
		Options: opts,
	})
}
