package swiper

type Settings struct {
	ImagePath string `json:"image_path,omitempty"`
}

func DefaultSettings() *Settings {
	return &Settings{
		"./content/markdowns/swiper",
	}
}
