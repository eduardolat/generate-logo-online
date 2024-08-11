package page

import (
	"github.com/eduardolat/generate-logo-online/internal/web/alpine"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func indexEditor() gomponents.Node {
	navBtn := func(tab string, text string) gomponents.Node {
		return html.Div(
			alpine.XBind("class", "editorTab == '"+tab+"' && 'border-neutral'"),
			html.Class("border-b border-base-300"),
			html.Button(
				alpine.XOn("click", "editorTab = '"+tab+"'"),
				alpine.XBind("class", "editorTab != '"+tab+"' && 'btn-ghost'"),
				html.Class("btn btn-neutral btn-block rounded-none no-animation"),
				gomponents.Text(text),
			),
		)
	}

	return html.Aside(
		components.Classes{
			"w-full max-w-[250px] bg-base-100 rounded-box overflow-hidden": true,
			"h-[calc(100dvh-90px)] max-h-[calc(1080px-100px)]":             true,
			"flex-none flex flex-col":                                      true,
		},
		html.Nav(
			html.Class("grid grid-cols-2"),
			navBtn("icon", "Icon"),
			navBtn("background", "Background"),
		),
		html.Div(
			html.Class("flex-grow p-2"),
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
