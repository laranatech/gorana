package layout

import (
	"github.com/laranatech/gorana/layout/keys"
)

func (node *NodeItem) IsAlongAxis(axis Axis) bool {
	if axis == XAxis && node.Direction == keys.Row {
		return true
	}
	if axis == YAxis && node.Direction == keys.Column {
		return true
	}
	return false
}

func Row() *Argument {
	return &Argument{
		Key:   keys.DirectionArg,
		Value: keys.Row,
	}
}

func Column() *Argument {
	return &Argument{
		Key:   keys.DirectionArg,
		Value: keys.Column,
	}
}
