package tts

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"os"
)

type Config struct {
	SecretId  string
	SecretKey string
	Region    string
	Codec     string
	VoiceType int64
}

func NewConfig() *Config {
	return &Config{
		SecretId:  os.Getenv("SECRETID"),
		SecretKey: os.Getenv("SECRETKEY"),
		Region:    regions.Guangzhou,
		Codec:     "mp3",
		VoiceType: 101023,
	}
}
