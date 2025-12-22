package lareq

import "github.com/laranatech/gorana/lareq/command"

func (q *RenderQueue) BeginPath() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.BeginPath,
	})
}

func (q *RenderQueue) Clip() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Clip,
	})
}

func (q *RenderQueue) ClosePath() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.ClosePath,
	})
}

func (q *RenderQueue) Fill() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Fill,
	})
}

func (q *RenderQueue) Reset() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Reset,
	})
}

func (q *RenderQueue) ResetTransform() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.ResetTransform,
	})
}

func (q *RenderQueue) Restore() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Restore,
	})
}

func (q *RenderQueue) Rotate(angle float32) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Rotate,
		Options: angle,
	})
}

func (q *RenderQueue) Save() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Save,
	})
}

func (q *RenderQueue) Scale(p Point) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Rotate,
		Options: p,
	})
}

func (q *RenderQueue) SetLineDash(opts []float32) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.SetLineDash,
		Options: opts,
	})
}

func (q *RenderQueue) SetTransform(matrix []float32) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.SetTransform,
		Options: matrix,
	})
}

func (q *RenderQueue) Transform(matrix []float32) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Transform,
		Options: matrix,
	})
}

func (q *RenderQueue) Translate(p Point) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Transform,
		Options: p,
	})
}

func (q *RenderQueue) Stroke() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Stroke,
	})
}

func (q *RenderQueue) Tick() {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.Tick,
	})
}
