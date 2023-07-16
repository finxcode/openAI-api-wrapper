package mapper

import (
	"chatGPT-api-wrapper/application/port/in"
	"chatGPT-api-wrapper/application/port/in/response"
)

func AsrRespToCompletionCommand(response response.AsrResponse) in.CompletionCommand {
	var messages []in.Message
	message := in.Message{
		Role:    "user",
		Content: response.Result,
	}

	messages = append(messages, message)

	return in.CompletionCommand{
		Model:       "gpt-3.5-turbo",
		Temperature: 0.7,
		Messages:    messages,
	}
}
