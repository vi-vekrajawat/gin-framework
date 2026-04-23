package repository

import (
	"context"
	"go-framework-learing/config"
	"go-framework-learing/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(task *models.Task) error{
	collection := config.DB.Collection("tasks")
	ctx , cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	_,err:= collection.InsertOne(ctx,task)
	return err
}

func GetTask(userID primitive.ObjectID) ([]models.Task, error) {
	collection := config.DB.Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"userid": userID} // ✅ correct field

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}