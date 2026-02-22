package style

import (
	"larana.tech/go/gorana/lareq/fontstyle"
	"larana.tech/go/gorana/lareq/fontweight"
	"larana.tech/go/gorana/lareq/linecap"
	"larana.tech/go/gorana/lareq/linejoin"
	"larana.tech/go/gorana/lareq/style/keys"
	"larana.tech/go/gorana/lareq/textalign"
	"larana.tech/go/gorana/lareq/textbaseline"
)

type StyleOption struct {
	Key   keys.Key `json:"k"`
	Value any      `json:"v"`
}

func FillStyle(value string) StyleOption {
	return StyleOption{Key: keys.FillStyle, Value: value}
}

func LineCap(value linecap.LineCap) StyleOption {
	return StyleOption{Key: keys.LineCap, Value: value}
}

func LineWidth(value float64) StyleOption {
	return StyleOption{Key: keys.LineWidth, Value: value}
}

func StrokeStyle(value string) StyleOption {
	return StyleOption{Key: keys.StrokeStyle, Value: value}
}

func FontFace(value string) StyleOption {
	return StyleOption{Key: keys.FontFace, Value: value}
}

func FontSize(value float64) StyleOption {
	return StyleOption{Key: keys.FontSize, Value: value}
}

func FontStyle(value fontstyle.FontStyle) StyleOption {
	return StyleOption{Key: keys.FontStyle, Value: value}
}

func FontWeight(value fontweight.FontWeight) StyleOption {
	return StyleOption{Key: keys.FontWeight, Value: value}
}

func TextBaseline(value textbaseline.TextBaseline) StyleOption {
	return StyleOption{Key: keys.TextBaseline, Value: value}
}

func TextAlign(value textalign.TextAlign) StyleOption {
	return StyleOption{Key: keys.TextAlign, Value: value}
}

func Radius(value float32) StyleOption {
	return StyleOption{Key: keys.Radius, Value: value}
}

func LineDashOffset(value float32) StyleOption {
	return StyleOption{Key: keys.LineDashOffset, Value: value}
}

func LineJoin(value linejoin.LineJoin) StyleOption {
	return StyleOption{Key: keys.LineJoin, Value: value}
}

func Pattern(value string) StyleOption {
	return StyleOption{Key: keys.Pattern, Value: value}
}

func MatchStyleOption(a, b StyleOption) bool {
	if a.Key != b.Key {
		return false
	}
	if a.Value != b.Value {
		return false
	}
	return true
}
