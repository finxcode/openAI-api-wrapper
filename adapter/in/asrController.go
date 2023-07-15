package in

import (
	"chatGPT-api-wrapper/adapter/in/utils"
	"chatGPT-api-wrapper/application/port/in"
	"chatGPT-api-wrapper/application/port/in/common"
	"chatGPT-api-wrapper/application/port/in/request"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type AsrController struct {
	asrUseCase in.ASRUseCase
}

func NewAsrController(asrUseCase in.ASRUseCase) *AsrController {
	return &AsrController{
		asrUseCase: asrUseCase,
	}
}

// GetASR todo verify the audio base64 code is valid
func (a *AsrController) GetASR() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mr *utils.MalformedRequest
		command := new(request.ASRCommand)
		if err := utils.DecodeJSONBody(c, command); err != nil {
			if errors.As(err, &mr) {
				resp := common.Response{
					ErrCode: mr.Status,
					Message: mr.Msg,
					Data:    nil,
				}
				return c.JSON(resp)
			} else {
				resp := common.Response{
					ErrCode: fiber.StatusPreconditionFailed,
					Message: "error on decode request body, please check again",
					Data:    nil,
				}
				return c.JSON(resp)
			}
		}
		respBody, err := a.asrUseCase.RecognizeAudio(command)
		if err != nil || respBody == nil {
			resp := common.Response{
				ErrCode: fiber.StatusServiceUnavailable,
				Message: "service unavailable, please try later",
				Data:    nil,
			}
			return c.JSON(resp)
		} else {
			resp := common.Response{
				ErrCode: 200,
				Message: "ok",
				Data:    respBody,
			}
			return c.JSON(resp)
		}
	}
}
