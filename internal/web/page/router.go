package page

import (
	"github.com/labstack/echo/v4"
)

type handlers struct{}

func MountRouter(parent *echo.Group) {
	h := handlers{}

	parent.GET("/", h.indexHandler)
}
