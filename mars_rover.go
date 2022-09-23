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

func (rover *marsRover) move_forward() {
	rover.position_x += 1
}

func (rover *marsRover) move_backward() {
	rover.position_x -= 1
}
