package entity

type Response struct {
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}