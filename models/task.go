package models
import(
	
	"go.mongodb.org/mongo-driver/bson/primitive"
	) 

type Task struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"userid"`
	Title  string             `bson:"title"`
	Description string        `bson:"description"`
}