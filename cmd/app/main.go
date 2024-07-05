package main

import (
	"github.com/eduardolat/generate-logo/internal/logger"
	"github.com/eduardolat/generate-logo/internal/web"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.HideBanner = true
	app.HidePort = true

	web.MountRouter(app.Group(""))

	logger.Info("starting server", logger.KV{
		"port": "8088",
		"url":  "http://localhost:8088",
	})
	if err := app.Start(":8088"); err != nil {
		logger.FatalError("failed to start server", logger.KV{
			"error": err,
		})
	}
}
