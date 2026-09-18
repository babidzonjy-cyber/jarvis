package cli

import (
	"bufio"
	"context"
	"fmt"
	"jarvis/internal/pipeline"
	"jarvis/internal/service/audio"
	"jarvis/internal/service/hotkey"
	"log"
	"os"
	"strings"
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

		state, err := daemonStateRepo.GetState(context.Background())
		if err != nil {
			return fmt.Errorf("cannot get state: %w", err)
		}

		fmt.Printf("Current mode is %q. Press Space to continue, or type another mode(work/chill) to switch.\n", state.Mode)

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)

			if input != "work" && input != "chill" {
				fmt.Println("incorret mode, keep the previous one:", state.Mode)
			} else if input == "" {
				fmt.Println("we're keeping the same mode:", state.Mode)
			} else {
				if err := daemonStateRepo.SetMode(context.Background(), input); err != nil {
					return err
				}
			}
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("scanner: %w", err)
		}

		fmt.Println("Jarvis is running. Hold down Cmd+Z to record, Option+Q to exit.")
		fmt.Println()

		go func() {
			if err := hotkey.ListenExit(p.OnExit); err != nil {
				log.Printf("cannot exit: %v", err)
			}
		}()

		go func() {
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()

			for range ticker.C {
				if err := daemonStateRepo.Heartbeat(context.Background()); err != nil {
					log.Printf("heartbeat failed: %v\n", err)
				}
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
