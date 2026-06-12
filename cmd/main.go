package main

import (
	"crud-go/config"
	"crud-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
