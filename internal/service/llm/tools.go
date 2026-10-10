package llm

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

var openAppTool = Tool{
	Type: "function",
	Function: Function{
		Name:        "open_app",
		Description: "Открывает приложение",
		Parameters: Parameters{
			Type: "object",
			Properties: map[string]Property{
				"app_name": {
					Type:        "string",
					Description: "Название приложения, которое нужно открыть",
				},
			},
			Required: []string{"app_name"},
		},
	},
}

var getWeatherTool = Tool{
	Type: "function",
	Function: Function{
		Name:        "get_weather",
		Description: "показывает погоду",
		Parameters: Parameters{
			Type: "object",
			Properties: map[string]Property{
				"city_name": {
					Type:        "string",
					Description: "Название города, погоду которого нужно узнать, его нужно писать на английском языке",
				},
				"lang_name": {
					Type:        "string",
					Description: "Название языка для гео, например, ru или en",
				},
			},
			Required: []string{"city_name"},
		},
	},
}

var searchWebTool = Tool{
	Type: "function",
	Function: Function{
		Name:        "search_web",
		Description: "Открывает браузер с результатами поиска по заданному запросу",
		Parameters: Parameters{
			Type: "object",
			Properties: map[string]Property{
				"query": {
					Type:        "string",
					Description: "Текст поискового запроса, например: Go channels tutorial",
				},
			},
			Required: []string{"query"},
		},
	},
}

var searchYoutubeTool = Tool{
	Type: "function",
	Function: Function{
		Name:        "search_youtube",
		Description: "Открывает в браузере ссылку на ютуб с результатами поиска по заданному запросу",
		Parameters: Parameters{
			Type: "object",
			Properties: map[string]Property{
				"query": {
					Type:        "string",
					Description: "Текст поискового запроса, например: Go channels tutorial",
				},
			},
			Required: []string{"query"},
		},
	},
}

var openProject = Tool{
	Type: "function",
	Function: Function{
		Name:        "open_project",
		Description: "открывает через терминал нужную мне директорию в IDE",
		Parameters: Parameters{
			Type: "object",
			Properties: map[string]Property{
				"project_name": {
					Type:        "string",
					Description: "Название диреткории, в которую я хочу попасть",
				},
				"ide": {
					Type:        "string",
					Description: "IDE в которой я хочу открыть проект",
					Enum:        []string{"zed", "vscode", "nvim"},
				},
			},
			Required: []string{"project_name", "ide"},
		},
	},
}
