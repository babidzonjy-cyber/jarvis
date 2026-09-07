package cli

import (
	"jarvis/internal/pipeline"
	"jarvis/internal/service/audio"
	"jarvis/internal/service/hotkey"
	"log"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "running AI assistant",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputPath := "cmd/test.wav"

		p := &pipeline.Pipeline{
			Recorder: &audio.Recorder{},
			WavPath:  outputPath,
		}

		go func() {
			if err := hotkey.ListenExit(p.OnExit); err != nil {
				log.Printf("cannot exit: %v", err)
			}
		}()

		if err := hotkey.ListenPushToTalk(p.OnDown, p.OnUp); err != nil {
			log.Printf("cannot listen: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
