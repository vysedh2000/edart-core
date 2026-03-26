package repositories

import (
	"edart-core/app/models"
	"edart-core/database"

	"gorm.io/gorm"
)

type AccRepository struct {
	DB * gorm.DB
}

func NewAccountRepo() *AccRepository{
	db := database.DB
	return &AccRepository{DB: db}
}

func (r*AccRepository) Create(acc models.AccountBalance) error {
	return r.DB.Create(&acc).Error
}

func (r*AccRepository) Inquiry(accNo string) (*models.AccountBalance, error) {
	var accSummary models.AccountBalance

	err := r.DB.Where(`"accNo"`, "= ?", accNo).First(&accSummary).Error
	if err != nil {
		return nil, err
	}

	return &accSummary, nil
}