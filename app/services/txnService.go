package services

import (
	"edart-core/app/dtos"
	"edart-core/app/models"
	"edart-core/app/repositories"
	"errors"
	"strconv"
	"strings"
	"time"
)

type TxnService struct{
	repo * repositories.TxnRepository
	othSer * OtherService
}

func NewTxnService() * TxnService{
	return &TxnService{repo: repositories.NewTxnRepo(), 
		othSer: NewOtherService(), }
}

func (t*TxnService) FundTxnTransfer(request *dtos.FundTxnRequest) (string, error){

	credExists, err := t.othSer.FindAccService(request.CreditAcc)
	if !credExists {
		return "Transaction Fail", errors.New("Credit Account is not found")
	}
	if err != nil {
		return "Transaction Fail", err
	}

	debExists, err := t.othSer.FindAccService(request.DebitAcc)

	if !debExists {
		return "Transaction Fail", errors.New("Debit Account is not found")
	}

	var debTxn models.TxnSuccess
	var credTxn models.TxnSuccess

	creditUser := strings.TrimSuffix(request.CreditAcc, request.Asset)
	debitUser := strings.TrimSuffix(request.DebitAcc, request.Asset)

	//credit txn
	credTxn.TxnId = request.BatchId+"C"
	credTxn.RefId = request.BatchId
	credTxn.AccId = request.CreditAcc
	credTxn.Narative = request.Narative
	credTxn.TxnCode = request.TxnCode
	credTxn.Asset = request.Asset
	credTxn.TxnType = "C"
	credTxn.Amount, _ = strconv.ParseFloat(request.Amount, 64)
	credTxn.UserId = creditUser
	credTxn.TxnDate = time.Now()

	//debit txn
	debTxn.TxnId = request.BatchId+"D"
	debTxn.RefId = request.BatchId
	debTxn.AccId = request.DebitAcc
	debTxn.Narative = request.Narative
	debTxn.TxnCode = request.TxnCode
	debTxn.Asset = request.Asset
	debTxn.Amount = -1 * credTxn.Amount
	debTxn.UserId = debitUser
	debTxn.TxnDate = time.Now()
	
	var txns = []models.TxnSuccess{debTxn, credTxn}

	//save to db
	t.repo.CreateTxnBatch(txns)

	return request.BatchId, nil
}