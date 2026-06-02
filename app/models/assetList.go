package models

import "time"

type AssetListing struct {
	Symbol string `gorm:"column:symbol"`
	Name string `gorm:"column:name"`
	Category string `gorm:"column:category"`
	Logo string `gorm:"column:logo"`
	AddedDate time.Time `gorm:"column:addedDate"`
}