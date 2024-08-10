package page

import (
	"github.com/eduardolat/generate-logo/internal/web/alpine"
	"github.com/eduardolat/generate-logo/internal/web/component"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func indexEditorBg() gomponents.Node {
	return html.Div(
		html.Class("space-y-2"),
		indexEditorBgType(),
		indexEditorBgRadius(),
		indexEditorBgColor(),
		indexEditorBgGradient(),
	)
}

func indexEditorBgType() gomponents.Node {
	return html.Div(
		html.Class("form-control w-full"),
		html.Label(
			html.Class("label"),
			html.For("bgType"),
			component.SpanText("Background type"),
		),
		html.Select(
			html.ID("bgType"),
			html.Class("select select-bordered"),
			alpine.XModel("bgType"),
			html.Option(html.Value("solid"), gomponents.Text("Solid")),
			html.Option(html.Value("gradient"), gomponents.Text("Gradient")),
		),
	)
}

func indexEditorBgRadius() gomponents.Node {
	return html.Div(
		html.Class("form-control w-full"),
		html.Label(
			html.Class("label"),
			html.For("bgRadius"),
			component.SpanText("Border radius"),
		),
		html.Input(
			html.ID("bgRadius"),
			html.Class("range"),
			html.Type("range"),
			html.Min("0"),
			html.Max("50"),
			html.Step("1"),
			alpine.XModel("bgRadius"),
		),
	)
}

func indexEditorBgColor() gomponents.Node {
	return html.Div(
		html.Class("form-control w-full"),
		alpine.XShow("bgType === 'solid'"),

		html.Label(
			html.Class("label"),
			html.For("bgColor"),
			component.SpanText("Background color"),
		),

		html.Input(
			html.ID("bgColor"),
			html.Class("w-full h-[60px]"),
			html.Type("color"),
			html.Class("w-full"),
			alpine.XModel("bgColor"),
		),
	)
}

func indexEditorBgGradient() gomponents.Node {
	return html.Div(
		alpine.XShow("bgType === 'gradient'"),
		html.Div(
			html.Label(html.For("bgGradientStart"), component.SpanText("Gradient Start Color:")),
			html.Input(
				html.ID("bgGradientStart"),
				html.Type("color"),
				html.Class("w-full"),
				alpine.XModel("bgGradientStart"),
			),
		),
		html.Div(
			html.Label(html.For("bgGradientEnd"), component.SpanText("Gradient End Color:")),
			html.Input(
				html.ID("bgGradientEnd"),
				html.Type("color"),
				html.Class("w-full"),
				alpine.XModel("bgGradientEnd"),
			),
		),
		html.Div(
			html.Label(html.For("bgGradientType"), component.SpanText("Gradient Type:")),
			html.Select(
				html.ID("bgGradientType"),
				alpine.XModel("bgGradientType"),
				html.Option(html.Value("linear"), gomponents.Text("Linear")),
				html.Option(html.Value("radial"), gomponents.Text("Radial")),
			),
		),
		html.Div(
			alpine.XShow("bgGradientType === 'linear'"),
			html.Label(html.For("bgGradientAngle"), component.SpanText("Gradient Angle:")),
			html.Input(
				html.ID("bgGradientAngle"),
				html.Type("range"),
				html.Min("0"),
				html.Max("360"),
				html.Step("1"),
				html.Class("range"),
				alpine.XModel("bgGradientAngle"),
			),
		),
		html.Div(
			html.Label(html.For("bgGradientCutLine"), component.SpanText("Gradient Cut Line:")),
			html.Input(
				html.ID("bgGradientCutLine"),
				html.Type("range"),
				html.Min("0"),
				html.Max("100"),
				html.Step("1"),
				html.Class("range"),
				alpine.XModel("bgGradientCutLine"),
			),
		),
		html.Div(
			html.Label(html.For("bgGradientBlur"), component.SpanText("Gradient blur:")),
			html.Input(
				html.ID("bgGradientBlur"),
				html.Type("range"),
				html.Min("0"),
				html.Max("100"),
				html.Step("1"),
				html.Class("range"),
				alpine.XModel("bgGradientBlur"),
			),
		),
	)
}
