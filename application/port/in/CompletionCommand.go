package in

type CompletionCommand struct {
	Model       string  `json:"model"`
	Message     Message `json:"message"`
	Temperature float32 `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
