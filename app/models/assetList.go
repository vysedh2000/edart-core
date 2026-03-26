package models

import "time"

type AssetListing struct {
	Symbol string `gorm:"column:symbol"`
	Name string `gorm:"column:name"`
	AddedDate time.Time `gorm:"column:addedDate"`
}