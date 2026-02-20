package layout

type PaddingValue struct {
	Top    float32 `json:"t"`
	Bottom float32 `json:"b"`
	Left   float32 `json:"l"`
	Right  float32 `json:"r"`
}

func (n *node) GetPaddingByAxis(axis Axis) float32 {
	if n.padding == nil {
		return 0
	}

	if axis == XAxis {
		return n.padding.Left + n.padding.Right
	}
	return n.padding.Bottom + n.padding.Top
}

func (n *node) GetInitialPaddingByAxis(axis Axis) float32 {
	if n.padding == nil {
		return 0
	}

	if axis == XAxis {
		return n.padding.Left
	}
	return n.padding.Top
}
