package services

import (
	"context"
	"edart-core/app/dtos"
	"edart-core/app/models"
	"edart-core/app/repositories"
	"edart-core/database"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
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

	ctx := context.Background()
	val, err := database.Rdb.Get(ctx, request.DebitAcc).Result()

	if val == request.DebitAcc {
		for i := 0; i < 5; i++ {
		val, err := database.Rdb.Get(ctx, request.DebitAcc).Result()
		if err == redis.Nil {
			break
		}
		if err != nil {
			return "", err
		}
		if val != "" {
			if i == 4 {
				return "", errors.New("account is locked")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	} 

	//before start txn lock first
	database.Rdb.Set(ctx, request.DebitAcc, request.DebitAcc, 50)

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

	//after txn done unlock account
	database.Rdb.Del(ctx, request.DebitAcc)

	return request.BatchId, nil
}