package layout

import (
	"github.com/laranatech/gorana/layout/keys"
)

func (n *node) SetPositionByAxis(axis Axis, value float32) {
	switch axis {
	case XAxis:
		n.cube.X = value
	case YAxis:
		n.cube.Y = value
	case ZAxis:
		n.cube.Z = value
	}
}

func (n *node) GetPositionByAxis(axis Axis) float32 {
	switch axis {
	case XAxis:
		return n.cube.X
	case YAxis:
		return n.cube.Y
	case ZAxis:
		return n.cube.Z
	}
	return 0
}

func ComputePosition(axis Axis, n *node) error {
	if n.IsRoot() {
		n.SetPositionByAxis(axis, 0)
	}

	if err := computeChildrenPositions(axis, n); err != nil {
		return err
	}

	for _, child := range n.children {
		if err := ComputePosition(axis, child); err != nil {
			return err
		}
	}

	return nil
}

func computeChildrenPositions(axis Axis, n *node) error {
	var totalSide float32 = 0

	for _, child := range n.children {
		totalSide += child.GetSideByAxis(axis)
	}

	initialOffset := computeInitialOffset(axis, n, totalSide)

	var offset float32 = initialOffset

	for i, child := range n.children {
		if i == 0 || !n.IsAlongAxis(axis) {
			child.SetPositionByAxis(axis, initialOffset)
			offset += n.gap + child.GetSideByAxis(axis)
			continue
		}
		child.SetPositionByAxis(axis, offset)
		offset += n.gap + child.GetSideByAxis(axis)
	}

	return nil
}

func computeInitialOffset(axis Axis, n *node, total float32) float32 {
	var totalGap float32 = float32(len(n.children)-1) * n.gap
	p := n.GetPositionByAxis(axis)
	initialPadding := n.GetInitialPaddingByAxis(axis)

	al, _ := n.alignment[axis]

	switch al {
	case keys.Center:
		return p - n.GetSideByAxis(axis) - ((total / 2) + totalGap)
	case keys.End:
		return p - n.GetSideByAxis(axis) - (total + totalGap)
	}
	return p + initialPadding
}
