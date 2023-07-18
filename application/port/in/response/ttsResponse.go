package response

type TtsResponse struct {
	Audio         string `json:"audio"`
	AudioDuration int64  `json:"audio_duration"`
}
