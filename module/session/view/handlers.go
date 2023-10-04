package sessionview

import (
	sessionv1 "sdm/pb/sdm/session/v1"

	"github.com/gofiber/fiber/v2"
)

func (s *SessionRouter) Connect(c *fiber.Ctx) error {
	ctx := c.UserContext()
	body := &sessionv1.ConnectRequest{}
	c.BodyParser(body)

	res, err := s.svc.Connect(ctx, body)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (s *SessionRouter) Disconnect(c *fiber.Ctx) error {
	ctx := c.UserContext()
	body := &sessionv1.DisconnectRequest{
		Token: c.Get("authorization"),
	}
	res, err := s.svc.Disconnect(ctx, body)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (s *SessionRouter) Info(c *fiber.Ctx) error {
	ctx := c.UserContext()
	body := &sessionv1.InfoRequest{
		Token: c.Get("authorization"),
	}

	res, err := s.svc.Info(ctx, body)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
