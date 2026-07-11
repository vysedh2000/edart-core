package services

import (
	"edart-core/app/dtos"
	"edart-core/app/models"
	"edart-core/app/repositories"
	"fmt"
	"time"
)

type AccountService struct{
	repo *repositories.AccRepository
	txnRepo * repositories.TxnRepository
	assetRepo * repositories.AssetRepository
}

func NewAccountService() *AccountService{
	return &AccountService{repo: repositories.NewAccountRepo(),
		txnRepo: repositories.NewTxnRepo(),
		assetRepo: repositories.NewAssetRepo(),
	}
}

func (a*AccountService) AccSummaryService(accNo string) (any, error) {
	accList, err := a.repo.AllAccInquiry(accNo)
	if err != nil {
		return nil, err
	}

	fmt.Print("AccList",accList)

	for i := range accList {
    acc := &accList[i] // Grab a pointer reference to the actual slice slot

	fmt.Print("AccNo", acc.AccNo)
    ledgerSum, err := a.txnRepo.GetBalByAcc(acc.AccNo)
	fmt.Print("AccBal", i,ledgerSum)
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