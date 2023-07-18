package model

// Message chatGptApi 消息封装对象
// role: system/user/assistant/function
// content: 消息内容
// name: role=function时候必填,传入调用function的name
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
}
