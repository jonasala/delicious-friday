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
	"github.com/jonasala/delicious-friday/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterRoutes(router fiber.Router) {
	group := router.Group("/work-orders")
	group.Get("/", list)
	group.Get("/:id", get)
	group.Post("/", create)
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
	result, err := mongo.DB.Collection("work_orders").InsertOne(c.Context(), wo)
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
	return c.SendString("get work order")
}

func list(c *fiber.Ctx) error {
	return c.SendString("list work orders")
}
