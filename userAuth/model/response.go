package model

type Response struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
