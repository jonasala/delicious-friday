package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jonasala/delicious-friday/auth"
	"github.com/jonasala/delicious-friday/mongo"
)

func main() {
	mongoURI := os.Getenv("DELICIOUS_FRIDAY_MONGO_URI")
	dbName := os.Getenv("DELICIOUS_FRIDAY_DBNAME")
	httpPort := os.Getenv("DELICIOUS_FRIDAY_HTTP_PORT")
	jwtSecret := os.Getenv("DELICIOUS_FRIDAY_JWT_SECRET")

	if httpPort == "" {
		log.Fatalln("DELICIOUS_FRIDAY_HTTP_PORT is undefined")
	}
	if mongoURI == "" {
		log.Fatalln("DELICIOUS_FRIDAY_MONGO_URI is undefined")
	}
	if dbName == "" {
		log.Fatalln("DELICIOUS_FRIDAY_DBNAME is undefined")
	}
	if jwtSecret == "" {
		log.Fatalln("DELICIOUS_FRIDAY_JWT_SECRET is undefined")
	}

	if err := mongo.Connect(mongoURI, dbName); err != nil {
		log.Fatalln("unable to connect with mongodb", err)
	}

	if err := auth.LoadUsers(); err != nil {
		log.Fatalln("unable to load users", err)
	}

	app := fiber.New()

	apiGroup := app.Group("/api")
	authGroup := apiGroup.Group("/auth")
	restrictedGroup := apiGroup.Group("/restricted")
	restrictedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
	}))

	auth.RegisterPublicRoutes(authGroup)

	if err := serveUI(app); err != nil {
		log.Fatalln("unable to serve ui.", err)
	}

	app.Listen(":" + httpPort)
}

func serveUI(app *fiber.App) error {
	if _, err := os.Stat("./ui/dist"); err != nil {
		return err
	}
	app.Static("/", "./ui/dist")
	return nil
}
