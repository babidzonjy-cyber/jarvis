package tools

import (
	"fmt"
	"os/exec"
)

func InspectYandexMusic() error {
	script := `
tell application "System Events"
	tell process "Yandex Music"
		get name of every window
	end tell
end tell
`

	out, err := exec.Command("osascript", "-e", script).CombinedOutput()
	if err != nil {
		return fmt.Errorf("inspect yandex music: %w: %s", err, out)
	}

	fmt.Println(string(out))

	return nil
}
