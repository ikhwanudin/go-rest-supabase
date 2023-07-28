package Photo

import "github.com/gofiber/fiber/v2"

func PhotoRoute(route fiber.Router) {
	route.Get("/photos", getAllPhotos)
	route.Get("/photo/:id", getPhotoByid)
}
