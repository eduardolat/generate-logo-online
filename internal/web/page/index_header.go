package page

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func indexHeader() gomponents.Node {
	return html.Header(
		components.Classes{
			"normal-width h-[70px]":                       true,
			"flex justify-between items-center space-x-2": true,
		},
		html.Div(
			html.Class("flex justify-start items-center space-x-2"),
			html.Img(
				html.Class("h-full max-h-[50px] w-auto"),
				html.Src("/images/logo.svg"),
			),
			html.Span(
				html.Class("text-lg font-bold desk:text-2xl"),
				gomponents.Text("Generate Logo Online"),
			),
		),
		html.Nav(),
	)
}
