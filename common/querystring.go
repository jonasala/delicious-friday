package common

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PaginationInfo(c *fiber.Ctx) (int, int, error) {
	page, err := strconv.Atoi(c.Query("page", "0"))
	if err != nil {
		return 0, 0, err
	}
	itemsPerPage, err := strconv.Atoi(c.Query("items_per_page", "0"))
	if err != nil {
		return 0, 0, err
	}

	return page, itemsPerPage, nil
}
