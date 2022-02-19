package workorder

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkOrder struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Number    string             `bson:"number" json:"number"`
	FileURL   string             `bson:"file_url" json:"file_url"`
	CreatedBy string             `bson:"created_by" json:"created_by"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
