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
	title := "Generate logo online"
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
				gomponents.Attr("type", "module"),
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
			html.Script(
				html.Src("https://cdn.jsdelivr.net/npm/jszip@3.10.1/dist/jszip.min.js"),
				html.Defer(),
			),
			html.Script(
				html.Src("https://cdn.jsdelivr.net/npm/theme-change@2.0.2/index.js"),
			),
		},
		Body: []gomponents.Node{
			alpine.XData(`gloapp`),
			alpine.XCloak(),
			html.Class("space-y-2 bg-base-300 w-[100dvw] h-[100dvh]"),
			baseHeader(),
			gomponents.Group(props.Body),
		},
	})
}
