package page

import (
	"github.com/eduardolat/generate-logo/internal/web/alpine"
	"github.com/eduardolat/generate-logo/internal/web/component"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func indexPreview() gomponents.Node {
	return html.Div(
		html.ID("preview-container"),

		components.Classes{
			"pattern-paper rounded flex-grow w-full select-none": true,
			"relative border border-gray-300":                    true,
			"h-[calc(100dvh-90px)] max-h-[calc(1080px-90px)]":    true,
		},
		html.Div(
			html.Class("absolute top-1 left-2"),
			component.H2Text("Preview"),
		),
		html.Div(
			html.ID("preview-container"),

			components.Classes{
				"absolute w-full h-full top-0 left-0": true,
				"flex justify-center items-center":    true,
			},

			indexPreviewLogo(),
		),
	)
}

func indexPreviewLogo() gomponents.Node {

	return html.Div(
		alpine.XBind("style", "previewSizeStyle"),
		alpine.XHTML("previewSvg"),
	)
}
