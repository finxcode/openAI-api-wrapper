package tts

import (
	"chatGPT-api-wrapper/adapter/in/mapper"
	"chatGPT-api-wrapper/application/port/in/request"
	"chatGPT-api-wrapper/application/port/in/response"
	errors2 "errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts/v20190823"
)

type Adapter struct {
	config *Config
}

func NewAdapter() *Adapter {
	return &Adapter{
		config: NewConfig(),
	}
}

func (a *Adapter) GetTtsResult(command *request.TtsCommand) (*response.TtsResponse, error) {
	credential := common.NewCredential(
		a.config.SecretId,
		a.config.SecretKey,
	)

	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tts.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := tts.NewClient(credential, a.config.Region, cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := tts.NewTextToVoiceRequest()

	request.Text = common.StringPtr(command.Text)
	request.SessionId = common.StringPtr(GenerateTtsSessionId())
	request.Codec = common.StringPtr(a.config.Codec)
	request.VoiceType = common.Int64Ptr(a.config.VoiceType)

	// 返回的resp是一个TextToVoiceResponse的实例，与请求对象对应
	voice, err := client.TextToVoice(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if voice == nil {
		return nil, errors2.New("voice timeout")
	}

	resp, err := mapper.TencentTtsResponseMapper(*voice)

	return resp, err
}
