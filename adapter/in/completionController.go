package in

import (
	"chatGPT-api-wrapper/adapter/in/utils"
	"chatGPT-api-wrapper/application/port/in"
	"chatGPT-api-wrapper/application/port/in/common"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
)

type CompletionController struct {
	getChatGPTCompletionUseCase in.GetChatGPTCompletionUseCase
}

func NewCompletionController(getChatGPTCompletionUseCase in.GetChatGPTCompletionUseCase) *CompletionController {
	return &CompletionController{
		getChatGPTCompletionUseCase: getChatGPTCompletionUseCase,
	}
}

func (ctl *CompletionController) GetCompletion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mr *utils.MalformedRequest
		command := new(in.CompletionCommand)
		if err := utils.DecodeJSONBody(c, command); err != nil {
			if errors.As(err, &mr) {
				return c.Status(mr.Status).JSON(fiber.Map{
					"message": mr.Msg,
				})
			} else {
				log.Print(err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}
		}
		respBody := ctl.getChatGPTCompletionUseCase.GetChatGPTCompletion(*command)
		if respBody == nil {
			resp := common.Response{
				ErrCode: fiber.StatusServiceUnavailable,
				Message: "service unavailable, please try later",
				Data:    nil,
			}
			return c.JSON(resp)
		} else {
			resp := common.Response{
				ErrCode: 0,
				Message: "ok",
				Data:    respBody,
			}
			return c.JSON(resp)
		}
	}
}
