package main

import (
	"edart-core/app/routes"
	"edart-core/app/services"
	"edart-core/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cobSvc := services.GetCobService()
	defer cobSvc.Stop()
	router := gin.Default()
	database.Connect()
	routes.Router(router)
	router.Run(":8090")
}