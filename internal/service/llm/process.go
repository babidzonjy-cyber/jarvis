package llm

func Process(prompt string) (answer string, err error) {
	return AskWithTools(prompt)
}
