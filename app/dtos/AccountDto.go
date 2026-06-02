package dtos

import "time"

type GetAccSumRequest struct {
	AccNo string `json:"accNo"`
}

type CreateAccountRequest struct {
	Uid string `json:"uid"`
	Asset string `json:"asset"`
}

type AccInq struct{
	AccNo string `gorm:"column:accNo" json:"accNo"`
	CloseBal float64 `gorm:"column:workingBal" json:"workingBal"`
	Asset string `gorm:"column:asset" json:"asset"`
	Category string `gorm:"column:category" json:"category"`
}

type CobAccDto struct{
	AccNo string `gorm:"column:accNo" json:"accNo"`
	WorkingBal float64 `gorm:"column:workingBal" json:"workingBal"`
	ClosingDate time.Time `gorm:"column:closingDate" json:"closingDate"`
}