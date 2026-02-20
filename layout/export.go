package layout

type OutputItem struct {
	Id       string        `json:"id"`
	Parent   *OutputItem   `json:"parent"`
	Children []*OutputItem `json:"children"`
	X        float32       `json:"x"`
	Y        float32       `json:"y"`
	W        float32       `json:"w"`
	H        float32       `json:"h"`
}

func (n *node) Export() *OutputItem {
	res := &OutputItem{
		Id:     n.id,
		X:      n.box.X,
		Y:      n.box.Y,
		W:      n.box.W,
		H:      n.box.H,
		Parent: nil,
	}

	children := make([]*OutputItem, 0, len(n.children))

	for _, child := range n.children {
		c := child.Export()
		c.Parent = res
		children = append(children, c)
	}

	res.Children = children

	return res
}
