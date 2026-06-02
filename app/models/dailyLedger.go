package models

import "time"

type DailyLedger struct {
	TxnNo string `gorm:"column:txnId;primaryKey"`
	RefId string `gorm:"column:refId"`
	AccNo string `gorm:"column:accNo"`
	Amount float64 `gorm:"column:amount"`
	Asset string `gorm:"column:asset"`
	ValueDate time.Time `gorm:"column:valueDate"`
}