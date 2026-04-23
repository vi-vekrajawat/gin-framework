package controller

import (
	"go-framework-learing/models"
	"go-framework-learing/repository"
	"go-framework-learing/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *gin.Context) {

	userID, _ := c.Get("user_id")

objID, err := primitive.ObjectIDFromHex(userID.(string))
if err != nil {
	utils.Error(c, "Invalid user ID", 400)
	return
}

var task models.Task
task.UserID = objID

	if err := c.ShouldBindJSON(&task); err != nil {
		utils.Error(c, "Invalid input: "+err.Error(), 400)
		return
	}


	err = repository.CreateTask(&task)
	if err != nil {
		utils.Error(c, "Failed to create task", 500)
		return
	}

	utils.Success(c, task)
}

func GetTask(c *gin.Context)  {
  
	userID, _ := c.Get("user_id")

objID, _ := primitive.ObjectIDFromHex(userID.(string))

	tasks,_ := repository.GetTask(objID)

	utils.Success(c , tasks)
	
	
}