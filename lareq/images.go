package lareq

import "larana.tech/go/gorana/lareq/command"

type DrawImageOptions struct {
	Img string `json:"i"`
	Box
}

func (q *RenderQueue) DrawImage(opts DrawImageOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.DrawImage,
		Options: opts,
	})
}

type DrawSpriteOptions struct {
	Img         string `json:"i"`
	Source      Box    `json:"s"`
	Destination Box    `json:"d"`
}

func (q *RenderQueue) DrawSprite(opts DrawSpriteOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.DrawSprite,
		Options: opts,
	})
}

type PasteBitmapOptions struct {
	Bitmap   []byte `json:"b"`
	Length   int    `json:"l"`
	Channels int    `json:"c"`
}

func (q *RenderQueue) PasteBitmap(opts PasteBitmapOptions) {
	q.Commands = append(q.Commands, RenderCommand{
		Command: command.PasteBitmap,
		Options: opts,
	})
}
