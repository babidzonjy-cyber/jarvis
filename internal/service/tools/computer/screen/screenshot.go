package screen

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/go-vgo/robotgo"
)

func Screenshot() (img image.Image, err error) {
	bitmap := robotgo.CaptureScreen()
	defer robotgo.FreeBitmap(bitmap)

	img = robotgo.ToImage(bitmap)
	if img == nil {
		return nil, errors.New("failed to capture screen")
	}

	return img, nil
}

func SaveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("encode image to file: %w", err)
	}

	return nil
}
