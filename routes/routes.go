package routes

import (
	"go-framework-learing/controller"
	"go-framework-learing/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine){
	auth:=router.Group("/auth")
	{
		auth.POST("/ragister",controller.Register)
		auth.POST("/login",controller.Login)
	}

	task := router.Group("/task")
	task.Use(middleware.AuthMiddleware())
	{
		task.POST("/",controller.CreateTask)
		task.GET("/",controller.GetTask)
	}
}