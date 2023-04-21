package in

type GetChatGPTCompletionUseCase interface {
	GetChatGPTCompletion(command CompletionCommand) *CompletionResponse
}
