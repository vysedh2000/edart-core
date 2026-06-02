package dtos

type FundTxnRequest struct {
	Asset string `json:"asset"`
	DebitAcc string `json:"debitAcc"`
	CreditAcc string `json:"creditAcc"`
	TxnCode string `json:"txnCode"`
	Amount string `json:"amount"`
	Narative string `json:"narative"`
	UserId string `json:"userId"`
}