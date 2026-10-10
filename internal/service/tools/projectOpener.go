package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var projects = map[string]bool{
	"leetcode": true,
	"jarvis":   true,
	"webchat":  true,
	"study":    true,
}

var ideCommands = map[string][]string{
	"vscode": {"open", "-a", "Visual Studio Code"},
	"zed":    {"open", "-a", "Zed"},
	"cursor": {"open", "-a", "Cursor"},
}

func OpenProject(projectName, ideName string) (string, error) {
	if !projects[projectName] {
		return "", fmt.Errorf("project %q is not allowed", projectName)
	}

	ide, ok := ideCommands[ideName]
	if !ok {
		return "", fmt.Errorf("unknown ide: %s", ideName)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("user home dir: %w", err)
	}

	path := filepath.Join(home, "dev", projectName)

	_, err = os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("file stat: %w", err)
	}

	cmd := exec.Command(ide[0], ide[1], ide[2], path)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cannot run command: %w, ide: %s", err, ide[2])
	}

	return "Successfuly opened", nil
}
