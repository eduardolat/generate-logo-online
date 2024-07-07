package page

import (
	"github.com/eduardolat/generate-logo/internal/web/component"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func indexEditorBackground() gomponents.Node {
	return html.Div(
		component.H2Text("Background"),
	)
}
