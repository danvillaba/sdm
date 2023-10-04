package sessionview

import (
	"sdm/middle"
	sessionsvc "sdm/module/session/service"

	"github.com/gofiber/fiber/v2"
)

type SessionRouter struct {
	svc *sessionsvc.SessionService
}

func CreateSessionRouter(
	svc *sessionsvc.SessionService,
) *SessionRouter {

	return &SessionRouter{svc}
}

func (s *SessionRouter) Register(router fiber.Router) {

	router.Post("/connect", s.Connect)
	router.Post("/disconnect", middle.TokenMiddleware, s.Disconnect)
	router.Get("/info", middle.TokenMiddleware, s.Info)
}
