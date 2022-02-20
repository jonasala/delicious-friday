package workorder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jonasala/delicious-friday/common"
	"github.com/jonasala/delicious-friday/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterRoutes(router fiber.Router) {
	group := router.Group("/work-orders")
	group.Get("/", list)
	group.Get("/:id", get)
	group.Post("/", create)
	group.Delete("/:id", delete)
}

func create(c *fiber.Ctx) error {
	reqPayload := new(CreateRequest)
	if err := c.BodyParser(reqPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := reqPayload.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	file, err := c.FormFile("document")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if mime, err := common.DetectContentType(file); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	} else if !strings.HasPrefix(mime, "image/") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid file type, only images are accepted",
		})
	}

	filepath := fmt.Sprintf("files/work-orders/%v%v", uuid.New().String(), filepath.Ext(file.Filename))
	if err := c.SaveFile(file, "./"+filepath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	wo := WorkOrder{
		Number:    reqPayload.Number,
		FileURL:   filepath,
		CreatedAt: time.Now(),
		CreatedBy: common.UserFromContext(c).Username,
	}
	result, err := db.DB.Collection("work_orders").InsertOne(c.Context(), wo)
	if err != nil {
		os.Remove("./" + filepath)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	wo.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(wo)
}

func get(c *fiber.Ctx) error {
	wo, err := GetWorkOrder(c.Context(), c.Params("id"))
	if err == mongo.ErrNoDocuments {
		return c.SendStatus(fiber.StatusNotFound)
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(wo)
}

func list(c *fiber.Ctx) error {
	prefix := c.Query("prefix")
	filter := bson.M{}
	if prefix != "" {
		filter = bson.M{
			"number": bson.M{
				"$regex": primitive.Regex{
					Pattern: "^" + prefix,
					Options: "i",
				},
			},
		}
	}

	options := options.Find()
	options.SetSort(bson.M{"created_at": -1})
	options.SetLimit(5)

	cursor, err := db.DB.Collection("work_orders").Find(c.Context(), filter, options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	workOrders := []WorkOrder{}
	if err := cursor.All(c.Context(), &workOrders); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(workOrders)
}

func delete(c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	result := db.DB.Collection("work_orders").FindOneAndDelete(c.Context(), bson.M{"_id": objectID})
	if result.Err() == mongo.ErrNoDocuments {
		return c.SendStatus(fiber.StatusNotFound)
	} else if result.Err() != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
