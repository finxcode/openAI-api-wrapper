package service

import (
	"chatGPT-api-wrapper/application/port/in"
	"chatGPT-api-wrapper/application/port/out"
	"log"
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
	if resp, err := c.getChatGPTCompletionPort.GetChatGPTCompletionOutgoing(command); err != nil {
		log.Println(err.Error())
		log.Println("here")
		return nil
	} else {
		return resp
	}

}
