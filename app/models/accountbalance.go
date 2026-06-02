package models

import "time"

type AccountBalance struct {
	AccNo string `gorm:"column:accNo;primaryKey"`
	WorkingBal float64 `gorm:"column:workingBal"`
	Asset string `gorm:"column:asset"`
	Category string `gorm:"column:category"`
	Uid string `gorm:"column:uid"`
	ClosingDate time.Time `gorm:"column:closingDate"`
	CreateAt time.Time `gorm:"column:createdAt"`
}

func(AccountBalance) TableName() string{
	return "AccountBalance"
}