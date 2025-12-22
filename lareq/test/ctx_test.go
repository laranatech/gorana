package lareq_test

import (
	"testing"

	"github.com/laranatech/gorana/lareq"
	"github.com/laranatech/gorana/lareq/command"
	"github.com/laranatech/gorana/lareq/linecap"
	"github.com/laranatech/gorana/lareq/style"
	"github.com/laranatech/gorana/lareq/style/keys"
)

func TestStxCommandContructor(t *testing.T) {
	q := lareq.New()
	q.SetCtx(style.FillStyle("#000"), style.LineCap(linecap.Butt), style.FillStyle("asd"), style.FillStyle("asd"), style.FillStyle("asd"))

	result := (*q.GetCommands())[0]
	expected := lareq.RenderCommand{
		Command: command.SetCtx,
		Options: []style.StyleOption{{Key: keys.FillStyle, Value: "#000"}, {Key: keys.LineCap, Value: linecap.Butt}},
	}

	if !lareq.MatchSetCtx(result, expected) {
		t.Errorf("Invalid setCtx constuction.\nGot: %v\nexpected: %v", result, expected)
	}
}
