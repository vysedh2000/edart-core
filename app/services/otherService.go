package services

import (
	"edart-core/app/models"
	"edart-core/app/repositories"
	"errors"

	"gorm.io/gorm"
)

type OtherService struct {
	accRepo * repositories.AccRepository
}

func NewOtherService() * OtherService{
	return &OtherService{accRepo: repositories.NewAccountRepo()}
}

func (o *OtherService) FindAccService(accountId string) (bool, error) {

	if o == nil || o.accRepo == nil || o.accRepo.DB == nil {
		return false, errors.New("database connection or repository is not initialized")
	}

	var accBal models.AccountBalance

	err := o.accRepo.DB.Select(`"accNo"`).First(&accBal, `"accNo" = ?`, accountId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

