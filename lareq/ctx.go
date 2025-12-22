package lareq

import (
	"slices"

	"github.com/laranatech/gorana/lareq/command"
	"github.com/laranatech/gorana/lareq/style"
)

// Adds `setCtx` command to queue
// The options are sorted
// If several options with the same key are provided, then only the first will count
func (q *RenderQueue) SetCtx(opts ...style.StyleOption) {
	slices.SortFunc(opts, func(a, b style.StyleOption) int {
		if a.Key < b.Key {
			return -1
		}
		if a.Key > b.Key {
			return 1
		}
		return 0
	})

	r := []style.StyleOption{}

	for i, v := range opts {
		if i == 0 {
			r = append(r, v)
			continue
		}

		if v.Key == opts[i-1].Key {
			continue
		}

		r = append(r, v)
	}

	q.Commands = append(q.Commands, RenderCommand{
		Command: command.SetCtx,
		Options: r,
	})
}

func MatchSetCtx(a, b RenderCommand) bool {
	if a.Command != command.SetCtx || b.Command != command.SetCtx {
		return false
	}

	aOpts := a.Options.([]style.StyleOption)
	bOpts := b.Options.([]style.StyleOption)

	if len(aOpts) != len(bOpts) {
		return false
	}

	for i := range len(aOpts) {
		if !style.MatchStyleOption(aOpts[i], bOpts[i]) {
			return false
		}
	}
	return true
}
