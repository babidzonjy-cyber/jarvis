package screen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScreenshot(t *testing.T) {
	_, err := Screenshot()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSaveImage(t *testing.T) {
	img, err := Screenshot()
	if err != nil {
		t.Fatal(err)
	}

	path, err := filepath.Abs(os.Getenv("SCREENSHOTPATH"))
	if err != nil {
		t.Fatal(err)
	}

	if err := SaveImage(img, path); err != nil {
		t.Fatal(err)
	}
}
