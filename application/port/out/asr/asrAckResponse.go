package asr

type AckResponse struct {
	RequestId string  `json:"RequestId"`
	Data      AckData `json:"Data"`
}

type AckData struct {
	TaskId int `json:"TaskId"`
}
