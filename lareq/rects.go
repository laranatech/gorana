package lareq

import "github.com/laranatech/gorana/lareq/command"

func (q *RenderQueue) Rect(b Box) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Rect,
		Options: b,
	})
}

func (q *RenderQueue) ClearRect(b Box) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.ClearRect,
		Options: b,
	})
}

type RoundedRectOptions struct {
	Box
	R float32 `json:"r"`
}

func (q *RenderQueue) RoundedRect(opts RoundedRectOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.RoundedRect,
		Options: opts,
	})
}
