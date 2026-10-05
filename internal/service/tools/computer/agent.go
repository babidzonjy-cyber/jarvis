package computer

import (
	"fmt"
	"jarvis/internal/service/llm"
	"jarvis/internal/service/tools/computer/screen"
	"os"
	"path/filepath"
)

func Run(task string) error {
	raw := os.Getenv("SCREENSHOT")
	if raw == "" {
		return fmt.Errorf("SCREENSHOTPATH is empty")
	}

	path, err := filepath.Abs(raw)
	if err != nil {
		return fmt.Errorf("path error: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	img, err := screen.Screenshot()
	if err != nil {
		return fmt.Errorf("cannot screen %w", err)
	}

	if err := screen.SaveImage(img, path); err != nil {
		return fmt.Errorf("save screenshot: %w", err)
	}

	answer, err := llm.AskVision(task, path)
	if err != nil {
		return fmt.Errorf("vision: %w", err)
	}

	fmt.Println(answer)

	return nil
}
