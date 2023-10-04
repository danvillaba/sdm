package inspectorview

import (
	inspectorsvc "sdm/module/inspector/service"

	"github.com/gofiber/fiber/v2"
)

func (s *InspectorRouter) inspectorMiddle(c *fiber.Ctx) error {
	session := s.manager.Get(c.Get("authorization"))
	if session == nil {
		return c.SendStatus(401)
	}

	current, err := inspectorsvc.CreateInspectorService(session.Db, session.Driver)
	if err != nil {
		return err
	}

	c.Locals("inspector", current)
	return c.Next()
}
