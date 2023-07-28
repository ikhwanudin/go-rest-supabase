package Photo

import (
	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
	"os"
)

func getAllPhotos(c *fiber.Ctx) error {
	supabase := supa.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	type Photos struct {
		Id     int8   `json:"id"`
		UserId string `json:"user_id"`
		Photo  string `json:"photo"`
	}

	var results []Photos

	err := supabase.DB.From("photo").Select("*").Execute(&results)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.Status(200).JSON(results)
}

func getPhotoByid(c *fiber.Ctx) error {
	id := c.Params("id")

	supabase := supa.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	type Photos struct {
		Id     int8   `json:"id"`
		UserId string `json:"user_id"`
		Photo  string `json:"photo"`
	}

	var results []Photos

	err := supabase.DB.From("photo").Select("*").Eq("id", id).Execute(&results)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(200).JSON(results)
}

//func uploadPhoto(c *fiber.Ctx) error {
//	file, err := c.FormFile("photo")
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": true,
//			"msg":   err.Error(),
//		})
//	}
//
//	buffer, err := file.Open()
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": true,
//			"msg":   err.Error(),
//		})
//	}
//	defer buffer.Close()
//
//	supabase := supa.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
//
//	objectName := file.Filename
//	fileBuffer := buffer
//
//	info, err := supabase.Storage.From("photo").Upload(objectName, fileBuffer)
//	if err != nil {
//		if err != nil {
//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//				"error": true,
//				"msg":   err.Error(),
//			})
//		}
//	}
//
//	if info != nil {
//		if err != nil {
//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//				"error": true,
//				"msg":   err.Error(),
//			})
//		}
//	}
//
//	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
//
//}
