package model

// Message chatGptApi 消息封装对象
// role: system/user/assistant/function
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
