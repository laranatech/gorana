package lareq

import "github.com/laranatech/gorana/lareq/command"

type ArcOptions struct {
	R                float64 `json:"r"`
	X                float64 `json:"x"`
	Y                float64 `json:"y"`
	Start            float64 `json:"s"`
	End              float64 `json:"e"`
	CounterClockwise bool    `json:"c"`
}

func (q *RenderQueue) Arc(opts ArcOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Arc,
		Options: opts,
	})
}

type EllipseOpts struct {
	P                Point   `json:"p"`
	R                Point   `json:"r"`
	Rotation         float64 `json:"rt"`
	Start            float64 `json:"s"`
	End              float64 `json:"e"`
	Counterclockwise bool    `json:"c"`
}

func (q *RenderQueue) Ellipse(opts EllipseOpts) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Ellipse,
		Options: opts,
	})
}
