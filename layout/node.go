package layout

import (
	"larana.tech/go/gorana/layout/keys"
	"larana.tech/go/gorana/utils"
)

const (
	NodeIdLength = 8
)

type node struct {
	id        string
	parent    *node
	children  []*node
	padding   *PaddingValue
	gap       float32
	sizes     map[Axis]*AxisSize
	cube      Cube
	direction keys.DirectionKey
	alignment map[Axis]keys.AlignmentKey
	computed  map[Axis]bool
}

// Computed values

func (n *node) IsRoot() bool {
	return n.parent == nil
}

func (n *node) IsComputed(axis Axis) bool {
	c, ok := n.computed[axis]

	return ok && c
}

// Setters

func (n *node) Id(id string) *node {
	n.id = id
	return n
}

func (n *node) RandId() *node {
	n.id = utils.RandString(NodeIdLength)
	return n
}

func (n *node) Gap(value float32) *node {
	n.gap = value
	return n
}

func (n *node) Padding(args ...float32) *node {
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

	n.padding = p

	return n
}

func (n *node) Children(children ...*node) *node {
	for _, child := range children {
		child.parent = n
	}

	n.children = children

	return n
}

func (n *node) Row() *node {
	n.direction = keys.Row
	return n
}

func (n *node) Column() *node {
	n.direction = keys.Column
	return n
}

func (n *node) Stack() *node {
	n.direction = keys.Stack
	return n
}

func (n *node) Align(axis Axis, alignment keys.AlignmentKey) *node {
	n.alignment[axis] = alignment
	return n
}

// constructor

func New() *node {
	n := &node{
		sizes:     map[Axis]*AxisSize{},
		alignment: map[Axis]keys.AlignmentKey{},
		computed:  map[Axis]bool{},
		padding:   &PaddingValue{},
	}

	n.alignment[XAxis] = keys.Start
	n.alignment[YAxis] = keys.Start
	n.alignment[ZAxis] = keys.Start

	n.sizes[XAxis] = &AxisSize{Type: keys.FitSize, Axis: XAxis}
	n.sizes[YAxis] = &AxisSize{Type: keys.FitSize, Axis: YAxis}
	n.sizes[ZAxis] = &AxisSize{Type: keys.FitSize, Axis: ZAxis}

	n.id = utils.RandString(NodeIdLength)

	return n
}
