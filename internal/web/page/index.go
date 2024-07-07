package page

import (
	"net/http"

	"github.com/eduardolat/generate-logo/internal/util/echoutil"
	"github.com/eduardolat/generate-logo/internal/web/layout"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) indexHandler(c echo.Context) error {
	return echoutil.RenderGomponent(c, http.StatusOK, indexPage())
}

func indexPage() gomponents.Node {
	return layout.Base(layout.BaseProps{
		Body: []gomponents.Node{
			html.Main(
				html.Class("normal-width flex space-x-2"),
				indexEditor(),
				indexPreview(),
			),
		},
	})
}
