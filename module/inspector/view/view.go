package inspectorview

import (
	"sdm/middle"
	"sdm/module/session"

	"github.com/gofiber/fiber/v2"
)

type InspectorRouter struct {
	manager *session.SessionManager
}

func CreateInspectorRouter(manager *session.SessionManager) *InspectorRouter {
	return &InspectorRouter{manager}
}

func (s *InspectorRouter) Register(router fiber.Router) {
	router.Use(middle.TokenMiddleware)
	router.Use(s.inspectorMiddle)

	router.Get("/databases", s.DatabaseList)
	router.Get("/objects", s.ObjectList)
	router.Get("/schemas", s.SchemaList)
	router.Get("/schemas/:schemaid/tables", s.SchemaTableList)
	router.Get("/schemas/:schemaid/tables/:tableid", s.TableInfo)
	router.Post("/query", s.Query)
}
