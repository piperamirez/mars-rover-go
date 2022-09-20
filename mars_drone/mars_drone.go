package mars_drone

import (
	"time"
)

type MarsDrone interface {
	TakePhoto() string
	GetPhotosTaken() int
}

type marsDrone struct {
	photos_taken int
}

func NewMarsDrone() MarsDrone {
	time.Sleep(1 * time.Second)
	drone := marsDrone{
		photos_taken: 0,
	}
	return &drone
}

func (drone *marsDrone) TakePhoto() string {
	drone.photos_taken = drone.photos_taken + 1
	time.Sleep(3 * time.Second)
	return "🗾"
}

func (drone *marsDrone) GetPhotosTaken() int {
	time.Sleep(1 * time.Second)
	return drone.photos_taken
}
