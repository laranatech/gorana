package lareq

import "github.com/laranatech/gorana/lareq/command"

type TextOptions struct {
	Text     string  `json:"t"`
	MaxWidth float64 `json:"mw"`
	Point
}

func (q *RenderQueue) FillText(opts TextOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.FillText,
		Options: opts,
	})
}

func (q *RenderQueue) StrokeText(opts TextOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.StrokeText,
		Options: opts,
	})
}
