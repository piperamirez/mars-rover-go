package mars_rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarsRoverExists(t *testing.T) {
	mars_rover := NewMarsRover()
	assert.NotNil(t, mars_rover)
}

func TestMarsRoverMoveForward(t *testing.T) {
	// Were assuming that the mars rover is
	// facing north at (0, 0)
	mars_rover := NewMarsRover()
	mars_rover.MoveForward()
	assert.Equal(t, mars_rover.y, 1)
	assert.Equal(t, mars_rover.x, 0)
	assert.Equal(t, mars_rover.direction, "north")
}
