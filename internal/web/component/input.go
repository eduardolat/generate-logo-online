package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type InputParams struct {
	// Name is the name attribute of the input.
	Name string
	// Type is the type attribute of the input. Default is "text".
	Type inputType
	// Placeholder is the placeholder attribute of the input.
	Placeholder string
	// Class is the additional class name(s) for the input.
	Class string
	// Default is neutral.
	Color color
	// Block makes the input 100% width. Default is false.
	Block bool
	// Children are the contents of the input.
	Children []gomponents.Node
}

func Input(props InputParams) gomponents.Node {
	color := ColorBlack
	if props.Color.Value != "" {
		color = props.Color
	}

	colorClasses := inputColorClasses(color)

	return html.Input(
		components.Classes{
			"disabled:cursor-not-allowed":                        true,
			"rounded bg-white":                                   true,
			"placeholder:text-black placeholder:text-opacity-70": true,
			"focus:ring focus:ring-opacity-50":                   true,
			"w-full":                                             props.Block,
			colorClasses:                                         true,
			props.Class:                                          props.Class != "",
		},
		gomponents.If(
			props.Name != "",
			html.Name(props.Name),
		),
		gomponents.If(
			props.Type.Value != "",
			html.Type(props.Type.Value),
		),
		gomponents.If(
			props.Placeholder != "",
			html.Placeholder(props.Placeholder),
		),
		gomponents.Group(props.Children),
	)
}

func inputColorClasses(color color) string {
	blackClasses := "border-black focus:border-black focus:ring-black"
	redClasses := "border-red-500 focus:border-red-500 focus:ring-red-500"

	switch color {
	case ColorBlack:
		return blackClasses
	case ColorRed:
		return redClasses
	default:
		return blackClasses
	}
}
