package routes

import (
	"chatGPT-api-wrapper/adapter/in"
	"github.com/gofiber/fiber/v2"
)

func SetApiV1Routes(router fiber.Router,
	completionAdapter *in.CompletionController,
	asrController *in.AsrController,
	asrCompletionController *in.AsrCompletionController,
	asrGptTtsController *in.AsrGptTtsController) {
	router.Post("/completion", completionAdapter.GetCompletion())
	router.Post("/asr", asrController.GetASR())
	router.Post("/asrCompletion", asrCompletionController.GetAsrCompletion())
	router.Post("/asrCompletionTts", asrGptTtsController.GetAsrGptTts())
	router.Post("/completionTts", asrGptTtsController.GetTxtGptTts())
}
