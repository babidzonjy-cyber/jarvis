package llm

import (
	"encoding/base64"
	"fmt"
	"os"
)

func AskVision(prompt string, imagePath string) (string, error) {
	data, err := os.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("read image: %w", err)
	}

	imageBased64 := base64.StdEncoding.EncodeToString(data)

	req := RequestLLMToolCalling{
		Model: "qwen2.5vl:3b",
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
				Images:  []string{imageBased64},
			},
		},
		Stream:  false,
		Options: map[string]any{"num_ctx": 8192},
	}

	resp, err := wrapper(req)
	if err != nil {
		return "", fmt.Errorf("vision request: %w", err)
	}

	return resp.Message.Content, nil
}
