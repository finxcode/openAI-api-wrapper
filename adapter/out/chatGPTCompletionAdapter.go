package out

import (
	"bytes"
	"chatGPT-api-wrapper/adapter/common"
	"chatGPT-api-wrapper/application/port/in"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const URL = "https://api.openai.com/v1/chat/completions"

type ChatGPTCompletionAdapter struct {
}

func NewChatGPTCompletionAdapter() *ChatGPTCompletionAdapter {
	return &ChatGPTCompletionAdapter{}
}

func (c *ChatGPTCompletionAdapter) GetChatGPTCompletionOutgoing(command in.CompletionCommand) (*in.CompletionResponse, error) {
	respBody := in.CompletionResponse{}
	marshal, err := json.Marshal(command)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewReader(marshal)

	r, err := http.NewRequest("POST", URL, reqBody)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&respBody)
	if err != nil {
		return nil, err
	}

	return nil, &common.AdapterError{
		Code: 1001,
		Msg:  "tencent api timeout",
	}
}
