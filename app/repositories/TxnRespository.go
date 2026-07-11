package repositories

import (
	"edart-core/app/models"
	"edart-core/database"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TxnRepository struct{
	DB * gorm.DB
}

func NewTxnRepo() * TxnRepository{
	db := database.DB
	return &TxnRepository{DB: db}
}

func (r*TxnRepository) CreateTxn(txn models.TxnSuccess) error {
	return r.DB.Create(&txn).Error;
}

func (r*TxnRepository) CreateTxnBatch(txn []models.TxnSuccess) error {
	return r.DB.Create(&txn).Error;
}

func (r *TxnRepository) GetNextTxnId() (string, error) {
	var seq int64

	err := r.DB.Raw("SELECT nextval('txn_id_seq')").Scan(&seq).Error
	if err != nil {
		return "", err
	}

	now := time.Now()
	year := now.Year()
	dayOfYear := now.YearDay()

	txnId := fmt.Sprintf("%04d%03d%06d", year, dayOfYear, seq)

	return txnId, nil
}

func (r *TxnRepository) GetBalByAcc(accNo string) (float64, error) {
	var totalAmt float64

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tmr := today.AddDate(0, 0, 1)
	
	query := `
		SELECT COALESCE(SUM(amount), 0) as total_amount 
		FROM public."TxnSuccess"
		WHERE "accId" = ? AND "txnDate" >= ? AND "txnDate" < ?
	`
	err := r.DB.Raw(query, accNo, today, tmr).Scan(&totalAmt).Error
	if err != nil {
		return 0, err
	}
	return totalAmt, nil
}