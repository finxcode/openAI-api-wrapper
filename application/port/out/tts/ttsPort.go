package tts

import (
	"chatGPT-api-wrapper/application/port/in/request"
	"chatGPT-api-wrapper/application/port/in/response"
)

type Port interface {
	GetTtsResult(command *request.TtsCommand) (*response.TtsResponse, error)
}
