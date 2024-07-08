package page

import (
	"github.com/eduardolat/generate-logo/internal/web/alpine"
	"github.com/eduardolat/generate-logo/internal/web/component"
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

				component.Button(component.ButtonParams{
					Type:  component.ButtonTypeButton,
					Color: component.ColorBlack,
					Size:  component.SizeMd,
					Block: true,
					Children: []gomponents.Node{
						alpine.XOn("click", "downloadZippedLogos"),
						component.SpanText("Download"),
						lucide.CloudDownload(html.Class("ml-1")),
					},
				}),
			),
		},
	})

	return html.Div(
		mo.HTML,
		component.Button(component.ButtonParams{
			Type:  component.ButtonTypeButton,
			Color: component.ColorBlack,
			Size:  component.SizeMd,
			Block: true,
			Class: "rounded-t-none",
			Children: []gomponents.Node{
				mo.OpenerAttr,
				component.SpanText("Download"),
				lucide.CloudDownload(html.Class("ml-1")),
			},
		}),
	)
}
