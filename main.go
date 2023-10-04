package main

import (
	"log"
	"sdm/module/session"
	"sdm/provider"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func main() {
	provider.Init()

	log.Fatal(provider.Container.Invoke(start))
}

func start(app *fiber.App, man *session.SessionManager) error {
	db, _ := sqlx.Open("postgres", "postgres://postgres:admin@localhost/main?sslmode=disable")
	ses := session.Session{
		Id:       "testing",
		Db:       db,
		Driver:   "postgres",
		Database: "main",
	}
	man.Add(&ses)
	defer man.Close()

	return app.Listen(":8080")
}
