package in

import (
	"chatGPT-api-wrapper/application/port/in/request"
	"chatGPT-api-wrapper/application/port/in/response"
)

type ASRUseCase interface {
	RecognizeAudio(asrCommand *request.ASRCommand) (*response.AsrResponse, error)
}
