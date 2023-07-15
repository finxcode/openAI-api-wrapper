package service

import (
	"chatGPT-api-wrapper/application/port/in/request"
	"chatGPT-api-wrapper/application/port/in/response"
	"chatGPT-api-wrapper/application/port/out/asr"
)

type AsrService struct {
	asrPort asr.Port
}

func NewAsrService(asrPort asr.Port) *AsrService {
	return &AsrService{
		asrPort: asrPort,
	}
}

func (a *AsrService) RecognizeAudio(command *request.ASRCommand) (*response.AsrResponse, error) {
	asrResp, err := a.asrPort.GetAsrResult(command)
	if err != nil {
		return nil, err
	}
	return asrResp, nil
}
