package service

import (
	"chatGPT-api-wrapper/application/port/in"
	"chatGPT-api-wrapper/application/port/out"
)

type ChatGPTService struct {
	getChatGPTCompletionPort out.GetChatGPTCompletionPort
}

func NewChatGPTService(getChatGPTCompletionPort out.GetChatGPTCompletionPort) *ChatGPTService {
	return &ChatGPTService{
		getChatGPTCompletionPort: getChatGPTCompletionPort,
	}
}

func (c *ChatGPTService) GetChatGPTCompletion(command in.CompletionCommand) *in.CompletionResponse {
	return c.getChatGPTCompletionPort.GetChatGPTCompletionOutgoing(command)
}
