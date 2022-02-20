package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jonasala/delicious-friday/auth"
)

func UserFromContext(c *fiber.Ctx) auth.User {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return auth.User{
		Name:     claims["name"].(string),
		Username: claims["username"].(string),
	}
}
