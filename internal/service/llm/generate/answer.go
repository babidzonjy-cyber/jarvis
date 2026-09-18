package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LLMRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
	System string `json:"system"`
}

type LLMResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`
}

func Ask(prompt, systemPrompt string) (string, error) {
	defModel := "qwen2.5:7b"

	if systemPrompt == "" {
		systemPrompt = "Голосовой помощник, разговор на русском языке."
	}

	req := LLMRequest{
		Prompt: prompt,
		Model:  defModel,
		Stream: false,
		System: systemPrompt,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("cannot marshal data: %w", err)
	}

	body := bytes.NewReader(data)

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", body)

	if err != nil {
		return "", fmt.Errorf("http post: %w", err)
	}

	defer resp.Body.Close()

	var response LLMResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("cannot decode response: %w", err)
	}

	return response.Response, nil
}
