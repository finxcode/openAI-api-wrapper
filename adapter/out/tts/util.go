package tts

import "github.com/google/uuid"

func GenerateTtsSessionId() string {
	return uuid.New().String()
}
