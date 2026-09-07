package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var modeCmd = &cobra.Command{
	Use:   "mode [work|chill]",
	Short: "swithes modes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mode := args[0]
		fmt.Println("mode:", mode)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(modeCmd)
}
