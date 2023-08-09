package model

type Sms struct {
	Tel        string            `json:"tel" binding:"required"`
	TemplateId string            `json:"template_id"`
	Param      map[string]string `json:"param"`
}
