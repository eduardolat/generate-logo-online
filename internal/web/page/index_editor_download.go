package page

import (
	"github.com/eduardolat/generate-logo/internal/web/component"
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func indexEditorDownload() gomponents.Node {
	return html.Div(
		component.Button(component.ButtonParams{
			Type:  component.ButtonTypeButton,
			Color: component.ColorBlack,
			Size:  component.SizeMd,
			Block: true,
			Class: "rounded-t-none",
			Children: []gomponents.Node{
				component.SpanText("Download"),
				lucide.CloudDownload(html.Class("ml-1")),
			},
		}),
	)
}
