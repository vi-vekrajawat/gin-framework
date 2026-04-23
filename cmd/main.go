package main

import (
	"go-framework-learing/config"
	"go-framework-learing/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	router := gin.Default()  // gin.New() , gin.Logger(), gin.Recovery() we will get inbuilt in default 

	config.ConnectDB()

	routes.SetUpRoutes(router)

	router.Run(":8081")

}