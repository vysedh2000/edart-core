package repositories

import (
	"edart-core/app/models"
	"edart-core/database"

	"gorm.io/gorm"
)

type DailyLedgerRepository struct {
	DB * gorm.DB
}

type AssetBalance struct {
	Asset       string  `gorm:"column:asset"`
	TotalAmount float64 `gorm:"column:total_amount"`
}

func NewDailyBalRepository() *DailyLedgerRepository{
	db := database.DB
	return &DailyLedgerRepository{DB : db}
}

func (d * DailyLedgerRepository) CreateDBal(ledgerBal models.DailyLedger) error{
	return d.DB.Create(&ledgerBal).Error
}

func (d * DailyLedgerRepository) CreateDailyBalBatch(ledgerBal []models.DailyLedger) error{
	return d.DB.Create(&ledgerBal).Error
}

func (r * DailyLedgerRepository) SumBalByAcc(accNo string) (float64, error) {
	var totalAmount float64

	query := `
		SELECT COALESCE(SUM(amount), 0) as total_amount 
		FROM public.daily_ledgers 
		WHERE "accNo" = ?
	`
	
	err := r.DB.Raw(query, accNo).Scan(&totalAmount).Error
	if err != nil {
		return 0, err
	}

	return totalAmount, nil
}