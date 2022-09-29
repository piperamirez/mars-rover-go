package mars_rover

type MarsRover struct {
	x         int
	y         int
	direction string
}

func NewMarsRover() MarsRover {
	mars_rover := MarsRover{0, 0, "north"}
	return mars_rover
}

func (mars_rover *MarsRover) MoveForward() {
	mars_rover.y++
}
