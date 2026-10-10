package computer

import (
	"fmt"
	"jarvis/internal/service/tools/computer/keyboard"
	"jarvis/internal/service/tools/computer/mouse"
)

type Action struct {
	Type string `json:"type"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Text string `json:"text"`
}

func ExecuteActionVision(action Action) error {
	switch action.Type {
	case "move":
		mouse.MoveMouse(action.X, action.Y)
	case "click":
		if err := mouse.ClickAt(action.X, action.Y); err != nil {
			return err
		}
	case "type":
		keyboard.Type(action.Text)
	default:
		return fmt.Errorf("unknown action: %s", action.Type)
	}

	return nil
}
