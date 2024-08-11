package page

import (
	"github.com/eduardolat/generate-logo-online/internal/web/alpine"
	"github.com/eduardolat/generate-logo-online/internal/web/component"
	"github.com/eduardolat/generate-logo-online/internal/web/htmx"
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

				html.Input(
					html.Class("input input-bordered w-full"),
					html.Type("text"),
					html.Name("q"),
					html.Placeholder("Search for an icon"),
					htmx.HxGet("/icons"),
					htmx.HxTrigger("input changed delay:500ms, search"),
					htmx.HxTarget("#icon-picker-results"),
					htmx.HxIndicator("#icon-picker-indicator"),
				),

				html.Div(
					html.ID("icon-picker-results"),
					html.Class("grid grid-cols-7 gap-2"),
					htmx.HxGet("/icons"),
					htmx.HxTrigger("intersect once"),
					htmx.HxIndicator("#icon-picker-indicator"),

					html.Div(
						html.Class("col-span-7"),
						component.SpinnerContainerLg(),
					),
				),
			),
		},
	})

	button := html.Button(
		html.Type("button"),
		html.Class("btn btn-neutral btn-block"),
		mo.OpenerAttr,
		component.SpanText("Pick an icon"),
		lucide.Pointer(html.Class("ml-1")),
	)

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
