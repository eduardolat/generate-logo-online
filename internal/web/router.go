package web

import (
	"github.com/eduardolat/generate-logo/internal/web/pages"
	"github.com/eduardolat/generate-logo/internal/web/static"
	"github.com/labstack/echo/v4"
)

func MountRouter(parent *echo.Group) {
	parent.StaticFS("/", static.StaticFS)
	pages.MountRouter(parent.Group(""))
}
