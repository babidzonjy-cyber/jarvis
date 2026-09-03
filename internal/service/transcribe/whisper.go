package transcribe

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Transcribe(wavPath string) (string, error) {
	cmd := exec.Command("whisper", wavPath, "--model", "base", "--language", "Russian", "--output_format", "txt")

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
