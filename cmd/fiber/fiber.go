package fiber

import (
	"chatGPT-api-wrapper/adapter/in"
	"chatGPT-api-wrapper/cmd/fiber/middleware"
	"chatGPT-api-wrapper/cmd/fiber/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

const (
	PORT    = ":3000"
	API     = "api"
	Version = "v1"
)

func setRouteGroupApiV1(app *fiber.App) fiber.Router {
	prefix := fmt.Sprintf("/%s/%s", API, Version)
	return app.Group(prefix, middleware.ApiKeyAuth(), middleware.FiberLogger())
}

func StartSrv(completionController *in.CompletionController,
	asrController *in.AsrController) {
	app := fiber.New()
	api := setRouteGroupApiV1(app)
	routes.SetApiV1Routes(api, completionController, asrController)
	err := app.Listen(PORT)
	log.Fatalf("server started failed with error:%s", err.Error())
}
