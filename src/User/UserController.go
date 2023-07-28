package User

import (
	"context"
	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
	"os"
)

func login(c *fiber.Ctx) error {
	supabase := supa.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	ctx := context.Background()

	p := new(SignDto)
	if err := c.BodyParser(p); err != nil {
		return c.Status(500).JSON(supa.ErrorResponse{Code: 500, Message: "err"})
	}

	user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{Email: p.Email, Password: p.Password})
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.Status(200).JSON(user)
}
