package layout

import (
	"github.com/laranatech/gorana/layout/keys"
)

func (n *node) IsAlongAxis(axis Axis) bool {
	if axis == XAxis && n.direction == keys.Row {
		return true
	}
	if axis == YAxis && n.direction == keys.Column {
		return true
	}
	if axis == ZAxis && n.direction == keys.Stack {
		return true
	}
	return false
}
