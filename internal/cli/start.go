package cli

import (
	"context"
	"jarvis/internal/pipeline"
	"jarvis/internal/service/audio"
	"jarvis/internal/service/hotkey"
	"log"
	"time"

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

		go func() {
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()

			for range ticker.C {
				if err := daemonStateRepo.Heartbeat(context.Background()); err != nil {
					log.Printf("heartbeat failed: %v\n", err)
				}
			}
		}()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
