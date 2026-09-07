package main

import (
	"jarvis/internal/cli"
	"os"

	"golang.design/x/hotkey/mainthread"
)

func main() {
	mainthread.Init(func() {
		if err := cli.Execute(); err != nil {
			os.Exit(1)
		}
	})
}
