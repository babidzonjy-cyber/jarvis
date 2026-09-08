package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var modeCmd = &cobra.Command{
	Use:   "mode [work|chill]",
	Short: "swithes modes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mode := args[0]

		if err := daemonStateRepo.SetMode(context.Background(), mode); err != nil {
			return fmt.Errorf("cannot set mode: %w", err)
		}

		fmt.Println("mode set:", mode)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(modeCmd)
}
