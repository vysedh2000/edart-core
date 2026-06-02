package controllers

import (
	"edart-core/app/dtos"
	"edart-core/app/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	service*services.AccountService
}

func NewAccController() *AccountController {
	return &AccountController{service:services.NewAccountService()}
}

func (c *AccountController) AccSummary(ctx *gin.Context) {
	id := ctx.Param("id")


	fmt.Println("allo",id)

	// if err:= ctx.ShouldBindJSON(&request); err != nil {
	// 	ctx.JSON(400, gin.H{
	// 		"message": "Invalid body request!",
	// 	})
	// 	return
	// }

	accDetail, err := c.service.AccSummaryService(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Account not found!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"assetList": accDetail,
	})
}

func (c*AccountController) CreateNewAccount(ctx*gin.Context) {
	var request dtos.CreateAccountRequest

	if err:= ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Request Body",
		})
		return
	}

	acc, err := c.service.CreateAccountService(&request)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "Error to create account!",
			"error": err.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"message": "Account Created!",
		"accountID": acc.AccNo,
		"asset": acc.Asset,
		"userID": acc.Uid,
	})
}