package out

import "chatGPT-api-wrapper/application/port/in"

type ChatGPTCompletionAdapter struct {
}

func NewChatGPTCompletionAdapter() *ChatGPTCompletionAdapter {
	return &ChatGPTCompletionAdapter{}
}

func (c *ChatGPTCompletionAdapter) GetChatGPTCompletionOutgoing(command in.CompletionCommand) *in.CompletionResponse {
	return nil
}
