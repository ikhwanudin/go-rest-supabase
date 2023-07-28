package main

import (
	"go-rest/src/Photo"
	"go-rest/src/User"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error when loading .env")
		}
	}

	port := os.Getenv("PORT")
	app := fiber.New()

	api := app.Group("/api/")
	User.UserRoute(api)
	Photo.PhotoRoute(api)

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
