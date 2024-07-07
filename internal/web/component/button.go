package component

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type ButtonParams struct {
	Type     buttonType
	Color    color
	Size     size
	Block    bool
	Class    string
	Children []gomponents.Node
}

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

	colorClasses := func() string {
		blackClasses := "bg-black text-white"

		switch params.Color {
		case ColorBlack:
			return blackClasses
		default:
			return blackClasses
		}
	}

	sizeClasses := func() string {
		smClasses := "text-sm p-1"
		mdClasses := "text-base p-2"
		lgClasses := "text-lg p-4"

		switch params.Size {
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

	return html.Button(
		html.Type(params.Type.Value),
		components.Classes{
			"rounded cursor-pointer":           true,
			"flex justify-center items-center": true,
			"block w-full":                     params.Block,
			colorClasses():                     true,
			sizeClasses():                      true,
			params.Class:                       params.Class != "",
		},
		gomponents.Group(params.Children),
	)
}
