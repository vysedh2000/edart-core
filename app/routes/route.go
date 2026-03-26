package routes

import (
	"edart-core/app/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router*gin.Engine) {
	userController := controllers.NewUserController()
	accController := controllers.NewAccController()
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.GET("/accSummary/:id", accController.AccSummary)

	router.POST("/user/create", userController.CreateNewUser)

	router.POST("/acc/create", accController.CreateNewAccount)

}