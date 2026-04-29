package models

import "time"

type UserInfo struct {
	Uid string `gorm:"column:uid;primaryKey"`
	Fullname string `gorm:"column:fullname"`
	Nationality string `gorm:"column:nationality"`
	Country string `gorm:"column:country"`
	IdNum string `gorm:"column:idNum"`
	IdType string `gorm:"column:idType"`
	Dob time.Time `gorm:"column:dob"`
	CreateAt time.Time `gorm:"column:createdAt"`
}

func(UserInfo) TableName() string {
	return "UserInfo"
}