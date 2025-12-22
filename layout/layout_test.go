package layout_test

import (
	"fmt"
	"math"
	"strings"
	"testing"

	la "github.com/laranatech/gorana/layout"
)

func TestLayout(t *testing.T) {
	root := la.Node(
		la.Id("root"),
		la.Row(),
		la.Width(la.Fix(640)),
		la.Height(la.Fix(480)),
		la.Gap(16),
		la.Padding(2),
		la.Children(
			la.Node(
				la.Id("child_1"),
				la.Width(la.Grow(1)),
				la.Height(la.Grow(1)),
			),
			la.Node(
				la.Id("child_2"),
				la.Width(la.Grow(2), la.Max(150)),
				la.Height(la.Grow(1)),
			),
			la.Node(
				la.Id("child_3"),
				la.Width(la.Fit()),
				la.Gap(8),
				la.Padding(2),
				la.Children(
					la.Node(
						la.Id("grandchild_1"),
						la.Width(la.Fix(50)),
						la.Height(la.Fix(50)),
					),
					la.Node(
						la.Id("grandchild_2"),
						la.Width(la.Fix(15)),
						la.Height(la.Fix(15)),
					),
				),
			),
		),
	)

	la.Layout(root)

	result := la.Export(root)

	expected := &la.OutputItem{
		Id: "root",
		X:  0,
		Y:  0,
		W:  640,
		H:  480,
		Children: []*la.OutputItem{
			{
				Id: "child_1",
				X:  2,
				Y:  2,
				W:  377,
				H:  476,
			},
			{
				Id: "child_2",
				X:  395,
				Y:  2,
				W:  150,
				H:  476,
			},
			{
				Id: "child_3",
				X:  561,
				Y:  2,
				W:  77,
				H:  54,
				Children: []*la.OutputItem{
					{
						Id: "grandchild_1",
						X:  563,
						Y:  4,
						W:  50,
						H:  50,
					},
					{
						Id: "grandchild_1",
						X:  621,
						Y:  4,
						W:  15,
						H:  15,
					},
				},
			},
		},
	}

	res := matchNodes(expected, result)

	if len(res) == 0 {
		return
	}

	t.Error(strings.Join(res, "\n"))
}

func matchNodes(a, b *la.OutputItem) []string {
	res := make([]string, 0, 100)

	if !matchFloats(a.X, b.X) {
		res = append(res, fmt.Sprintf("%s !x: %f != %f", a.Id, a.X, b.X))
	}

	if !matchFloats(a.Y, b.Y) {
		res = append(res, fmt.Sprintf("%s !y: %f != %f", a.Id, a.Y, b.Y))
	}

	if !matchFloats(a.W, b.W) {
		res = append(res, fmt.Sprintf("%s !w: %f != %f", a.Id, a.W, b.W))
	}

	if !matchFloats(a.H, b.H) {
		res = append(res, fmt.Sprintf("%s !h: %f != %f", a.Id, a.H, b.H))
	}

	if len(a.Children) != len(b.Children) {
		res = append(
			res,
			fmt.Sprintf(
				"%s !children: %d != %d",
				a.Id,
				len(a.Children),
				len(b.Children),
			),
		)
		return res
	}

	for i := range a.Children {
		r := matchNodes(a.Children[i], b.Children[i])
		res = append(res, r...)
	}

	return res
}

func matchFloats(a, b float32) bool {
	var t float32 = 0.01
	d := math.Abs(float64(a - b))

	return float32(d) < t
}
