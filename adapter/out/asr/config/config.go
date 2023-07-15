package config

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"os"
)

type Config struct {
	SecretId      string
	SecretKey     string
	Region        string
	ModelType     string
	ChannelNum    uint64
	SpeakerNum    int64
	ResTextFormat uint64
	SourceType    uint64
	SubmitRetry   uint8
	ResultRetry   uint8
	RetryInterval uint16
}

func NewConfig() *Config {
	return &Config{
		SecretId:      os.Getenv("SECRETID"),
		SecretKey:     os.Getenv("SECRETKEY"),
		Region:        regions.Guangzhou,
		ModelType:     "16k_zh-PY",
		ChannelNum:    1,
		SpeakerNum:    1,
		ResTextFormat: 0,
		SourceType:    1,
		SubmitRetry:   3,
		ResultRetry:   5,
		RetryInterval: 3000,
	}
}
