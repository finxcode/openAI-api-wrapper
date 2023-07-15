package main

import (
	"chatGPT-api-wrapper/adapter/in"
	"chatGPT-api-wrapper/adapter/out"
	"chatGPT-api-wrapper/adapter/out/asr"
	"chatGPT-api-wrapper/application/service"
	"chatGPT-api-wrapper/cmd/fiber"
)

func main() {
	//os.Setenv("SECRETID", "AKIDiFYdegbNyNabTAfzA9B0g9RYUEpGZjpl")
	//os.Setenv("SECRETKEY", "zmS1QhpnbLuJYcYJuKXYMyMcDfuwlNn1")
	//os.Setenv("LOGFILE", "D:\\log\\hp-ai-service.log")

	chatGPTCompletionAdapter := out.NewChatGPTCompletionAdapter()
	chatGPTCompletionService := service.NewChatGPTService(chatGPTCompletionAdapter)
	completionController := in.NewCompletionController(chatGPTCompletionService)

	asrAdapter := asr.NewTencentAdapter()
	asrService := service.NewAsrService(asrAdapter)
	asrController := in.NewAsrController(asrService)

	fiber.StartSrv(completionController, asrController)
}
