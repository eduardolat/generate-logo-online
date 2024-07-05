package echoutil

import (
	"github.com/labstack/echo/v4"
)

func JsonError(c echo.Context, code int, err error, message ...string) error {
	if len(message) > 0 {
		return c.JSON(code, map[string]string{
			"error":   err.Error(),
			"message": message[0],
		})
	}

	return c.JSON(code, map[string]string{
		"error": err.Error(),
	})
}
