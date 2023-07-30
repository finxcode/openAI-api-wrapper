package in

import (
	adapterCommon "chatGPT-api-wrapper/adapter/common"
	"chatGPT-api-wrapper/adapter/in/mapper"
	"chatGPT-api-wrapper/adapter/in/utils"
	"chatGPT-api-wrapper/application/port/in"
	"chatGPT-api-wrapper/application/port/in/common"
	"chatGPT-api-wrapper/application/port/in/request"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
)

type AsrGptTtsController struct {
	asrUseCase        in.ASRUseCase
	completionUseCase in.GetChatGPTCompletionUseCase
	ttsUseCase        in.TtsUseCase
}

func NewAsrGptTtsController(asrUseCase in.ASRUseCase,
	getChatGPTCompletionUseCase in.GetChatGPTCompletionUseCase,
	ttsUseCase in.TtsUseCase) *AsrGptTtsController {
	return &AsrGptTtsController{
		asrUseCase:        asrUseCase,
		completionUseCase: getChatGPTCompletionUseCase,
		ttsUseCase:        ttsUseCase,
	}
}

func (a *AsrGptTtsController) GetAsrGptTts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mr *utils.MalformedRequest
		var er *utils.Base64EncodeError
		var timeoutError *adapterCommon.AdapterError
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

		if err := utils.DecodeBase64Input(command.AudioBase64); err != nil {
			if errors.As(err, &er) {
				resp := common.Response{
					ErrCode: er.Status,
					Message: er.Msg,
					Data:    nil,
				}
				return c.JSON(resp)
			} else {
				resp := common.Response{
					ErrCode: fiber.StatusPreconditionFailed,
					Message: "error on decode base64 string, please check again",
					Data:    nil,
				}
				return c.JSON(resp)
			}
		}

		command.AudioBase64 = utils.GetRequestBase64(command.AudioBase64)

		respBody, err := a.asrUseCase.RecognizeAudio(command)
		if err != nil {
			if errors.As(err, &timeoutError) {
				resp := common.Response{
					ErrCode: fiber.StatusServiceUnavailable,
					Message: timeoutError.Msg,
					Data:    nil,
				}
				return c.JSON(resp)
			} else {
				resp := common.Response{
					ErrCode: fiber.StatusServiceUnavailable,
					Message: "service unavailable, please try later",
					Data:    nil,
				}
				return c.JSON(resp)
			}
		}

		if respBody == nil {
			resp := common.Response{
				ErrCode: fiber.StatusServiceUnavailable,
				Message: "service unavailable, please try later",
				Data:    nil,
			}
			return c.JSON(resp)
		} else if respBody.Result == "" {
			resp := common.Response{
				ErrCode: fiber.StatusBadRequest,
				Message: "voice can not be recognized",
				Data:    nil,
			}
			return c.JSON(resp)
		} else {
			completionCommand := mapper.AsrRespToCompletionCommand(*respBody)
			completionResp := a.completionUseCase.GetChatGPTCompletion(completionCommand)

			if completionResp == nil {
				resp := common.Response{
					ErrCode: fiber.StatusServiceUnavailable,
					Message: "service unavailable, please try later",
					Data:    nil,
				}
				return c.JSON(resp)
			} else {
				ttsCommand := request.TtsCommand{
					Text: completionResp.Choices[0].Message.Content,
				}
				if ttsResp, err := a.ttsUseCase.GetTts(ttsCommand); err != nil {
					resp := common.Response{
						ErrCode: fiber.StatusServiceUnavailable,
						Message: err.Error(),
						Data:    nil,
					}
					return c.JSON(resp)
				} else {
					resp := common.Response{
						ErrCode: 200,
						Message: "ok",
						Data:    ttsResp,
					}
					return c.JSON(resp)
				}
			}
		}
	}
}

func (a *AsrGptTtsController) GetTxtGptTts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mr *utils.MalformedRequest
		command := new(in.CompletionCommand)

		if err := utils.DecodeJSONBody(c, command); err != nil {
			if errors.As(err, &mr) {
				resp := common.Response{
					ErrCode: mr.Status,
					Message: mr.Msg,
					Data:    nil,
				}
				return c.JSON(resp)
			} else {
				log.Print(err.Error())
				resp := common.Response{
					ErrCode: fiber.StatusServiceUnavailable,
					Message: "service unavailable, please try later",
					Data:    nil,
				}
				return c.JSON(resp)
			}
		}

		respBody := a.completionUseCase.GetChatGPTCompletion(*command)
		if respBody == nil {
			resp := common.Response{
				ErrCode: fiber.StatusServiceUnavailable,
				Message: "service unavailable, please try later",
				Data:    nil,
			}
			return c.JSON(resp)
		} else {
			ttsCommand := request.TtsCommand{
				Text: respBody.Choices[0].Message.Content,
			}
			if ttsResp, err := a.ttsUseCase.GetTts(ttsCommand); err != nil {
				resp := common.Response{
					ErrCode: fiber.StatusServiceUnavailable,
					Message: err.Error(),
					Data:    nil,
				}
				return c.JSON(resp)
			} else {
				resp := common.Response{
					ErrCode: 200,
					Message: "ok",
					Data:    ttsResp,
				}
				return c.JSON(resp)
			}
		}
	}
}