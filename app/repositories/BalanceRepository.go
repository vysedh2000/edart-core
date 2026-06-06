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

func (r * BalanceRepository) CobGetBalByAcc(accNo string) (float64, error) {
	var totalAmount float64

	query := `
		SELECT COALESCE(SUM(amount), 0) as total_amount 
		FROM public."BalanceLedger"
		WHERE "accNo" = ?
	`
	err := r.DB.Raw(query, accNo).Scan(&totalAmount).Error
	if err != nil {
		return 0, err
	}
	return totalAmount, nil
}