package services

import (
	"edart-core/app/dtos"
	"edart-core/app/models"
	"edart-core/app/repositories"
	"time"
)

type AccountService struct{
	repo *repositories.AccRepository
}

func NewAccountService() *AccountService{
	return &AccountService{repo: repositories.NewAccountRepo()}
}

func (a*AccountService) AccSummaryService(accNo string) (*models.AccountBalance, error) {
	return a.repo.Inquiry(accNo)
}

func (a*AccountService) CreateAccountService(request *dtos.CreateAccountRequest) (*models.AccountBalance, error) {
	var accountInfo models.AccountBalance
	accountInfo.AccNo=request.Uid+"-"+request.Asset
	accountInfo.Asset=request.Asset
	accountInfo.Uid=request.Uid
	accountInfo.WorkingBal=0
	accountInfo.CreateAt=time.Now()

	if err := a.repo.Create(accountInfo); err != nil {
		return nil, err
	}

	return &accountInfo, nil
}