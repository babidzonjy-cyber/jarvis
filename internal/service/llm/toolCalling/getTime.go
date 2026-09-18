package llm

import (
	"time"
)

var getTimeTool = Tool{
	Type: "function",
	Function: Function{
		Name:        "get_time",
		Description: "Возвращает текущее время",
		Parameters: Parameters{
			Type:       "object",
			Properties: map[string]Property{},
		},
	},
}

func getTime() string {
	return time.Now().Format("15:04:05")
}
