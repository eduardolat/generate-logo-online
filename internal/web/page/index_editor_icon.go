package page

import (
	"github.com/eduardolat/generate-logo/internal/web/alpine"
	"github.com/eduardolat/generate-logo/internal/web/component"
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func indexEditorIcon() gomponents.Node {
	return html.Div(
		indexEditorIconPicker(),
		indexEditorIconColor(),
		indexEditorIconSize(),
		indexEditorIconRotate(),
	)
}

func indexEditorIconPicker() gomponents.Node {
	mo := component.Modal(component.ModalParams{
		Size:          component.SizeMd,
		Title:         "Pick an icon",
		HTMXIndicator: "icon-picker-indicator",
	})

	button := component.Button(component.ButtonParams{
		Block: true,
		Children: []gomponents.Node{
			mo.OpenerAttr,
			component.SpanText("Pick an icon"),
			lucide.Pointer(html.Class("ml-1")),
		},
	})

	return html.Div(
		mo.HTML,
		button,
	)
}

func indexEditorIconSize() gomponents.Node {
	return html.Div(
		html.Input(
			html.Type("range"),
			html.Min("1"),
			html.Max("100"),
			html.Step("1"),
			html.Class("w-full"),
			alpine.XModel("iconSize"),
		),
	)
}

func indexEditorIconColor() gomponents.Node {
	return html.Div(
		html.Input(
			html.Type("color"),
			html.Class("w-full"),
			alpine.XModel("iconColor"),
		),
	)
}

func indexEditorIconRotate() gomponents.Node {
	return html.Div(
		html.Input(
			html.Type("range"),
			html.Min("-180"),
			html.Max("180"),
			html.Step("1"),
			html.Class("w-full"),
			alpine.XModel("iconRotate"),
		),
	)
}
