package layout

import (
	"larana.tech/go/gorana/layout/keys"
	"larana.tech/go/gorana/utils"
)

const (
	NodeIdLength = 8
)

type Node struct {
	id        string
	parent    *Node
	children  []*Node
	padding   *PaddingValue
	gap       float32
	sizes     map[Axis]*AxisSize
	cube      Cube
	direction keys.DirectionKey
	alignment map[Axis]keys.AlignmentKey
	computed  map[Axis]bool
}

// Computed values

func (n *Node) IsRoot() bool {
	return n.parent == nil
}

func (n *Node) IsComputed(axis Axis) bool {
	c, ok := n.computed[axis]

	return ok && c
}

// Setters

func (n *Node) Id(id string) *Node {
	n.id = id
	return n
}

func (n *Node) RandId() *Node {
	n.id = utils.RandString(NodeIdLength)
	return n
}

func (n *Node) Gap(value float32) *Node {
	n.gap = value
	return n
}

func (n *Node) Padding(args ...float32) *Node {
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

func (n *Node) Children(children ...*Node) *Node {
	for _, child := range children {
		child.parent = n
	}

	n.children = children

	return n
}

func (n *Node) Row() *Node {
	n.direction = keys.Row
	return n
}

func (n *Node) Column() *Node {
	n.direction = keys.Column
	return n
}

func (n *Node) Stack() *Node {
	n.direction = keys.Stack
	return n
}

func (n *Node) Align(axis Axis, alignment keys.AlignmentKey) *Node {
	n.alignment[axis] = alignment
	return n
}

// constructor

func New() *Node {
	n := &Node{
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
