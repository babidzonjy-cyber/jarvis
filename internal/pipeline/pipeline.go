package pipeline

import (
	"fmt"
	"jarvis/internal/service/audio"
	"jarvis/internal/service/transcribe"
	"os"
)

type Pipeline struct {
	Recorder *audio.Recorder
	WavPath  string
}

func (p *Pipeline) OnDown() error {
	fmt.Println("[LISTENING]")
	fmt.Println()

	if err := p.Recorder.Start(p.WavPath); err != nil {
		return fmt.Errorf("keyDown error: %w", err)
	}

	return nil
}

func (p *Pipeline) OnUp() error {
	fmt.Println("[STOP LISTENING]")
	fmt.Println()

	if err := p.Recorder.Stop(); err != nil {
		return fmt.Errorf("keyUp error: %w", err)
	}

	data, err := transcribe.Transcribe(p.WavPath)
	if err != nil {
		return fmt.Errorf("cannot transcribe: %w", err)
	}

	fmt.Println(string(data))

	return nil
}

func (p *Pipeline) OnExit() {
	fmt.Println("Stopping programm")
	os.Exit(0)
}
