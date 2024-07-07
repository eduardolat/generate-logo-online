package layout

import (
	"github.com/eduardolat/generate-logo/internal/web/component"
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func baseHeader() gomponents.Node {
	return html.Header(
		components.Classes{
			"normal-width h-[60px]":                       true,
			"flex justify-between items-center space-x-2": true,
		},
		html.Div(
			html.Class("flex justify-start items-center space-x-2"),
			html.Img(
				html.Class("h-full max-h-[40px] w-auto rounded"),
				html.Src("/images/logo.svg"),
			),
			html.Span(
				html.Class("text-lg font-bold desk:text-2xl"),
				gomponents.Text("Generate Logo Online"),
			),
		),
		html.Nav(
			html.Class("hidden desk:inline"),
			html.Ul(
				html.Class("text-xl flex justify-end items-center space-x-2"),
				html.Li(
					html.A(
						html.Class("flex items-center space-x-1 hover:underline"),
						html.Href("https://x.com/eduardoolat"),
						html.Target("_blank"),
						lucide.Twitter(html.Class("size-6")),
						component.SpanText("X (Twitter)"),
					),
				),
				html.Li(
					html.A(
						html.Class("flex items-center space-x-1 hover:underline"),
						html.Href("https://github.com/eduardolat/generate-logo"),
						html.Target("_blank"),
						lucide.Github(html.Class("size-6")),
						component.SpanText("GitHub"),
					),
				),
			),
		),
	)
}
