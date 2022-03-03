package auth

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RegisterPublicRoutes(router fiber.Router) {
	router.Post("/login", func(c *fiber.Ctx) error {
		request := LoginRequest{}
		if err := c.BodyParser(&request); err != nil {
			return err
		}

		for _, user := range Users {
			if user.Username == request.Username && user.Password == request.Password {
				claims := jwt.MapClaims{
					"username": user.Username,
					"name":     user.Name,
					"exp":      time.Now().Add(time.Hour * 72).Unix(),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				t, err := token.SignedString([]byte(os.Getenv("DELICIOUS_FRIDAY_JWT_SECRET")))
				if err != nil {
					return c.SendStatus(fiber.StatusInternalServerError)
				}

				return c.JSON(LoginResponse{
					Username: user.Username,
					Name:     user.Name,
					Token:    t,
				})
			}
		}
		return c.SendStatus(fiber.StatusUnauthorized)
	})
}
