package models

// Implementation of huggin face services using curl command
type Generate struct {
	provider    string `json:"provider"`
	model       string `json:"model"`
	contentType string `json:"content_type"`
	url		    string `json:"url"`
	token	   string `json:"token"`
}