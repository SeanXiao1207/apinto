package wenxin

type ClientRequest struct {
	Messages []*Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Id           string `json:"id"`
	Object       string `json:"object"`
	Created      int    `json:"created"`
	Result       string `json:"result"`
	FinishReason string `json:"finish_reason"`
	ErrorCode    int    `json:"error_code"`
	ErrorMsg     string `json:"error_msg"`
	Usage        Usage  `json:"usage"`
}

// Usage provides information about the token counts for the request and response.
// {"prompt_tokens":19,"completion_tokens":21,"total_tokens":40}
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
