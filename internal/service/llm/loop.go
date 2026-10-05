package llm

import (
	"fmt"
	"jarvis/internal/service/tools"
)

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
		Tools: []Tool{
			getTimeTool,
			openAppTool,
			getWeatherTool,
			searchWebTool,
			searchYoutubeTool,
		},
	}

	resp, err := wrapper(req)
	if err != nil {
		return "", err
	}

	if len(resp.Message.ToolCalls) == 0 {
		fmt.Println("Ответ без tool calling")
		return resp.Message.Content, nil
	}

	messages := []Message{
		{
			Role:    "user",
			Content: prompt,
		},
		{
			Role:      "assistant",
			Content:   resp.Message.Content,
			ToolCalls: resp.Message.ToolCalls,
		},
	}

	for _, toolCall := range resp.Message.ToolCalls {
		var res string

		switch toolCall.Function.Name {
		case "get_time":
			res = tools.GetTime()

		case "open_app":
			appName, ok := toolCall.Function.Arguments["app_name"].(string)
			if !ok {
				return "", fmt.Errorf("invalid app_name")
			}

			if err := tools.OpenApp(appName); err != nil {
				return "", fmt.Errorf("open app: %w", err)
			}

			res = fmt.Sprintf("Приложение %s открыто", appName)

		case "get_weather":
			cityName, ok := toolCall.Function.Arguments["city_name"].(string)
			if !ok {
				return "", fmt.Errorf("invalid city_name")
			}

			langName, ok := toolCall.Function.Arguments["lang_name"].(string)
			if !ok {
				return "", fmt.Errorf("invalid lang_name")
			}

			res, err = tools.GetWeather(cityName, langName)
			if err != nil {
				return "", fmt.Errorf("get weather: %w", err)
			}

		case "search_web":
			query, ok := toolCall.Function.Arguments["query"].(string)
			if !ok {
				return "", fmt.Errorf("invalid query")
			}

			res, err = tools.SearchWeb(query)
			if err != nil {
				return "", fmt.Errorf("search web: %w", err)
			}

		case "search_youtube":
			query, ok := toolCall.Function.Arguments["query"].(string)
			if !ok {
				return "", fmt.Errorf("search youtube: %w", err)
			}

			res, err = tools.SearchYoutube(query)
			if err != nil {
				return "", fmt.Errorf("search youtube: %w", err)
			}

		default:
			return "", fmt.Errorf("unknown tool: %s", toolCall.Function.Name)
		}

		messages = append(messages, Message{
			Role:    "tool",
			Content: res,
		})
	}

	newReq := RequestLLMToolCalling{
		Model:    defModel,
		Messages: messages,
		Stream:   false,
		Tools:    []Tool{getTimeTool},
	}

	newResp, err := wrapper(newReq)
	if err != nil {
		return "", err
	}

	fmt.Println("Ответ с tool calling")
	return newResp.Message.Content, nil
}
