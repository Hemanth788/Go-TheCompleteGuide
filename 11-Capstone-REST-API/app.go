package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.com/rest-api/db"
	"go.com/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/", getHome)

	routes.RegisterEventRoutes(server)

	server.Run(":8080")
}

func getHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Homey! I'm home!"})
}