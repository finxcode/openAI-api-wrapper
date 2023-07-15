package asr

import (
	"chatGPT-api-wrapper/application/port/in/request"
	"chatGPT-api-wrapper/application/port/in/response"
)

type Port interface {
	GetAsrResult(command *request.ASRCommand) (*response.AsrResponse, error)
}
