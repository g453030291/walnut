package model

// Msg 消息接口通用封装入参
type Msg struct {
	Id      string     `json:"id"`
	Content string     `json:"content"`
	Model   ModelsEnum `json:"model"`
}
