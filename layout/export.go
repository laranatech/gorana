package layout

import (
	"fmt"
)

type OutputItem struct {
	Id       string       `json:"id"`
	Parent   *OutputItem  `json:"parent"`
	Children []*OutputItem `json:"children"`
	X        float32      `json:"x"`
	Y        float32      `json:"y"`
	W        float32      `json:"w"`
	H        float32      `json:"h"`
}

func Export(root *NodeItem) *OutputItem {
	node := &OutputItem{
		Id:     root.Id,
		X:      root.Box.X,
		Y:      root.Box.Y,
		W:      root.Box.W,
		H:      root.Box.H,
		Parent: nil,
	}

	children := make([]*OutputItem, 0, len(root.Children))

	for _, child := range root.Children {
		c := Export(child)
		c.Parent = node
		children = append(children, c)
	}

	node.Children = children

	return node
}

func PrintNodes(node *OutputItem) {
	fmt.Println(node.Id, node.X, node.Y, node.W, node.H)

	if len(node.Children) > 0 {
		fmt.Println(":::")

		for _, child := range node.Children {
			PrintNodes(child)
		}

		fmt.Println(";;;")
	}
}
