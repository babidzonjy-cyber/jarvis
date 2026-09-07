package audio

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func findDevice() (int, error) {
	cmd := exec.Command(
		"ffmpeg", "-f",
		"avfoundation", "-list_devices",
		"true", "-i", "",
	)

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return -1, fmt.Errorf("stderr pipe err: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return -1, fmt.Errorf("cannot start command: %w", err)
	}

	scanner := bufio.NewScanner(stderrPipe)

	numberIndex := -1
	for scanner.Scan() {
		line := scanner.Text()
		subStr := "MacBook Air Microphone"
		if strings.Contains(line, subStr) {
			idx := strings.LastIndex(line, "[")
			end := strings.LastIndex(line, "]")

			stringIndex := line[idx+1 : end]

			number, err := strconv.Atoi(stringIndex)
			if err != nil {
				log.Printf("cannot parse string to integer: %v\n", err)
			}

			numberIndex = number
		}
	}

	if err := scanner.Err(); err != nil {
		return -1, fmt.Errorf("scanner error: %w", err)
	}

	if numberIndex == -1 {
		return -1, fmt.Errorf("cannot find device: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		var errExit *exec.ExitError
		if errors.As(err, &errExit) {
			return numberIndex, nil
		}
		return -1, fmt.Errorf("cannot wait to stop command: %w", err)
	}

	return numberIndex, nil
}
