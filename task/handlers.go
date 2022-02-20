package task

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jonasala/delicious-friday/common"
	"github.com/jonasala/delicious-friday/db"
	"github.com/jonasala/delicious-friday/workorder"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(router fiber.Router) {
	group := router.Group("/tasks")
	group.Get("/", list)
	group.Get("/:id", get)
	group.Put("/:id", update)
	group.Delete("/:id", delete)
	group.Post("/", create)
}

func list(c *fiber.Ctx) error {
	return c.SendString("list")
}

func get(c *fiber.Ctx) error {
	return c.SendString("get")
}

func update(c *fiber.Ctx) error {
	return c.SendString("update")
}

func delete(c *fiber.Ctx) error {
	return c.SendString("delete")
}

func create(c *fiber.Ctx) error {
	reqPayload := new(CreateUpdateRequest)
	if err := c.BodyParser(reqPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := reqPayload.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var wo *workorder.WorkOrder
	if reqPayload.WorkOrder != "" {
		var err error
		wo, err = workorder.GetWorkOrder(c.Context(), reqPayload.WorkOrder)
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "work order not found",
			})
		} else if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	task := Task{
		WorkOrder:   wo,
		Description: reqPayload.Description,
		CreatedAt:   time.Now(),
		CreatedBy:   common.UserFromContext(c).Username,
	}
	result, err := db.DB.Collection("tasks").InsertOne(c.Context(), task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	task.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(task)
}
