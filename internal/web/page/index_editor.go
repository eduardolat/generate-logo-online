package page

import (
	"github.com/eduardolat/generate-logo/internal/web/alpine"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func indexEditor() gomponents.Node {
	li := func(pos string, tab string, text string) gomponents.Node {
		return html.Li(
			alpine.XOn("click", "editorTab = '"+tab+"'"),
			alpine.XBind("class", "editorTab == '"+tab+"' ? 'active' : 'inactive'"),
			components.Classes{
				"text-center py-2 cursor-pointer":  true,
				"font-bold border border-base-300": true,
				"[&.active]:border-black":          true,
				"[&.active]:bg-black":              true,
				"[&.active]:text-white":            true,
			},
			gomponents.Text(text),
		)
	}

	return html.Aside(
		components.Classes{
			"w-full max-w-[250px] bg-base-100 rounded-box overflow-hidden": true,
			"h-[calc(100dvh-90px)] max-h-[calc(1080px-100px)]":             true,
			"flex-none flex flex-col":                                      true,
		},
		html.Nav(
			html.Ul(
				html.Class("grid grid-cols-2"),
				li("left", "icon", "Icon"),
				li("right", "background", "Background"),
			),
		),
		html.Div(
			html.Class("flex-grow p-2 border border-y-0 border-base-300"),
			html.Div(
				alpine.XShow("editorTab == 'icon'"),
				indexEditorIcon(),
			),
			html.Div(
				alpine.XShow("editorTab == 'background'"),
				indexEditorBg(),
			),
		),
		html.Div(
			indexEditorDownload(),
		),
	)
}
