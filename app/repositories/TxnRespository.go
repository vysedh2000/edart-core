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