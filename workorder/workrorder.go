package workorder

import (
	"context"
	"time"

	"github.com/jonasala/delicious-friday/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkOrder struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Number    string             `bson:"number" json:"number"`
	FileURL   string             `bson:"file_url" json:"file_url"`
	CreatedBy string             `bson:"created_by" json:"created_by"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func GetWorkOrder(ctx context.Context, id string) (*WorkOrder, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := db.DB.Collection("work_orders").FindOne(ctx, bson.M{"_id": objectID})
	if result.Err() != nil {
		return nil, result.Err()
	}
	wo := new(WorkOrder)
	result.Decode(wo)
	return wo, nil
}
