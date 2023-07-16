package asr

import (
	adapterCommon "chatGPT-api-wrapper/adapter/common"
	"chatGPT-api-wrapper/adapter/out/asr/config"
	"chatGPT-api-wrapper/application/port/in/request"
	"chatGPT-api-wrapper/application/port/in/response"
	"fmt"
	v20190614 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"strings"
	"time"
)

type TencentAdapter struct {
	config *config.Config
}

func NewTencentAdapter() *TencentAdapter {
	return &TencentAdapter{
		config: config.NewConfig(),
	}
}

func (t *TencentAdapter) GetAsrResult(command *request.ASRCommand) (*response.AsrResponse, error) {
	asrResp := response.AsrResponse{}

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "asr.tencentcloudapi.com"

	credential := common.NewCredential(
		t.config.SecretId,
		t.config.SecretKey,
	)

	client, _ := v20190614.NewClient(credential, t.config.Region, cpf)
	data := command.AudioBase64

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := v20190614.NewCreateRecTaskRequest()

	request.EngineModelType = common.StringPtr(t.config.ModelType)
	request.ChannelNum = common.Uint64Ptr(t.config.ChannelNum)
	request.SpeakerNumber = common.Int64Ptr(t.config.SpeakerNum)
	request.ResTextFormat = common.Uint64Ptr(t.config.ResTextFormat)
	request.SourceType = common.Uint64Ptr(t.config.SourceType)
	request.Data = common.StringPtr(data)

	var recResp *v20190614.CreateRecTaskResponse
	var submitError error
	for i := 0; i < int(t.config.SubmitRetry); i++ {
		recResp, submitError = client.CreateRecTask(request)

		if _, ok := submitError.(*errors.TencentCloudSDKError); ok {
			continue
		}
		if submitError != nil {
			continue
		}
		if *recResp.Response.Data.TaskId != 0 {
			break
		}
	}

	if recResp != nil && *recResp.Response.Data.TaskId != 0 {
		taskId := recResp.Response.Data.TaskId
		req := v20190614.NewDescribeTaskStatusRequest()
		req.TaskId = taskId

		for i := 0; i < int(t.config.ResultRetry); i++ {

			resResp, err := client.DescribeTaskStatus(req)
			if _, ok := err.(*errors.TencentCloudSDKError); ok {
				time.Sleep(time.Duration(int64(t.config.RetryInterval)) * time.Millisecond)
				continue
			}
			if err != nil {
				time.Sleep(time.Duration(int64(t.config.RetryInterval)) * time.Millisecond)
				continue
			}
			if *resResp.Response.Data.Status != 2 {
				time.Sleep(time.Duration(int64(t.config.RetryInterval)) * time.Millisecond)
				continue
			}
			if *resResp.Response.Data.Status == 2 {
				fmt.Println(resResp.ToJsonString())
				asrResp.Result = parseResult(*resResp.Response.Data.Result)
				asrResp.AudioDuration = *resResp.Response.Data.AudioDuration * 1000
				return &asrResp, nil
			}
		}
	} else {
		return nil, submitError
	}

	return nil, &adapterCommon.AdapterError{
		Code: 1001,
		Msg:  "tencent api timeout",
	}
}

func parseResult(res string) string {
	if len(res) == 0 {
		return res
	}

	sepRes := strings.Split(res, "  ")
	if len(sepRes) < 2 {
		return res
	}
	fmt.Println(sepRes)
	return sepRes[1]
}
