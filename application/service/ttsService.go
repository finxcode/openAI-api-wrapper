package service

import (
	"chatGPT-api-wrapper/application/port/in/request"
	"chatGPT-api-wrapper/application/port/in/response"
	"chatGPT-api-wrapper/application/port/out/tts"
)

type TtsService struct {
	ttsPort tts.Port
}

func NewTtsService(ttsPort tts.Port) *TtsService {
	return &TtsService{
		ttsPort: ttsPort,
	}
}

func (t *TtsService) GetTts(command request.TtsCommand) (response.TtsResponse, error) {
	ttsResp, err := t.ttsPort.GetTtsResult(&command)
	return *ttsResp, err
}
