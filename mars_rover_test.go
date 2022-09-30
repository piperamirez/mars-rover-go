package mars_rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarsRoverExists(t *testing.T) {
	rover := NewMarsRover("E")
	assert.NotNil(t, rover)
}

func TestMarsRoverMovesForward(t *testing.T) {
	rover := NewMarsRover("E")
	rover.moveForward()

	assert.Equal(t, 1, rover.position_x)
	assert.Equal(t, 0, rover.position_y)
}

func TestMarsRoverMovesBackward(t *testing.T) {
	rover := NewMarsRover("E")
	rover.moveBackward()

	assert.Equal(t, -1, rover.position_x)
	assert.Equal(t, 0, rover.position_y)
}

func TestMarsRoverMovesBackwardFacingSouth(t *testing.T) {
	rover := NewMarsRover("S")
	rover.moveBackward()

	assert.Equal(t, 0, rover.position_x)
	assert.Equal(t, 1, rover.position_y)
}

func TestMarsRoverStartsFacingNorth(t *testing.T) {
	rover := NewMarsRover("N")
	assert.Equal(t, "N", rover.orientation)
}

func TestMarsRoverMovesForwardFacingSouth(t *testing.T) {
	rover := NewMarsRover("S")
	rover.moveForward()

	assert.Equal(t, 0, rover.position_x)
	assert.Equal(t, -1, rover.position_y)
}

func TestMarsRoverMovesForwardFacingWest(t *testing.T) {
	rover := NewMarsRover("W")
	rover.moveForward()

	assert.Equal(t, -1, rover.position_x)
	assert.Equal(t, 0, rover.position_y)
}

func TestMarsRoverTurnLeft(t *testing.T) {
	rover := NewMarsRover("E")
	rover.turnLeft()

	assert.Equal(t, "N", rover.orientation)
}

func TestMarsRoverTurnsLeftFacingSouth(t *testing.T) {
	rover := NewMarsRover("S")
	rover.turnLeft()

	assert.Equal(t, "E", rover.orientation)
}
