package transcribe

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Transcribe(wavPath, initialPrompt string) (string, error) {
	outputDir := filepath.Dir(wavPath)

	cmd := exec.Command(
		"whisper", wavPath,
		"--model", "small",
		"--language", "Russian",
		"--output_format", "txt",
		"--output_dir", outputDir,
		"--initial_prompt", initialPrompt,
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("transcribe failed: %w", err)
	}

	txtPath := strings.TrimSuffix(wavPath, ".wav") + ".txt"

	data, err := os.ReadFile(txtPath)
	if err != nil {
		return "", fmt.Errorf("cannot read transcript: %w", err)
	}

	return string(data), nil
}
