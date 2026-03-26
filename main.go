package main

import (
	"edart-core/app/routes"
	"edart-core/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Connect()
	routes.Router(router)
	router.Run(":8080")
}