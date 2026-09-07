package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "jarvis",
	Short: "The AI voice assistant",
	Long:  "The AI voice assistant that runs local on youy PC",
}

func Execute() error {
	return rootCmd.Execute()
}
