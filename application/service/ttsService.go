package service

import (
	"chatGPT-api-wrapper/application/port/in/request"
	"chatGPT-api-wrapper/application/port/in/response"
	"chatGPT-api-wrapper/application/port/out/tts"
	"errors"
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
	if ttsResp == nil {
		return response.TtsResponse{}, errors.New("timeout")
	}
	return *ttsResp, err
}
