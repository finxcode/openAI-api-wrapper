package mapper

import (
	"bytes"
	"chatGPT-api-wrapper/adapter/in/utils/mp3"
	"chatGPT-api-wrapper/application/port/in/response"
	"encoding/base64"
	"errors"
	tts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts/v20190823"
	"io"
)

func TencentTtsResponseMapper(voiceResponse tts.TextToVoiceResponse) (*response.TtsResponse, error) {
	ttsResp := response.TtsResponse{}
	if voiceResponse.Response.Audio == nil {
		return nil, errors.New("timeout")
	}
	ttsResp.Audio = *voiceResponse.Response.Audio
	output, _ := base64.StdEncoding.DecodeString(ttsResp.Audio)
	duration, err := getTtsResponseAudioDuration(output)
	if err != nil {
		return nil, err
	}
	ttsResp.AudioDuration = duration
	return &ttsResp, nil
}

func getTtsResponseAudioDuration(output []byte) (int64, error) {
	d := mp3.NewDecoder(bytes.NewReader(output))
	var t int64

	var f mp3.Frame
	skipped := 0

	for {

		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		t = t + f.Duration().Milliseconds()
	}
	return t, nil
}
