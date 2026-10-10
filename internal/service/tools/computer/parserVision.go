package computer

import (
	"encoding/json"
	"fmt"
)

func ParseAction(response string) (Action, error) {
	var action Action

	fmt.Printf("VISION RAW: %q\n", response)

	if err := json.Unmarshal([]byte(response), &action); err != nil {
		return Action{}, fmt.Errorf("cannot unmarshal reponse: %w", err)
	}

	return action, nil
}
