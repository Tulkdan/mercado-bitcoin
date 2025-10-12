package dto

type AccountBalance struct {
	BTC int64 `json:"BTC"`
	BRL int64 `json:"BRL"`
}

type AccountOutput struct {
	Id      string         `json:"id"`
	Balance AccountBalance `json:"balance"`
}
