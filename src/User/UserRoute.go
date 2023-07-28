package User

import "github.com/gofiber/fiber/v2"

func Route(route fiber.Router) {
	route.Post("/login", login)
}
