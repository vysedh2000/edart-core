package dtos

type GetAccSumRequest struct {
	AccNo string `json:"accNo"`
}

type CreateAccountRequest struct {
	Uid string `json:"uid"`
	Asset string `json:"asset"`
}