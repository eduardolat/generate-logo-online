package page

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/eduardolat/generate-logo-online/internal/icons"
	"github.com/eduardolat/generate-logo-online/internal/util/echoutil"
	"github.com/eduardolat/generate-logo-online/internal/web/alpine"
	"github.com/eduardolat/generate-logo-online/internal/web/component"
	"github.com/eduardolat/generate-logo-online/internal/web/htmx"
	"github.com/labstack/echo/v4"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func (h *handlers) searchIconHandler(c echo.Context) error {
	const pageSize = 7 * 8

	query := c.QueryParam("q")
	page, err := strconv.Atoi(c.QueryParam("p"))
	if err != nil {
		page = 1
	}

	foundIcons, hasNextPage, nextPage := icons.PaginateIcons(page, pageSize, query)
	return echoutil.RenderGomponent(
		c, http.StatusOK,
		searchIconResults(foundIcons, query, hasNextPage, nextPage),
	)
}

func searchIconResults(
	foundIcons []icons.Icon, query string, hasNextPage bool, nextPage int,
) gomponents.Node {
	iconsLen := len(foundIcons)

	buttons := []gomponents.Node{}
	for i, icon := range foundIcons {
		buttons = append(buttons, html.Div(
			html.Button(
				html.Type("button"),
				html.Class("btn btn-outline btn-square btn-lg w-full"),
				alpine.XOn("click", `setOriginalSVG('`+icon.SVG+`')`),
				html.Title(icon.Name),
				icon.Icon(html.Class("size-8")),
			),

			gomponents.If(
				i == iconsLen-1 && hasNextPage,
				gomponents.Group([]gomponents.Node{
					htmx.HxGet(fmt.Sprintf(
						"/icons?q=%s&p=%d", query, nextPage,
					)),
					htmx.HxTrigger("intersect once"),
					htmx.HxSwap("afterend"),
					htmx.HxIndicator("#icon-picker-indicator"),
				}),
			),
		))
	}

	return component.RenderableGroup(buttons)
}
