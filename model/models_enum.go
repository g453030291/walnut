package model

type ModelsEnum string

// 4k->8k->16k->32k
const (
	GPT35    ModelsEnum = "gpt-3.5-turbo"
	GPT4     ModelsEnum = "gpt-4"
	GPT3516K ModelsEnum = "gpt-3.5-turbo-16k"
	GPT432K  ModelsEnum = "gpt-4-32k"
)
