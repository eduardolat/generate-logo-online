package layout

import (
	"github.com/eduardolat/generate-logo/internal/web/alpine"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type BaseProps struct {
	Title string
	Body  []gomponents.Node
}

func Base(props BaseProps) gomponents.Node {
	title := "Generate logo"
	if props.Title != "" {
		title = title + " - " + props.Title
	}

	return components.HTML5(components.HTML5Props{
		Language: "en",
		Title:    title,
		Head: []gomponents.Node{
			html.Link(
				html.Rel("shortcut icon"),
				html.Href("/favicon.ico"),
			),

			html.Link(
				html.Rel("stylesheet"),
				html.Href("/css/style.css"),
			),

			html.Script(
				html.Src("/js/app.js"),
				html.Defer(),
			),
			html.Script(
				html.Src("https://cdn.jsdelivr.net/npm/alpinejs@3.14.0/dist/cdn.min.js"),
				html.Defer(),
			),
			html.Script(
				html.Src("https://cdn.jsdelivr.net/npm/htmx.org@1.9.12/dist/htmx.min.js"),
				html.Defer(),
			),
		},
		Body: []gomponents.Node{
			alpine.XData(`gloapp`),
			alpine.XCloak(),
			html.Class("space-y-2"),
			baseHeader(),
			gomponents.Group(props.Body),
		},
	})
}
