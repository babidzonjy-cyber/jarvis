package audio

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type Recorder struct {
	cmd *exec.Cmd
}

func (r *Recorder) Start(outputPath string) error {
	deviceNumber, err := findDevice()
	if err != nil {
		fmt.Println("Cannot find the device, use the default input device")
		deviceNumber = 0
	}

	deviceStr := ":" + strconv.Itoa(deviceNumber)

	r.cmd = exec.Command(
		"ffmpeg", "-y",
		"-f", "avfoundation",
		"-i", deviceStr, "-ar",
		"16000", "-ac",
		"1", outputPath,
	)

	// r.cmd.Stderr = os.Stderr
	// r.cmd.Stdout = os.Stdout

	return r.cmd.Start()
}

func (r *Recorder) Stop() error {
	if err := r.cmd.Process.Signal(os.Interrupt); err != nil {
		return err
	}

	if err := r.cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return nil
		}
		return err
	}

	return nil
}
