package routes

import (
	"chatGPT-api-wrapper/adapter/in"
	"github.com/gofiber/fiber/v2"
)

func SetApiV1Routes(router fiber.Router, completionAdapter *in.CompletionController) {
	router.Get("/getCompletion", completionAdapter.GetCompletion())
}
