package component

import (
	"bytes"

	"github.com/maragudk/gomponents"
)

// RenderableGroup renders a group of nodes without a parent element.
//
// This is because gomponents.Group() cannot be directly rendered and
// needs to be wrapped in a parent element.
func RenderableGroup(children []gomponents.Node) gomponents.Node {
	if len(children) == 0 {
		return gomponents.Raw("")
	}

	buf := bytes.Buffer{}
	for _, child := range children {
		err := child.Render(&buf)
		if err != nil {
			return gomponents.Raw("Error rendering group")
		}
	}
	return gomponents.Raw(buf.String())
}
