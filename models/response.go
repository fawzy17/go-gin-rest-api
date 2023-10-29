package models

type Response struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Status        string      `json:"status"`
	Data          interface{} `json:"data"`
}
