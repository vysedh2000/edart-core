package controllers

import (
	"edart-core/app/dtos"
	"edart-core/app/services"

	"github.com/gin-gonic/gin"
)

type TxnController struct {
	service * services.TxnService
	othService * services.OtherService
	cobService * services.CobService
}

func NewTxnController() * TxnController{
	return &TxnController{service: services.NewTxnService(),
	othService: services.NewOtherService(), 
	cobService: services.GetCobService(),}
}

func (c *TxnController) CobHandler(ctx *gin.Context) {
	c.cobService.DailyCob()
	ctx.JSON(200, gin.H{
		"status": "success",
		"message": "Cob is running",
	})
}

func (c *TxnController) FundTxnTransfer(ctx *gin.Context){
	var request dtos.FundTxnRequest
	if err:= ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Request Body",
		})
		return
	}
	txnId, err := c.service.FundTxnTransfer(&request)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to process transaction",
			"error": err.Error(),
		})
		return
	}
	
	ctx.JSON(200, gin.H{
		"message": "Transaction successful",
		"txnId": txnId,
	})
}