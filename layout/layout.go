package layout

type Axis byte

const (
	XAxis Axis = iota
	YAxis
	ZAxis
)

type Box struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	W float32 `json:"w"`
	H float32 `json:"h"`
}

type Cube struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	W float32 `json:"w"`
	H float32 `json:"h"`
	D float32 `json:"d"`
}

func (n *node) Layout() error {
	if err := ComputeSize(XAxis, n); err != nil {
		return err
	}

	if err := ComputeSize(YAxis, n); err != nil {
		return err
	}

	if err := ComputeSize(ZAxis, n); err != nil {
		return err
	}

	if err := ComputePosition(XAxis, n); err != nil {
		return err
	}

	if err := ComputePosition(YAxis, n); err != nil {
		return err
	}

	if err := ComputePosition(ZAxis, n); err != nil {
		return err
	}

	return nil
}
