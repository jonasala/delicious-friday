package task

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jonasala/delicious-friday/common"
	"github.com/jonasala/delicious-friday/db"
	"github.com/jonasala/delicious-friday/workorder"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	search := c.Query("search")
	all := c.Query("all") == "1"
	page, perPage, err := common.PaginationInfo(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resultset := common.Resultset{
		Page:         page,
		ItemsPerPage: perPage,
	}

	filters := bson.M{}
	if !all {
		filters["created_by"] = common.UserFromContext(c).Username
	}
	if search != "" {
		filters["$or"] = bson.A{
			bson.M{"created_by": search},
			bson.M{"work_order.number": bson.M{
				"$regex": primitive.Regex{
					Pattern: "^" + search,
					Options: "i",
				},
			},
			},
			bson.M{"description": bson.M{
				"$regex": primitive.Regex{
					Pattern: search,
					Options: "i",
				},
			}},
		}
	}

	total, err := db.DB.Collection("tasks").CountDocuments(c.Context(), filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resultset.TotalItems = int(total)
	resultset.Calculate()

	options := options.Find()
	options.SetSort(bson.M{"created_at": -1})
	options.SetLimit(int64(resultset.ItemsPerPage))
	options.SetSkip(int64(resultset.Offset))

	cursor, err := db.DB.Collection("tasks").Find(c.Context(), filters, options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tasks := []Task{}
	if err := cursor.All(c.Context(), &tasks); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	resultset.Data = tasks
	return c.JSON(resultset)

}

func get(c *fiber.Ctx) error {
	task, err := GetTask(c.Context(), c.Params("id"))
	if err == mongo.ErrNoDocuments {
		return c.SendStatus(fiber.StatusNotFound)
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(task)
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
