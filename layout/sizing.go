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

func (n *node) IsFix(axis Axis) bool {
	s, ok := n.sizes[axis]
	if !ok {
		return false
	}
	return s.Type == keys.FixSize
}

func (n *node) IsFit(axis Axis) bool {
	s, ok := n.sizes[axis]
	if !ok {
		return false
	}
	return s.Type == keys.FitSize
}

func (n *node) IsGrow(axis Axis) bool {
	s, ok := n.sizes[axis]
	if !ok {
		return false
	}
	return s.Type == keys.GrowSize
}

func (n *node) HasGrowChildren(axis Axis) bool {
	for _, child := range n.children {
		if child.IsGrow(axis) {
			return true
		}
	}

	return false
}

func (n *node) SetSideByAxis(axis Axis, value float32) {
	switch axis {
	case XAxis:
		n.box.W = value
	case YAxis:
		n.box.H = value
	}
}

func (n *node) GetSideByAxis(axis Axis) float32 {
	switch axis {
	case XAxis:
		return n.box.W
	case YAxis:
		return n.box.H
	}
	return 0
}

func clampSide(axis Axis, n *node, value float32) float32 {
	s, ok := n.sizes[axis]
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

func (n *node) Size(axis Axis, args ...*SizeArgument) *node {
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

	n.sizes[axis] = s

	return n
}

func (n *node) Width(args ...*SizeArgument) *node {
	return n.Size(XAxis, args...)
}

func (n *node) Height(args ...*SizeArgument) *node {
	return n.Size(YAxis, args...)
}

// func (n *node) Depth(args ...*SizeArgument) *node {
// 	return n.Size(ZAxis, args...)
// }

func ComputeSize(axis Axis, root *node) error {
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

func computeFix(axis Axis, n *node) error {
	s, _ := n.sizes[axis]

	if n.IsFix(axis) {
		n.SetSideByAxis(axis, s.Value)
		n.computed[axis] = true
	}

	for _, child := range n.children {
		err := computeFix(axis, child)
		if err != nil {
			return err
		}
	}

	return nil
}

func computeFit(axis Axis, n *node) error {
	if n.IsComputed(axis) || !n.IsFit(axis) {
		for _, child := range n.children {
			computeFit(axis, child)
		}
		return nil
	}

	if n.HasGrowChildren(axis) {
		return errors.New("fit n can't have grow children")
	}

	p := n.GetPaddingByAxis(axis)

	if len(n.children) == 0 {
		n.SetSideByAxis(axis, p)
		n.computed[axis] = true
		return nil
	}

	var maxChild float32 = 0
	var totalChildren float32 = 0

	for _, child := range n.children {
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

	if n.IsAlongAxis(axis) {
		side += n.gap*float32(len(n.children)-1) + totalChildren
	} else {
		side += maxChild
	}

	n.SetSideByAxis(axis, side)
	n.computed[axis] = true

	return nil
}

func computeGrow(axis Axis, n *node) error {
	var err error

	if n.IsAlongAxis(axis) {
		err = growChildrenAlongAxis(axis, n)
	} else {
		err = growChildrenCrossAxis(axis, n)
	}

	for _, child := range n.children {
		err := computeGrow(axis, child)
		if err != nil {
			return err
		}
	}

	return err
}

func growChildrenAlongAxis(axis Axis, n *node) error {
	p := n.GetPaddingByAxis(axis)
	gap := float32(len(n.children)-1) * n.gap
	w := n.GetSideByAxis(axis)
	taken := p + gap

	var totalShare float32 = 0

	for _, child := range n.children {
		if !child.IsGrow(axis) {
			taken += child.GetSideByAxis(axis)
			continue
		}

		s, _ := child.sizes[axis]

		totalShare += s.Value
	}

	available := w - taken

	for {
		if totalShare == 0 {
			return nil
		}

		changed := false

		for _, child := range n.children {
			if child.IsComputed(axis) || !child.IsGrow(axis) {
				continue
			}

			s, _ := child.sizes[axis]

			side := s.Value / totalShare * available

			if s.Max > 0 && side > s.Max {
				child.SetSideByAxis(axis, s.Max)
				available -= s.Max
				totalShare -= s.Value
				child.computed[axis] = true
				changed = true
			}
		}

		if !changed {
			break
		}
	}

	for _, child := range n.children {
		if child.IsComputed(axis) || !child.IsGrow(axis) {
			continue
		}

		s, _ := child.sizes[axis]

		side := s.Value / totalShare * available

		child.SetSideByAxis(axis, side)
		child.computed[axis] = true
	}

	return nil
}

func growChildrenCrossAxis(axis Axis, n *node) error {
	for _, child := range n.children {
		if !child.IsGrow(axis) || child.IsComputed(axis) {
			continue
		}
		side := clampSide(axis, child, growCrossAxis(axis, child))
		child.SetSideByAxis(axis, side)
		child.computed[axis] = true
	}
	return nil
}

func growCrossAxis(axis Axis, n *node) float32 {
	if n.parent == nil || !n.parent.IsComputed(axis) {
		return 0
	}
	return n.parent.GetSideByAxis(axis) - n.parent.GetPaddingByAxis(axis)
}
