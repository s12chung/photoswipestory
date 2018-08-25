package swiper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type ImageData struct {
	settings *Settings
}

func NewImageData(settings *Settings) *ImageData {
	return &ImageData{settings}
}

const orderFile = "order.txt"

func (imageData *ImageData) OrderFilenames() ([]string, error) {
	imagePath := imageData.settings.ImagePath
	bytes, err := ioutil.ReadFile(path.Join(imagePath, orderFile))
	if err != nil {
		return nil, err
	}

	var orderFilenames []string
	for _, filename := range strings.Split(string(bytes), "\n") {
		filename = strings.TrimSpace(filename)
		if filename == "" {
			continue
		}
		orderFilenames = append(orderFilenames, filename)
	}
	return orderFilenames, nil
}

func (imageData *ImageData) Paths() ([]string, error) {
	orderFilenames, err := imageData.OrderFilenames()
	if err != nil {
		return nil, err
	}

	imagePath := imageData.settings.ImagePath
	var imagePaths []string
	for _, filename := range orderFilenames {
		filename = strings.TrimSpace(filename)
		_, err := os.Stat(path.Join(imagePath, filename))
		if err != nil {
			return nil, err
		}
		imagePaths = append(imagePaths, filename)
	}

	dirFilePaths, err := ioutil.ReadDir(imagePath)
	if err != nil {
		return nil, err
	}
	if len(imagePaths) != len(dirFilePaths)-1 {
		return nil, fmt.Errorf("image count in %v (%v) do not match count in %v (%v)", orderFile, len(imagePaths), imagePath, len(dirFilePaths))
	}
	return imagePaths, nil
}
