package main

import (
	"chatGPT-api-wrapper/adapter/in"
	"chatGPT-api-wrapper/adapter/out"
	"chatGPT-api-wrapper/application/service"
	"chatGPT-api-wrapper/cmd/fiber"
)

func main() {
	chatGPTCompletionAdapter := out.NewChatGPTCompletionAdapter()
	chatGPTCompletionService := service.NewChatGPTService(chatGPTCompletionAdapter)
	completionController := in.NewCompletionController(chatGPTCompletionService)
	fiber.StartSrv(completionController)
}
