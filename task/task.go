package task

import (
	"context"
	"time"

	"github.com/jonasala/delicious-friday/db"
	"github.com/jonasala/delicious-friday/workorder"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID   `bson:"_id" json:"id"`
	WorkOrder   *workorder.WorkOrder `bson:"work_order" json:"work_order"`
	Description string               `bson:"description" json:"description"`
	CreatedBy   string               `bson:"created_by" json:"created_by"`
	CreatedAt   time.Time            `bson:"created_at" json:"created_at"`
}

func GetTask(ctx context.Context, id string) (*Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := db.DB.Collection("tasks").FindOne(ctx, bson.M{"_id": objectID})
	if result.Err() != nil {
		return nil, result.Err()
	}
	task := new(Task)
	result.Decode(task)
	return task, nil
}
