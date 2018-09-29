package tennis

func IsPointSliceEqual(a, b []Point) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

type Point bool

func (p Point) IsForPlayerOne() bool {
	return bool(p)
}

func (p Point) IsForPlayerTwo() bool {
	return !bool(p)
}

const (
	PlayerOnePoint Point = true
	PlayerTwoPoint Point = false
)
