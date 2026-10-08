package llm

import (
	"encoding/base64"
	"fmt"
	"os"
)

func AskVision(task string, imagePath string) (string, error) {
	data, err := os.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("read image: %w", err)
	}

	defer func() {
		if err := os.Remove(imagePath); err != nil {
			fmt.Println("remove screenshot:", err)
		}
	}()

	imageBased64 := base64.StdEncoding.EncodeToString(data)

	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"type": map[string]any{
				"type": "string",
				"enum": []string{"move", "click", "type"},
			},
			"x": map[string]any{
				"type": "integer",
			},
			"y": map[string]any{
				"type": "integer",
			},
			"text": map[string]any{
				"type": "string",
			},
		},
		"required":             []string{"type", "x", "y", "text"},
		"additionalProperties": false,
	}

	visionPrompt := `
	Ты управляешь компьютером по скриншоту.
	Возвращай ТОЛЬКО JSON:

	{
  "type": "move | click | type",
  "x": <координата X>,
  "y": <координата Y>,
  "text": ""
	}

	Размер исходного изображения: 2560x1440.

	Координаты должны быть координатами исходного изображения.
	Не используй фиксированные или примерные координаты.
	Сначала найди объект на изображении, затем определи его положение.

	Правила:
	- move — переместить мышь
	- click — кликнуть
	- type — напечатать текст
	- для move/click указывай x и y
	- для type указывай text
	`

	content := visionPrompt + "\nЗадача: " + task

	req := RequestLLMToolCalling{
		Model: "qwen3-vl:4b",
		Messages: []Message{
			{
				Role:    "user",
				Content: content,
				Images:  []string{imageBased64},
			},
		},
		Stream:  false,
		Options: map[string]any{"num_ctx": 8192},
		Format:  schema,
	}

	fmt.Println("TASK:", task)
	fmt.Println("PROMPT:", content)

	resp, err := wrapper(req)
	if err != nil {
		return "", fmt.Errorf("vision request: %w", err)
	}

	fmt.Printf("CONTENT: %q\n", resp.Message.Content)
	fmt.Printf("MESSAGE: %+v\n", resp.Message)

	return resp.Message.Content, nil
}
