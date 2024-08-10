package page

import (
	"github.com/eduardolat/generate-logo/internal/web/alpine"
	"github.com/eduardolat/generate-logo/internal/web/component"
	"github.com/eduardolat/generate-logo/internal/web/htmx"
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
		Content: []gomponents.Node{
			html.Div(
				html.Class("space-y-4 h-full"),

				component.Input(component.InputParams{
					Name:        "q",
					Type:        component.InputTypeText,
					Color:       component.ColorBlack,
					Placeholder: "Search for an icon",
					Block:       true,
					Children: []gomponents.Node{
						htmx.HxGet("/icons"),
						htmx.HxTrigger("input changed delay:500ms, search"),
						htmx.HxTarget("#icon-picker-results"),
						htmx.HxIndicator("#icon-picker-indicator"),
					},
				}),

				html.Div(
					html.ID("icon-picker-results"),
					htmx.HxGet("/icons"),
					htmx.HxTrigger("intersect once"),
					htmx.HxIndicator("#icon-picker-indicator"),

					component.SpinnerContainerLg(),
				),
			),
		},
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
			html.Class("range"),
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
			html.Class("range"),
			alpine.XModel("iconRotate"),
		),
	)
}
