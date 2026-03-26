package models

import "time"

type AccountBalance struct {
	AccNo string `gorm:"column:accNo;primaryKey"`
	WorkingBal float64 `gorm:"column:workingBal"`
	Asset string `gorm:"column:asset"`
	Uid string `gorm:"column:uid"`
	CreateAt time.Time `gorm:"column:createdAt"`
}

func(AccountBalance) TableName() string{
	return "AccountBalance"
}