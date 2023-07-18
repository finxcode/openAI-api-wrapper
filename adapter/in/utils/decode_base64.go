package utils

import (
	"encoding/base64"
	"net/http"
	"strings"
)

var supportedTypes = [11]string{
	"wav", "mp3", "m4a", "flv", "mp4", "wma", "3gp", "amr", "aac", "ogg-opus", "flac",
}

type Base64EncodeError struct {
	Status int
	Msg    string
}

func (b *Base64EncodeError) Error() string {
	return b.Msg
}

func DecodeBase64Input(encoded string) error {
	output, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		return &Base64EncodeError{
			Status: http.StatusBadRequest,
			Msg:    "wrong base64 encoding format",
		}
	}

	if len(output) < 512 {
		return &Base64EncodeError{
			Status: http.StatusBadRequest,
			Msg:    "data too short",
		}
	}

	contentType := http.DetectContentType(output[:512])
	inputType := strings.Split(contentType, "/")

	if len(inputType) < 2 {
		return &Base64EncodeError{
			Status: http.StatusBadRequest,
			Msg:    "cannot identify input type",
		}
	}

	if !isSupported(inputType[1], supportedTypes) {
		return &Base64EncodeError{
			Status: http.StatusBadRequest,
			Msg:    "input data type not supported",
		}
	}

	return nil
}

func DecodeBase64TtsOutput(encoded string) error {
	_, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		return &Base64EncodeError{
			Status: http.StatusServiceUnavailable,
			Msg:    "tts response audio base64 cannot be decoded",
		}
	}

	return nil
}

func isSupported(t string, supportedTypes [11]string) bool {
	for i := 0; i < len(supportedTypes); i++ {
		if t == supportedTypes[i] {
			return true
		}
	}
	return false
}
