package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func wrapper(req RequestLLMToolCalling) (ResponseLLMToolCalling, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return ResponseLLMToolCalling{}, fmt.Errorf("cannot marshal data: %w", err)
	}
	fmt.Println(req.Model, req.Messages[0].Content)

	body := bytes.NewReader(data)

	resp, err := http.Post("http://localhost:11434/api/chat", "application/json", body)
	if err != nil {
		return ResponseLLMToolCalling{}, fmt.Errorf("http post: %w", err)
	}

	defer resp.Body.Close()

	var response ResponseLLMToolCalling
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return ResponseLLMToolCalling{}, fmt.Errorf("cannot decode response: %w", err)
	}

	return response, nil
}
