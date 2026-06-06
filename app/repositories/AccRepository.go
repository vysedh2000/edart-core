package repositories

import (
	"edart-core/app/dtos"
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

func (r*AccRepository) AllAccInquiry(userId string) ([]dtos.AccInq, error) {
	var accSum []dtos.AccInq

	query := `SELECT "accNo", "workingBal", asset, category
			FROM public."AccountBalance" where uid = ?;`

	err := r.DB.Raw(query, userId).Scan(&accSum).Error
	if err != nil {
		return nil, err
	}

	return accSum, nil	
}

func (r*AccRepository) CobGetAccList() ([]dtos.CobAccDto, error) {
	var accList []dtos.CobAccDto

	query := `SELECT "accNo", "workingBal", "closingDate"
			FROM public."AccountBalance";`
			
	err := r.DB.Raw(query).Scan(&accList).Error
	if err != nil {
		return nil, err
	}

	return accList, nil
}

func (r *AccRepository) CobUpdateBal(accNo string, amt float64) (string, error) {
	query := `SELECT cobbalance_test(?, ?);`
	var status string

	// Use Raw and Scan to pull the result into the status variable
	err := r.DB.Raw(query, accNo, amt).Scan(&status).Error
	if err != nil {
		return "", err
	}

	return status, nil
}
