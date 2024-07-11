package component

import (
	"github.com/google/uuid"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// InputControlParams is the properties for the InputControl component.
type InputControlParams struct {
	// ID is the HTML id attribute value for the input control.
	ID string
	// Name is the HTML name attribute value for the input control.
	Name string
	// Type is the type attribute of the input. Default is "text".
	Type inputType
	// Title is the text label associated with the input control.
	Title string
	// Placeholder is the HTML placeholder attribute value, displayed when the input is empty.
	Placeholder string
	// HelpText is additional text displayed below the input, usually providing guidance.
	HelpText string
	// Color is the color class applied to the input control for styling.
	Color color
	// Required, when true, indicates that the input is mandatory for form submission.
	Required bool
	// Error, when true, applies error styling to the input control.
	Error bool
	// Children are additional HTML nodes to be included inside the input element.
	Children []gomponents.Node
}

// InputControl is a component that renders a group of nodes
// that are used to present a form input control.
func InputControl(props InputControlParams) gomponents.Node {
	id := props.ID
	if id == "" {
		id = "input-" + uuid.NewString()
	}

	hasHelperText := props.HelpText != ""
	describedById := id + "-help-text"

	color := props.Color
	if props.Error {
		color = ColorRed
	}

	return html.Label(
		components.Classes{
			"input-control-container block space-y-1": true,
		},
		html.For(id),

		html.Span(
			components.Classes{
				"input-control-label block": true,
				"text-error":                props.Error,
			},

			gomponents.If(
				props.Title != "",
				gomponents.Text(props.Title),
			),

			gomponents.If(props.Required, html.Span(
				html.Class("text-error"),
				gomponents.Text(" *"),
			)),
		),
		Input(InputParams{
			Color:       color,
			Block:       true,
			Name:        props.Name,
			Type:        props.Type,
			Placeholder: props.Placeholder,
			Children: []gomponents.Node{
				html.ID(id),
				gomponents.If(props.Required, html.Required()),
				gomponents.If(hasHelperText, html.Aria("describedby", describedById)),
				gomponents.Group(props.Children),
			},
		}),
		gomponents.If(hasHelperText,
			html.Span(
				components.Classes{
					"input-control-help-text block text-sm": true,
					"text-error":                            props.Error,
				},
				html.ID(describedById),
				gomponents.Text(props.HelpText),
			),
		),
	)
}
