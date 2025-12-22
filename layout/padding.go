package layout

import (
	"github.com/laranatech/gorana/layout/keys"
)

type PaddingValue struct {
	Top    float32 `json:"t"`
	Bottom float32 `json:"b"`
	Left   float32 `json:"l"`
	Right  float32 `json:"r"`
}

func Padding(args ...float32) *Argument {
	var p *PaddingValue = nil
	if len(args) == 1 {
		p = &PaddingValue{
			Top:    args[0],
			Bottom: args[0],
			Left:   args[0],
			Right:  args[0],
		}
	} else if len(args) == 2 {
		p = &PaddingValue{
			Top:    args[0],
			Bottom: args[0],
			Left:   args[1],
			Right:  args[1],
		}
	} else if len(args) == 4 {
		p = &PaddingValue{
			Top:    args[0],
			Right:  args[1],
			Bottom: args[2],
			Left:   args[3],
		}
	}

	return &Argument{
		Key:   keys.PaddingArg,
		Value: p,
	}
}

func (node *NodeItem) GetPaddingByAxis(axis Axis) float32 {
	if node.Padding == nil {
		return 0
	}

	if axis == XAxis {
		return node.Padding.Left + node.Padding.Right
	}
	return node.Padding.Bottom + node.Padding.Top
}

func (node *NodeItem) GetInitialpaddingByAxis(axis Axis) float32 {
	if node.Padding == nil {
		return 0
	}

	if axis == XAxis {
		return node.Padding.Left
	}
	return node.Padding.Top
}
