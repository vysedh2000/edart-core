package models

import "time"

type BalanceLedger struct {
	TxnNo uint `gorm:"column:txnId;primaryKey;autoIncrement"`
	RefId string `gorm:"column:refId"`
	AccNo string `gorm:"column:accNo"`
	Amount float64 `gorm:"column:amount"`
	Asset string `gorm:"column:asset"`
	ValueDate time.Time `gorm:"column:valueDate"`
}

func(BalanceLedger) TableName() string{
	return "BalanceLedger"
}