package utils

import (
	"encoding/base64"
	"github.com/vincent-petithory/dataurl"
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

	dataUrl, err := dataurl.DecodeString(encoded)
	if err != nil {
		return &Base64EncodeError{
			Status: http.StatusBadRequest,
			Msg:    "wrong base64 encode",
		}
	}
	subType := dataUrl.MediaType.Subtype

	if !isSupported(subType, supportedTypes) {
		return &Base64EncodeError{
			Status: http.StatusBadRequest,
			Msg:    "input data type not supported",
		}
	}

	if len(strings.Split(encoded, ",")) < 2 {
		return &Base64EncodeError{
			Status: http.StatusBadRequest,
			Msg:    "input content should be separate by comma",
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

func GetRequestBase64(encoded string) string {
	return strings.Split(encoded, ",")[1]
}

func isSupported(t string, supportedTypes [11]string) bool {
	for i := 0; i < len(supportedTypes); i++ {
		if t == supportedTypes[i] {
			return true
		}
	}
	return false
}
