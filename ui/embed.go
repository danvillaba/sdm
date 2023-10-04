package ui

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed build/*
var buildDist embed.FS

var config = filesystem.Config{
	Root:         http.FS(buildDist),
	PathPrefix:   "build",
	Index:        "index.html",
	NotFoundFile: "index.html",
}

var UiFS = filesystem.New(config)
