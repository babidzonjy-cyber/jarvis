package main

import (
	"fmt"
	"jarvis/internal/pipeline"
	"jarvis/internal/service/audio"
	"jarvis/internal/service/hotkey"
	"log"

	"golang.design/x/hotkey/mainthread"
)

func main() {
	mainthread.Init(run)
}

func run() {
	outputPath := "cmd/test.wav"

	pipeline := &pipeline.Pipeline{
		Recorder: &audio.Recorder{},
		WavPath:  outputPath,
	}

	fmt.Println("Start")

	go func() {
		if err := hotkey.ListenExit(pipeline.OnExit); err != nil {
			log.Fatal(err)
		}
	}()

	if err := hotkey.ListenPushToTalk(pipeline.OnDown, pipeline.OnUp); err != nil {
		log.Fatal(err)
	}
}
