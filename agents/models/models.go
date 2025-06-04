package models

// Implementation of huggin face services using curl command
type Generate struct {
	Prompt      string `json:"prompt"`
	Provider    string `json:"provider"`
	Model       string `json:"model"`
	ContentType string `json:"content_type"`
	Url         string `json:"url"`
	Token       string `json:"token"`
}
