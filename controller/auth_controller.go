package controller

import (
	"go-framework-learing/models"
	"go-framework-learing/repository"
	"go-framework-learing/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, "Invalid input: "+err.Error(), 400)
		return
	}

	// 🔥 Hash password
	hashedPassword, err := utils.GenerateToken(user.Password)
	if err != nil {
		utils.Error(c, "Failed to hash password", 500)
		return
	}
	user.Password = hashedPassword

	user.CreatedAt = time.Now()

	err = repository.CreateUser(&user)
	if err != nil {
		utils.Error(c, err.Error(), 500)
		return
	}

	utils.Success(c, gin.H{
		"message": "User registered successfully",
	})
}

func Login (c *gin.Context){
	var input models.User

	c.BindJSON(&input)

	user,err := repository.GetuserByEmail(input.Email)

	if err!=nil{
		utils.Error(c,"User not found",http.StatusNotFound)
		return
	}

	if user.Password != input.Password{
		utils.Error(c , "Invalid Credentails",401)
		return
	}

	token,_:= utils.GenerateToken((user.ID).String())

	utils.Success(c,gin.H{
		"token":token,
		
	})

}