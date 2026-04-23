package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	ctx , cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	uri := os.Getenv("MONGO_URI")

	client , err := mongo.Connect(ctx,options.Client().ApplyURI(uri))

	if err!=nil{
		panic("Mongo connected failed : " + err.Error())
	}

	if err = client.Ping(ctx,nil); err!=nil{
		panic("Mongo pind failed "+ err.Error())
	}

	fmt.Println("Mongodb connected successfully")



	DB = client.Database(os.Getenv("MONGO_DB"))
}