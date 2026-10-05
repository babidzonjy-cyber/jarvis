package tools

import (
	"fmt"
	"os/exec"
)

var allowedApps = map[string]bool{
	"Yandex Music":       true,
	"Telegram":           true,
	"Claude":             true,
	"Zen":                true,
	"Zed":                true,
	"Visual Studio Code": true,
	"Ghostty":            true,
	"Happ":               true,
	"OrbStack":           true,
	"Notes":              true,
	"Activity Monitor":   true,
	"TablePlus":          true,
	"Bruno":              true,
	"System Settings":    true,
	"Obsidian":           true,
}

func OpenApp(appName string) error {
	if !allowedApps[appName] {
		return fmt.Errorf("applicaton %q is not allowed", appName)
	}

	return exec.Command("open", "-a", appName).Run()
}
