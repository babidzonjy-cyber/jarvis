package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "shows which mode is currently active",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("какой статус сейчас")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
