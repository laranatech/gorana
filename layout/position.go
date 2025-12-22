package layout

import "github.com/laranatech/gorana/layout/keys"

func (node *NodeItem) SetPositionByAxis(axis Axis, value float32) {
	switch axis {
	case XAxis:
		node.Box.X = value
	case YAxis:
		node.Box.Y = value
	}
}

func (node *NodeItem) GetPositionByAxis(axis Axis) float32 {
	switch axis {
	case XAxis:
		return node.Box.X
	case YAxis:
		return node.Box.Y
	}
	return 0
}

func ComputePosition(axis Axis, node *NodeItem) error {
	if node.IsRoot() {
		node.SetPositionByAxis(axis, 0)
	}

	err := computeChildrenPositions(axis, node)

	if err != nil {
		return err
	}

	for _, child := range node.Children {
		err := ComputePosition(axis, child)
		if err != nil {
			return err
		}
	}

	return nil
}

func computeChildrenPositions(axis Axis, node *NodeItem) error {
	var totalSide float32 = 0

	for _, child := range node.Children {
		totalSide += child.GetSideByAxis(axis)
	}

	initialOffset := computeInitialOffset(axis, node, totalSide)
	var offset float32 = 0

	for i, child := range node.Children {
		if i == 0 || !node.IsAlongAxis(axis) {
			child.SetPositionByAxis(axis, initialOffset)
			offset += node.Gap + child.GetSideByAxis(axis)
			continue
		}
		child.SetPositionByAxis(axis, offset)
		offset += node.Gap + child.GetSideByAxis(axis)
	}

	return nil
}

func computeInitialOffset(axis Axis, node *NodeItem, total float32) float32 {
	var totalGap float32 = float32(len(node.Children)-1) * node.Gap
	p := node.GetPositionByAxis(axis)

	switch node.Alignment {
	case keys.Center:
		return p - node.GetSideByAxis(axis) - ((total / 2) + totalGap)
	case keys.End:
		return p - node.GetSideByAxis(axis) - (total + totalGap)
	}
	return node.GetInitialpaddingByAxis(axis)
}
