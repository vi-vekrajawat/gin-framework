package repository

import (
	"context"
	"go-framework-learing/config"
	"go-framework-learing/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user *models.User) error {

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(),time.Second*10)

	defer cancel()

	_ , err := collection.InsertOne(ctx , user)
	if err!=nil{
		panic("something went wrong while creating the user"+ err.Error())
	}

	return err
	// return  config.DB.Create(user).Error

}

func GetuserByEmail(email string) (*models.User,error){
	// var user models.User

	collection := config.DB.Collection("users")

	ctx , cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var user models.User

	err := collection.FindOne(ctx,bson.M{"email":email}).Decode(&user)

	return &user,err
}