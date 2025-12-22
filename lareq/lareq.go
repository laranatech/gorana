package lareq

import (
	"github.com/laranatech/gorana/lareq/command"
	"github.com/laranatech/gorana/lareq/style"
)

func New() Queue {
	return &RenderQueue{}
}

func (q *RenderQueue) GetCommands() *[]RenderCommand {
	return &q.Commands
}

func (q *RenderQueue) DedupeCtx() *[]RenderCommand {
	newCommands := []RenderCommand{}

	var prevSet *RenderCommand

	for _, v := range q.Commands {
		if v.Command != command.SetCtx {
			newCommands = append(newCommands, v)
			continue
		}

		if len(v.Options.([]style.StyleOption)) == 0 {
			continue
		}

		if prevSet != nil {
			if MatchSetCtx(*prevSet, v) {
				prevSet = &v
				continue
			}
		}

		newCommands = append(newCommands, v)
		prevSet = &v
	}

	return &newCommands
}
