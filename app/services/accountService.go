package services

import (
	"edart-core/app/dtos"
	"edart-core/app/models"
	"edart-core/app/repositories"
	"time"
)

type AccountService struct{
	repo *repositories.AccRepository
	dailyBalRepo * repositories.DailyLedgerRepository
	assetRepo * repositories.AssetRepository
}

func NewAccountService() *AccountService{
	return &AccountService{repo: repositories.NewAccountRepo(),
		dailyBalRepo: repositories.NewDailyBalRepository(),
		assetRepo: repositories.NewAssetRepo(),
	}
}

func (a*AccountService) AccSummaryService(accNo string) (any, error) {
	accList, err := a.repo.AllAccInquiry(accNo)
	if err != nil {
		return nil, err
	}

	for i := range accList {
    acc := &accList[i] // Grab a pointer reference to the actual slice slot

    ledgerSum, err := a.dailyBalRepo.SumBalByAcc(acc.AccNo)
    if err != nil {
        return nil, err
    }

    acc.CloseBal += ledgerSum
	}

	return accList, nil
}

func (a*AccountService) CreateAccountService(request *dtos.CreateAccountRequest) (*models.AccountBalance, error) {
	var accountInfo models.AccountBalance

	category, err := a.assetRepo.FindBySymbol(request.Asset)
	if err != nil {
		return nil, err
	}

	accountInfo.AccNo=request.Uid+request.Asset
	accountInfo.Asset=request.Asset
	accountInfo.Uid=request.Uid
	accountInfo.WorkingBal=0
	accountInfo.Category = *category
	accountInfo.CreateAt=time.Now()
	accountInfo.ClosingDate=time.Now()

	if err := a.repo.Create(accountInfo); err != nil {
		return nil, err
	}

	return &accountInfo, nil
}