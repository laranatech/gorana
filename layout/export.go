package layout

type OutputItem struct {
	Id       string        `json:"id"`
	Parent   *OutputItem   `json:"parent"`
	Children []*OutputItem `json:"children"`
	X        float32       `json:"x"`
	Y        float32       `json:"y"`
	Z        float32       `json:"z"`
	W        float32       `json:"w"`
	H        float32       `json:"h"`
	D        float32       `json:"d"`
}

func (n *node) Export() *OutputItem {
	res := &OutputItem{
		Id:     n.id,
		X:      n.cube.X,
		Y:      n.cube.Y,
		Z:      n.cube.Z,
		W:      n.cube.W,
		H:      n.cube.H,
		D:      n.cube.D,
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
