package page

import (
	"github.com/eduardolat/generate-logo/internal/web/component"
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func indexEditorDownload() gomponents.Node {
	return html.Div(
		html.Button(
			components.Classes{
				"bg-black rounded-b w-full py-2":      true,
				"text-white font-bold cursor-pointer": true,
				"flex justify-center items-center":    true,
			},
			component.SpanText("Download"),
			lucide.CloudDownload(html.Class("ml-1")),
		),
	)
}
