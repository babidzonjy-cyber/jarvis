package computer

import (
	"fmt"
	"jarvis/internal/service/llm"
)

func RunVisionAction(task string, imagePath string) error {
	jsonStrData, err := llm.AskVision(task, imagePath)
	fmt.Println(jsonStrData)
	if err != nil {
		return fmt.Errorf("ask vision: %w", err)
	}

	actionVision, err := ParseAction(jsonStrData)
	if err != nil {
		return fmt.Errorf("parse action: %w", err)
	}

	if err := ExecuteActionVision(actionVision); err != nil {
		return fmt.Errorf("execute action vision: %w", err)
	}

	return nil
}
