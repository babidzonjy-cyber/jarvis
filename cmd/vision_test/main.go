package main

import (
	"jarvis/internal/service/tools/computer"
	"jarvis/internal/service/tools/computer/screen"
	"log"
	"os"
	"path/filepath"
)

func main() {
	raw := os.Getenv("SCREENSHOTPATH")
	if raw == "" {
		log.Fatal("SCREENSHOTPATH is empty (check .env / make export)")
	}

	path, err := filepath.Abs(raw)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.Fatal(err)
	}

	img, err := screen.Screenshot()
	if err != nil {
		log.Fatal(err)
	}

	if err := screen.SaveImage(img, path); err != nil {
		log.Fatal(err)
	}

	task := "Найди файл pipeline.go и верни координаты его центра. и кликни по нему"
	if err := computer.RunVisionAction(task, path); err != nil {
		log.Fatal(err)
	}
}
