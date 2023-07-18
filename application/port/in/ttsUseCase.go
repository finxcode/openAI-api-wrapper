package in

import (
	"chatGPT-api-wrapper/application/port/in/request"
	"chatGPT-api-wrapper/application/port/in/response"
)

type TtsUseCase interface {
	GetTts(command request.TtsCommand) (response.TtsResponse, error)
}
