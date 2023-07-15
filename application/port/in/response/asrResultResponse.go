package response

type AsrResponse struct {
	Result        string  `json:"result"`
	AudioDuration float64 `json:"audio_duration"`
}
