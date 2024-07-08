package component

import (
	"fmt"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type ButtonParams struct {
	// Type is the type of the button. button, submit, or reset.
	// Default is button.
	Type buttonType
	// Color is the color of the button. Default is black.
	Color color
	// Size is the size of the button. Default is md.
	Size size
	// Block makes the button full width.
	Block bool
	// Circle makes the button a circle.
	Circle bool
	// Square makes the button a square.
	Square bool
	// Ghost makes the button a ghost button.
	Ghost bool
	// Outline makes the button an outline button.
	Outline bool
	// Class is an additional classes to add to the button.
	Class string
	// Children are the children of the button.
	Children []gomponents.Node
}

// Button renders a button.
func Button(params ButtonParams) gomponents.Node {
	if params.Type.Value == "" {
		params.Type = ButtonTypeButton
	}

	if params.Color.Value == "" {
		params.Color = ColorBlack
	}

	if params.Size.Value == "" {
		params.Size = SizeMd
	}

	colorClasses := buttonColorRegularClasses(params.Color)
	if params.Outline {
		colorClasses = buttonColorOutlineClasses(params.Color)
	}
	if params.Ghost {
		colorClasses = buttonColorGhostClasses(params.Color)
	}

	sizeClasses := buttonSizeRegularClasses(params.Size)
	if params.Square {
		sizeClasses = buttonSizeSquareClasses(params.Size)
	}
	if params.Circle {
		sizeClasses = buttonSizeCircleClasses(params.Size)
	}

	return html.Button(
		html.Type(params.Type.Value),
		components.Classes{
			"disabled:cursor-not-allowed disabled:opacity-50":            true,
			"rounded transition duration-100 border hover:bg-opacity-80": true,
			"inline-flex justify-center items-center":                    true,
			"w-full":     params.Block,
			colorClasses: true,
			sizeClasses:  true,
			params.Class: params.Class != "",
		},
		gomponents.Group(params.Children),
	)
}

func buttonColorRegularClasses(color color) string {
	blackClasses := "border-black bg-black text-white"

	switch color {
	case ColorBlack:
		return blackClasses
	default:
		return blackClasses
	}
}

func buttonColorOutlineClasses(color color) string {
	blackClasses := "border-black bg-transparent text-black hover:bg-black hover:text-white"

	switch color {
	case ColorBlack:
		return blackClasses
	default:
		return blackClasses
	}
}

func buttonColorGhostClasses(color color) string {
	blackClasses := "border-transparent bg-transparent text-black hover:border-black hover:bg-black hover:text-white"

	switch color {
	case ColorBlack:
		return blackClasses
	default:
		return blackClasses
	}
}

func buttonSizeRegularClasses(size size) string {
	smClasses := "text-sm px-2 py-1"
	mdClasses := "text-base px-4 py-2"
	lgClasses := "text-lg p-4 px-6 py-3"

	switch size {
	case SizeSm:
		return smClasses
	case SizeMd:
		return mdClasses
	case SizeLg:
		return lgClasses
	default:
		return mdClasses
	}
}

func buttonSizeSquareClasses(size size) string {
	smClasses := "w-6 h-6"
	mdClasses := "w-8 h-8"
	lgClasses := "w-10 h-10"

	switch size {
	case SizeSm:
		return smClasses
	case SizeMd:
		return mdClasses
	case SizeLg:
		return lgClasses
	default:
		return mdClasses
	}
}

func buttonSizeCircleClasses(size size) string {
	return fmt.Sprintf("rounded-full %s", buttonSizeSquareClasses(size))
}
