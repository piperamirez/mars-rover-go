package mars_drone

type MarsDrone interface {
	TakePhoto() string
	GetPhotosTaken() int
}

type marsDrone struct {
	photos_taken int
}

func NewMarsDrone() MarsDrone {
	drone := marsDrone{
		photos_taken: 0,
	}
	return &drone
}

func (drone *marsDrone) TakePhoto() string {
	drone.photos_taken = drone.photos_taken + 1
	return "🗾"
}

func (drone *marsDrone) GetPhotosTaken() int {
	return drone.photos_taken
}
