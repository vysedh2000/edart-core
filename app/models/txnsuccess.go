package models

import (
    "time"
)

type TxnSuccess struct {
	TxnId string `gorm:"column:txnId;primaryKey"`
	RefId string `gorm:"column:refId"`
	AccId string `gorm:"column:accId"`
	TxnType string `gorm:"column:txnType"`
	TxnCode string `gorm:"column:txnCode"`
	Amount string `gorm:"column:amount"`
	Narative string `gorm:"column:narative"`
	TxnDate time.Time `gorm:"column:txnDate"`
	UserId string `gorm:"column:userId"`
}

func(TxnSuccess) TableName() string {
	return "TxnSuccess"
}