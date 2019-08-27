package models

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
	Data   string `json:"data"`
}
