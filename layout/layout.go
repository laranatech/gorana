package layout

import (
	"github.com/google/uuid"
	"github.com/laranatech/gorana/layout/keys"
)

type Axis byte

const (
	XAxis Axis = iota
	YAxis
)

type Box struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	W float32 `json:"w"`
	H float32 `json:"h"`
}

type Argument struct {
	Key   keys.ArgumentKey
	Value any
}

type NodeItem struct {
	Id        string
	Parent    *NodeItem
	Children  []*NodeItem
	Padding   *PaddingValue
	Gap       float32
	Sizes     map[Axis]*AxisSize
	Box       Box
	Direction keys.DirectionKey
	Alignment keys.AlignmentKey
	Computed  map[Axis]bool
}

func (node *NodeItem) IsRoot() bool {
	return node.Parent == nil
}

func (node *NodeItem) IsComputed(axis Axis) bool {
	c, ok := node.Computed[axis]

	return ok && c
}

func Gap(value float32) *Argument {
	return &Argument{
		Key:   keys.GapArg,
		Value: value,
	}
}

func Node(args ...*Argument) *NodeItem {
	node := &NodeItem{
		Sizes:    map[Axis]*AxisSize{},
		Computed: map[Axis]bool{},
	}
	for _, arg := range args {
		switch arg.Key {
		case keys.GapArg:
			node.Gap = arg.Value.(float32)
		case keys.PaddingArg:
			node.Padding = arg.Value.(*PaddingValue)
		case keys.ChildrenArg:
			children := arg.Value.([]*NodeItem)
			for _, child := range children {
				child.Parent = node
			}
			node.Children = children
		case keys.SizeArg:
			s := arg.Value.(*AxisSize)
			node.Sizes[s.Axis] = s
		case keys.DirectionArg:
			node.Direction = arg.Value.(keys.DirectionKey)
		case keys.AlignmentArg:
			node.Alignment = arg.Value.(keys.AlignmentKey)
		case keys.IdArg:
			node.Id = arg.Value.(string)
		}
	}

	_, ok := node.Sizes[XAxis]

	if !ok {
		node.Sizes[XAxis] = &AxisSize{Type: keys.FitSize, Axis: XAxis}
	}

	_, ok = node.Sizes[YAxis]

	if !ok {
		node.Sizes[YAxis] = &AxisSize{Type: keys.FitSize, Axis: YAxis}
	}

	if node.Id == "" {
		// TODO: uuid.NewString() panics. handle error later
		node.Id = uuid.NewString()
	}

	return node
}

func Children(children ...*NodeItem) *Argument {
	return &Argument{
		Key:   keys.ChildrenArg,
		Value: children,
	}
}

func Id(value string) *Argument {
	return &Argument{
		Key:   keys.IdArg,
		Value: value,
	}
}

func Start() *Argument {
	return &Argument{
		Key: keys.AlignmentArg,
		Value: keys.Start,
	}
}

func Center() *Argument {
	return &Argument{
		Key: keys.AlignmentArg,
		Value: keys.Center,
	}
}

func End() *Argument {
	return &Argument{
		Key: keys.AlignmentArg,
		Value: keys.End,
	}
}

func Layout(root *NodeItem) error {
	if err := ComputeSize(XAxis, root); err != nil {
		return err
	}

	if err := ComputeSize(YAxis, root); err != nil {
		return err
	}

	if err := ComputePosition(XAxis, root); err != nil {
		return err
	}

	if err := ComputePosition(YAxis, root); err != nil {
		return err
	}

	return nil
}

func Test() {
	root := Node(
		Id("root"),
		Row(),
		Width(Fix(640)),
		Height(Fix(480)),
		Gap(16),
		Padding(2),
		Children(
			Node(
				Id("child_1"),
				Width(Grow(1)),
				Height(Grow(1)),
			),
			Node(
				Id("child_2"),
				Width(Grow(2), Max(150)),
				Height(Grow(1)),
			),
			Node(
				Id("child_3"),
				Width(Fit()),
				Gap(8),
				Padding(2),
				Children(
					Node(
						Id("grandchild_1"),
						Width(Fix(50)),
						Height(Fix(50)),
					),
					Node(
						Id("grandchild_2"),
						Width(Fix(15)),
						Height(Fix(15)),
					),
				),
			),
		),
	)

	Layout(root)

	output := Export(root)

	PrintNodes(output)
}
