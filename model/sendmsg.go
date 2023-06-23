package model

type SendMsg struct {
	ReceiveId string `json:"receive_id"`
	MsgType   string `json:"msg_type"`
	Content   string `json:"content"`
}
