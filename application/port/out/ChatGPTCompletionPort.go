package out

import "chatGPT-api-wrapper/application/port/in"

type GetChatGPTCompletionPort interface {
	GetChatGPTCompletionOutgoing(command in.CompletionCommand) *in.CompletionResponse
}
