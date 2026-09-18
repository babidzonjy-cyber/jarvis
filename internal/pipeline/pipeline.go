package pipeline

import (
	"fmt"
	"jarvis/internal/service/audio"
	llm "jarvis/internal/service/llm/generate"
	"jarvis/internal/service/transcribe"
	"os"
)

type Pipeline struct {
	Recorder *audio.Recorder
	WavPath  string
}

func (p *Pipeline) OnDown() error {
	fmt.Println("[LISTENING]")

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

	systemPrompt := "Голосовой помощник, разговор на русском языке."

	prompt, err := transcribe.Transcribe(p.WavPath, systemPrompt)
	if err != nil {
		return fmt.Errorf("cannot transcribe: %w", err)
	}

	fmt.Println("Ваш текст:", prompt)

	answer, err := llm.Ask(prompt, systemPrompt)
	if err != nil {
		return fmt.Errorf("llm cannot ask: %w", err)
	}

	fmt.Println("Ответ нейросети:", answer)
	fmt.Println()

	return nil
}

func (p *Pipeline) OnExit() {
	fmt.Println("Stopping programm")
	os.Exit(0)
}
