package mars_rover

type marsRover struct {
	position_x  int
	position_y  int
	orientation string
}

func NewMarsRover(orientation string) marsRover {
	rover := marsRover{
		position_x: 0, position_y: 0, orientation: orientation,
	}
	return rover
}

func (rover *marsRover) moveForward() {
	rover.move(1)
}

func (rover *marsRover) moveBackward() {
	rover.move(-1)
}

func (rover *marsRover) move(direction int) {
	switch rover.orientation {
	case "S":
		rover.position_y -= direction
	case "N":
		rover.position_y += direction
	case "E":
		rover.position_x += direction
	case "W":
		rover.position_x -= direction
	}
}

func (rover *marsRover) turnLeft() {
	// TODO
	// Find the index of current orientation in "NESW" and return the
	// value of index - 1
}
