package page

import (
	"github.com/eduardolat/generate-logo-online/internal/web/alpine"
	"github.com/eduardolat/generate-logo-online/internal/web/component"
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func indexEditorDownload() gomponents.Node {
	mo := component.Modal(component.ModalParams{
		Size:  component.SizeLg,
		Title: "Download your logo",
		Content: []gomponents.Node{
			html.Div(
				html.Class("space-y-4"),

				html.Div(
					html.Class("grid grid-cols-3 gap-2"),
					html.Div(
						html.Class("flex justify-center"),
						alpine.XHTML("dlPreviewSvg"),
					),
					html.Div(
						html.Class("flex justify-center"),
						alpine.XHTML("dlPreviewSvgWhite"),
					),
					html.Div(
						html.Class("flex justify-center"),
						alpine.XHTML("dlPreviewSvgBlack"),
					),
				),

				html.Button(
					html.Type("button"),
					html.Class("btn btn-primary btn-block"),
					alpine.XOn("click", "downloadZippedLogos"),
					component.SpanText("Download"),
					lucide.CloudDownload(html.Class("ml-1")),
				),
			),
		},
	})

	return html.Div(
		mo.HTML,
		html.Button(
			html.Type("button"),
			html.Class("btn btn-primary btn-block rounded-none no-animation"),
			mo.OpenerAttr,
			component.SpanText("Download"),
			lucide.CloudDownload(html.Class("ml-1")),
		),
	)
}
