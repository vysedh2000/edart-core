package routes

import (
	"edart-core/app/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	userController := controllers.NewUserController()
	accController := controllers.NewAccController()
	txnController := controllers.NewTxnController()

	router.Use(gin.Recovery())

	

	core := router.Group("/core")
	{

		core.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
		core.GET("/accSummary/:id", accController.AccSummary)

		core.POST("/user/create", userController.CreateNewUser)

		core.POST("/acc/create", accController.CreateNewAccount)

		core.POST("/txn/fundTxn", txnController.FundTxnTransfer)

		core.GET("/cob", txnController.CobHandler)
	}
}