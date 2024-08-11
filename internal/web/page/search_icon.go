package page

import (
	"net/http"

	"github.com/eduardolat/generate-logo-online/internal/icons"
	"github.com/eduardolat/generate-logo-online/internal/util/echoutil"
	"github.com/eduardolat/generate-logo-online/internal/web/alpine"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) searchIconHandler(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		allIcons := icons.GetAllIcons()
		return echoutil.RenderGomponent(
			c, http.StatusOK,
			searchIconResults(allIcons),
		)
	}

	foundIcons := icons.SearchIcons(query)
	return echoutil.RenderGomponent(
		c, http.StatusOK,
		searchIconResults(foundIcons),
	)
}

func searchIconResults(foundIcons []icons.Icon) gomponents.Node {
	return html.Div(
		html.Class("grid grid-cols-7 gap-2"),
		gomponents.Group(
			gomponents.Map(
				foundIcons,
				func(icon icons.Icon) gomponents.Node {
					return html.Div(
						html.Class("tooltip"),
						html.Button(
							html.Type("button"),
							html.Class("btn btn-outline btn-square btn-lg w-full"),
							alpine.XOn("click", `setOriginalSVG('`+icon.SVG+`')`),
							html.Title(icon.Name),
							icon.Icon(html.Class("size-8")),
						),
					)
				}),
		),
	)
}
