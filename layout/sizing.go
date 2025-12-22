package layout

import (
	"errors"

	"github.com/laranatech/gorana/layout/keys"
)

type SizeArgument struct {
	Key   keys.SizeKey
	Value float32
}

type AxisSize struct {
	Min   float32
	Max   float32
	Value float32
	Type  keys.SizeKey
	Axis  Axis
}

func (node *NodeItem) IsFix(axis Axis) bool {
	s, ok := node.Sizes[axis]
	if !ok {
		return false
	}
	return s.Type == keys.FixSize
}

func (node *NodeItem) IsFit(axis Axis) bool {
	s, ok := node.Sizes[axis]
	if !ok {
		return false
	}
	return s.Type == keys.FitSize
}

func (node *NodeItem) IsGrow(axis Axis) bool {
	s, ok := node.Sizes[axis]
	if !ok {
		return false
	}
	return s.Type == keys.GrowSize
}

func (node *NodeItem) HasGrowChildren(axis Axis) bool {
	for _, child := range node.Children {
		if child.IsGrow(axis) {
			return true
		}
	}

	return false
}

func (node *NodeItem) SetSideByAxis(axis Axis, value float32) {
	switch axis {
	case XAxis:
		node.Box.W = value
	case YAxis:
		node.Box.H = value
	}
}

func (node *NodeItem) GetSideByAxis(axis Axis) float32 {
	switch axis {
	case XAxis:
		return node.Box.W
	case YAxis:
		return node.Box.H
	}
	return 0
}

func clampSide(axis Axis, node *NodeItem, value float32) float32 {
	s, ok := node.Sizes[axis]
	v := value

	if !ok {
		return 0
	}

	if s.Max != 0 && v > s.Max {
		v = s.Max
	}

	if s.Min > value {
		v = s.Min
	}

	return v
}

func Grow(value float32) *SizeArgument {
	return &SizeArgument{
		Key:   keys.GrowSize,
		Value: value,
	}
}

func Fix(value float32) *SizeArgument {
	return &SizeArgument{
		Key:   keys.FixSize,
		Value: value,
	}
}

func Fit() *SizeArgument {
	return &SizeArgument{Key: keys.FitSize}
}

func Min(value float32) *SizeArgument {
	return &SizeArgument{
		Key:   keys.MinSize,
		Value: value,
	}
}

func Max(value float32) *SizeArgument {
	return &SizeArgument{
		Key:   keys.MaxSize,
		Value: value,
	}
}

func Size(axis Axis, args ...*SizeArgument) *Argument {
	s := &AxisSize{
		Axis: axis,
	}

	for _, arg := range args {
		if arg.Key == keys.MinSize {
			s.Min = arg.Value
			continue
		}
		if arg.Key == keys.MaxSize {
			s.Max = arg.Value
			continue
		}
		s.Type = arg.Key
		s.Value = arg.Value
	}

	return &Argument{
		Key:   keys.SizeArg,
		Value: s,
	}
}

func Width(args ...*SizeArgument) *Argument {
	return Size(XAxis, args...)
}

func Height(args ...*SizeArgument) *Argument {
	return Size(YAxis, args...)
}

func ComputeSize(axis Axis, root *NodeItem) error {
	err := computeFix(axis, root)

	if err != nil {
		return err
	}

	err = computeFit(axis, root)

	if err != nil {
		return err
	}

	err = computeGrow(axis, root)

	return err
}

func computeFix(axis Axis, node *NodeItem) error {
	s, _ := node.Sizes[axis]

	if node.IsFix(axis) {
		node.SetSideByAxis(axis, s.Value)
		node.Computed[axis] = true
	}

	for _, child := range node.Children {
		err := computeFix(axis, child)
		if err != nil {
			return err
		}
	}

	return nil
}

func computeFit(axis Axis, node *NodeItem) error {
	if node.IsComputed(axis) || !node.IsFit(axis) {
		for _, child := range node.Children {
			computeFit(axis, child)
		}
		return nil
	}

	if node.HasGrowChildren(axis) {
		return errors.New("fit node can't have grow children")
	}

	p := node.GetPaddingByAxis(axis)

	if len(node.Children) == 0 {
		node.SetSideByAxis(axis, p)
		node.Computed[axis] = true
		return nil
	}

	var maxChild float32 = 0
	var totalChildren float32 = 0

	for _, child := range node.Children {
		computeFit(axis, child)

		if child.IsComputed(axis) {
			cs := child.GetSideByAxis(axis)

			if cs > maxChild {
				maxChild = cs
			}

			totalChildren += cs
		}
	}

	var side float32 = p

	if node.IsAlongAxis(axis) {
		side += node.Gap*float32(len(node.Children)-1) + totalChildren
	} else {
		side += maxChild
	}

	node.SetSideByAxis(axis, side)
	node.Computed[axis] = true

	return nil
}

func computeGrow(axis Axis, node *NodeItem) error {
	var err error

	if node.IsAlongAxis(axis) {
		err = growChildrenAlongAxis(axis, node)
	} else {
		err = growChildrenCrossAxis(axis, node)
	}

	for _, child := range node.Children {
		err := computeGrow(axis, child)
		if err != nil {
			return err
		}
	}

	return err
}

func growChildrenAlongAxis(axis Axis, node *NodeItem) error {
	p := node.GetPaddingByAxis(axis)
	gap := float32(len(node.Children)-1) * node.Gap
	w := node.GetSideByAxis(axis)
	taken := p + gap

	var totalShare float32 = 0

	for _, child := range node.Children {
		if !child.IsGrow(axis) {
			taken += child.GetSideByAxis(axis)
			continue
		}

		s, _ := child.Sizes[axis]

		totalShare += s.Value
	}

	available := w - taken

	for {
		if totalShare == 0 {
			return nil
		}

		changed := false

		for _, child := range node.Children {
			if child.IsComputed(axis) || !child.IsGrow(axis) {
				continue
			}

			s, _ := child.Sizes[axis]

			side := s.Value / totalShare * available

			if s.Max > 0 && side > s.Max {
				child.SetSideByAxis(axis, s.Max)
				available -= s.Max
				totalShare -= s.Value
				child.Computed[axis] = true
				changed = true
			}
		}

		if !changed {
			break
		}
	}

	for _, child := range node.Children {
		if child.IsComputed(axis) || !child.IsGrow(axis) {
			continue
		}

		s, _ := child.Sizes[axis]

		side := s.Value / totalShare * available

		child.SetSideByAxis(axis, side)
		child.Computed[axis] = true
	}

	return nil
}

func growChildrenCrossAxis(axis Axis, node *NodeItem) error {
	for _, child := range node.Children {
		if !child.IsGrow(axis) || child.IsComputed(axis) {
			continue
		}
		side := clampSide(axis, child, growCrossAxis(axis, child))
		child.SetSideByAxis(axis, side)
		child.Computed[axis] = true
	}
	return nil
}

func growCrossAxis(axis Axis, node *NodeItem) float32 {
	if node.Parent == nil || !node.Parent.IsComputed(axis) {
		return 0
	}
	return node.Parent.GetSideByAxis(axis) - node.Parent.GetPaddingByAxis(axis)
}
