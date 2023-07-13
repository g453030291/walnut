package model

type Chat struct {
	Model       string      `json:"model"`
	Messages    []Message   `json:"messages"`
	Temperature float64     `json:"temperature"`
	Functions   []Functions `json:"functions"`
}

type Functions struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

type Parameters struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required"`
}
