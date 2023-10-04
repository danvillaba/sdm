package provider

import (
	"sdm/middle"
	inspectorview "sdm/module/inspector/view"
	sessionview "sdm/module/session/view"
	"sdm/ui"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/dig"
)

type appParams struct {
	dig.In
	Inspector *inspectorview.InspectorRouter
	Session   *sessionview.SessionRouter
}

func createApp(in appParams) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middle.ErrorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/api")

	in.Inspector.Register(api.Group("/inspector"))
	in.Session.Register(api.Group("/session"))

	app.Use("*", ui.UiFS)

	return app
}
