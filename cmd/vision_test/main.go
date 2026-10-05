package main

import (
	"fmt"
	"jarvis/internal/service/llm"
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

	answer, err := llm.AskVision("Что находится на этой картинке? Опиши кратко", path)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer)
}
