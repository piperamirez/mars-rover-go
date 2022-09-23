package mars_rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarsRoverExists(t *testing.T) {
	rover := NewMarsRover("E")
	assert.NotNil(t, rover)
}

func TestMarsRoverMoves(t *testing.T) {
	rover := NewMarsRover("E")
	rover.move_forward()

	assert.Equal(t, rover.position_x, 1)
	assert.Equal(t, rover.position_y, 0)
}

func TestMarsRoverMoveBackward(t *testing.T) {
	rover := NewMarsRover("E")
	rover.move_backward()

	assert.Equal(t, rover.position_x, -1)
	assert.Equal(t, rover.position_y, 0)
}

func TestMarsRoverStartsFacingNorth(t *testing.T) {
	rover := NewMarsRover("N")
	assert.Equal(t, "N", rover.orientation)
}
