package robot

const MaxPositionWidth = 25
const MaxPositionHeight = 10

type Position struct {
	X, Y int
}

func NewPosition() *Position {
	position := Position{0, 0}
	return &position
}
