package page

import (
	"github.com/eduardolat/generate-logo/internal/web/alpine"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func indexEditorBg() gomponents.Node {
	return html.Div(
		indexEditorBgRadius(),
		indexEditorBgColor(),
	)
}

func indexEditorBgRadius() gomponents.Node {
	return html.Div(
		html.Input(
			html.Type("range"),
			html.Min("0"),
			html.Max("50"),
			html.Step("1"),
			html.Class("w-full"),
			alpine.XModel("bgRadius"),
		),
	)
}

func indexEditorBgColor() gomponents.Node {
	return html.Div(
		html.Input(
			html.Type("color"),
			html.Class("w-full"),
			alpine.XModel("bgColor"),
		),
	)
}
