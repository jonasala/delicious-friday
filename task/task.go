package task

import (
	"time"

	"github.com/jonasala/delicious-friday/workorder"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID   `bson:"id" json:"id"`
	WorkOrder   *workorder.WorkOrder `bson:"work_order" json:"work_order"`
	Description string               `bson:"description" json:"description"`
	CreatedBy   string               `bson:"created_by" json:"created_by"`
	CreatedAt   time.Time            `bson:"created_at" json:"created_at"`
}
