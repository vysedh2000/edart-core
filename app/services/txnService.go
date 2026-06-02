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
	ledger * repositories.DailyLedgerRepository
	othSer * OtherService
}

func NewTxnService() * TxnService{
	return &TxnService{repo: repositories.NewTxnRepo(), 
		othSer: NewOtherService(), ledger: 
		repositories.NewDailyBalRepository()}
}

func (t*TxnService) FundTxnTransfer(request *dtos.FundTxnRequest) (string, error){

	credExists, err := t.othSer.FindAccService(request.CreditAcc)
	
	if !credExists {
		return "Transaction Fail", errors.New("Credit Account is not found")
	}

	debExists, err := t.othSer.FindAccService(request.DebitAcc)

	if !debExists {
		return "Transaction Fail", errors.New("Debit Account is not found")
	}

	txnId, err := t.repo.GetNextTxnId()
	if err != nil {
		return "", err
	}

	var credEntry models.DailyLedger
	var debEntry models.DailyLedger
	var debTxn models.TxnSuccess
	var credTxn models.TxnSuccess

	creditUser := strings.TrimSuffix(request.CreditAcc, request.Asset)
	debitUser := strings.TrimSuffix(request.DebitAcc, request.Asset)

	//credit entry
	credEntry.TxnNo = txnId+"C"
	credEntry.RefId = txnId
	credEntry.AccNo = request.CreditAcc
	credEntry.Amount, _ = strconv.ParseFloat(request.Amount, 64)
	credEntry.Asset = request.Asset
	credEntry.ValueDate = time.Now()

	//debit entry
	debEntry.TxnNo = txnId+"D"
	debEntry.RefId = txnId
	debEntry.AccNo = request.DebitAcc
	debEntry.Amount, _ = strconv.ParseFloat(request.Amount, 64)
	debEntry.Amount = -1*debEntry.Amount

	var entries = []models.DailyLedger{debEntry, credEntry}

	//credit txn
	credTxn.TxnId = txnId+"C"
	credTxn.RefId = txnId
	credTxn.AccId = request.CreditAcc
	credTxn.Narative = request.Narative
	credTxn.TxnCode = request.TxnCode
	credTxn.Asset = request.Asset
	credTxn.TxnType = "C"
	credTxn.Amount = credEntry.Amount
	credTxn.UserId = creditUser

	//debit txn
	debTxn.TxnId = txnId+"D"
	debTxn.RefId = txnId
	debTxn.AccId = request.DebitAcc
	debTxn.Narative = request.Narative
	debTxn.TxnCode = request.TxnCode
	debTxn.Asset = request.Asset
	debTxn.Amount = debEntry.Amount
	debTxn.UserId = debitUser

	var txns = []models.TxnSuccess{debTxn, credTxn}

	//save to db
	t.ledger.CreateDailyBalBatch(entries)
	t.repo.CreateTxnBatch(txns)

	return txnId, nil
}