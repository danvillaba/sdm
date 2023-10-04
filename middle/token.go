package middle

import "github.com/gofiber/fiber/v2"

func TokenMiddleware(c *fiber.Ctx) error {
	token := c.Get("authorization")
	if token == "" {
		return c.SendStatus(401)
	}

	return c.Next()
}
