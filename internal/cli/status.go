package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "shows which mode is currently active",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		const heartbeatTimeout = 10 * time.Second

		state, err := daemonStateRepo.GetState(context.Background())
		if err != nil {
			log.Printf("cannot get state: %v", err)
			return fmt.Errorf("cannot get state: %w", err)
		}

		if time.Since(state.UpdatedAt) > heartbeatTimeout {
			fmt.Printf("Демон сейчас не запущен, последний запущенный режим был: %q, %q\n", state.Mode, state.UpdatedAt)
		} else {
			fmt.Printf("Демон работает, режим: %q\n", state.Mode)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
