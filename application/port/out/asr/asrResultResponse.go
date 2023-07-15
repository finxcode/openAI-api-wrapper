package asr

type Response struct {
	RequestID string `json:"RequestId"`
	Data      Data   `json:"Data"`
}

type Data struct {
	AudioDuration string      `json:"AudioDuration"`
	ErrorMsg      string      `json:"ErrorMsg"`
	Result        string      `json:"Result"`
	ResultDetail  interface{} `json:"ResultDetail"`
	Status        int         `json:"Status"`
	StatusStr     string      `json:"StatusStr"`
	TaskId        int         `json:"TaskId"`
}
