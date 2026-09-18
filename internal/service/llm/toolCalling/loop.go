package llm

import "fmt"

func AskWithTools(prompt string) (string, error) {
	defModel := "qwen2.5:7b"
	req := RequestLLMToolCalling{
		Model: defModel,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Stream: false,
		Tools:  []Tool{getTimeTool},
	}

	resp, err := wrapper(req)
	if err != nil {
		return "", err
	}

	if len(resp.Message.ToolCalls) == 0 {
		return resp.Message.Content, nil
	}

	firstFunc := resp.Message.ToolCalls[0]
	if firstFunc.Function.Name != "get_time" {
		return "", fmt.Errorf("проверка на то, что это функция get_time, если уже другие функции добавил, то отладить этот говнокод")
	}

	tool := getTime()

	newReq := RequestLLMToolCalling{
		Model: defModel,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
			{
				Role:      "assistant",
				Content:   resp.Message.Content,
				ToolCalls: resp.Message.ToolCalls,
			},
			{
				Role:    "tool",
				Content: tool,
			},
		},
		Stream: false,
		Tools:  []Tool{getTimeTool},
	}

	newResp, err := wrapper(newReq)
	if err != nil {
		return "", err
	}

	return newResp.Message.Content, nil
}
