package llm

type RequestLLMToolCalling struct {
	Model    string         `json:"model"`
	Messages []Message      `json:"messages"`
	Stream   bool           `json:"stream"`
	Tools    []Tool         `json:"tools,omitempty"`
	Options  map[string]any `json:"options,omitempty"`
	Format   any            `json:"format,omitempty"`
}

type Message struct {
	Role      string     `json:"role"`
	Content   string     `json:"content"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	Images    []string   `json:"images,omitempty"`
}

type Tool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Function struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

type Parameters struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"` // аргументы для функции
	Required   []string            `json:"required"`
}

type Property struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Enum        []string `json:"enum"`
}
