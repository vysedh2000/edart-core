package repositories

import (
	"edart-core/database"

	"gorm.io/gorm"
)

type BalanceRepository struct{
	DB * gorm.DB
}

func NewBalanceRepo() *BalanceRepository{
	db := database.DB
	return &BalanceRepository{DB: db}
}

