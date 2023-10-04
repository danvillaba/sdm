package middle

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	if err == sql.ErrNoRows {
		return c.Status(404).SendString(err.Error())
	}

	return c.Status(code).SendString(err.Error())
}
