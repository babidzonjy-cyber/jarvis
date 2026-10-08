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

func openProject(projectName string) (string, error) {
	if !projects[projectName] {
		return "", fmt.Errorf("project %q is not allowed", projectName)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("user home dir: %w", err)
	}

	path := filepath.Join(home, "dev", projectName)

	return path, nil
}

func OpenProjectZed(projectName string) (string, error) {
	path, err := openProject(projectName)
	if err != nil {
		return "", fmt.Errorf("open project: %w", err)
	}

	return "Successfuly opened", exec.Command("zed", path).Run()
}

func OpenProjectVsCode(projectName string) (string, error) {
	path, err := openProject(projectName)
	if err != nil {
		return "", fmt.Errorf("open project: %w", err)
	}

	return "Successfuly opened", exec.Command("code", path).Run()
}

func OpenProjectNvim(projectName string) (string, error) {
	path, err := openProject(projectName)
	if err != nil {
		return "", fmt.Errorf("open project: %w", err)
	}

	return "Successfuly opened", exec.Command("nvim", path).Run()
}
