package component

import (
	"github.com/maragudk/gomponents"
	gcomponents "github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// HxLoadingSm returns a small loading indicator.
func HxLoadingSm(centered bool, id ...string) gomponents.Node {
	return hxLoading(centered, SizeSm, id...)
}

// HxLoadingMd returns a loading indicator.
func HxLoadingMd(centered bool, id ...string) gomponents.Node {
	return hxLoading(centered, SizeMd, id...)
}

// HxLoadingLg returns a large loading indicator.
func HxLoadingLg(centered bool, id ...string) gomponents.Node {
	return hxLoading(centered, SizeLg, id...)
}

func hxLoading(centered bool, size size, id ...string) gomponents.Node {
	style := func() string {
		styleSm := "width: 15px; height: 15px;"
		styleMd := "width: 25px; height: 25px;"
		styleLg := "width: 40px; height: 40px;"

		switch size {
		case SizeSm:
			return styleSm
		case SizeMd:
			return styleMd
		case SizeLg:
			return styleLg
		default:
			return styleMd
		}
	}()

	pickedID := ""
	if len(id) > 0 {
		pickedID = id[0]
	}

	return html.Div(
		gomponents.If(
			pickedID != "",
			html.ID(pickedID),
		),
		html.Class("htmx-indicator"),
		html.Div(
			gcomponents.Classes{
				"flex justify-center items-center": centered,
				"w-full h-full":                    true,
			},
			html.Div(
				html.Style(style),
				gcomponents.Classes{
					"border-black border-t-transparent animate-spin rounded-full": true,
					"border-[2px]": size == SizeSm,
					"border-[3px]": size == SizeMd,
					"border-[5px]": size == SizeLg,
				},
			),
		),
	)
}
