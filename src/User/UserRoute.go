package User

import "github.com/gofiber/fiber/v2"

func UserRoute(route fiber.Router) {
	route.Get("/authors", getUser)
	route.Post("/login", login)
}
