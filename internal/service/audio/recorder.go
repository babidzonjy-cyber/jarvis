package audio

import (
	"os"
	"os/exec"
)

type Recorder struct {
	cmd *exec.Cmd
}

func (r *Recorder) Start(outputPath string) error {
	r.cmd = exec.Command("ffmpeg", "-f", "avfoundation", "-i", ":0", "-ar", "16000", "-ac", "1", outputPath)
	return r.cmd.Start()
}

func (r *Recorder) Stop() error {
	if err := r.cmd.Process.Signal(os.Interrupt); err != nil {
		return err
	}

	return r.cmd.Wait()
}
