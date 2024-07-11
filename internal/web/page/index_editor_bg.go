package page

import (
	"github.com/eduardolat/generate-logo/internal/web/alpine"
	"github.com/eduardolat/generate-logo/internal/web/component"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func indexEditorBg() gomponents.Node {
	return html.Div(
		indexEditorBgType(),
		indexEditorBgRadius(),
		indexEditorBgColor(),
		indexEditorBgGradient(),
	)
}

func indexEditorBgType() gomponents.Node {
	return html.Div(
		html.Label(html.For("bgType"), component.SpanText("Background Type:")),
		html.Select(
			html.ID("bgType"),
			alpine.XModel("bgType"),
			html.Option(html.Value("solid"), gomponents.Text("Solid")),
			html.Option(html.Value("gradient"), gomponents.Text("Gradient")),
		),
	)
}

func indexEditorBgRadius() gomponents.Node {
	return html.Div(
		html.Label(html.For("bgRadius"), component.SpanText("Border Radius:")),
		html.Input(
			html.ID("bgRadius"),
			html.Type("range"),
			html.Min("0"),
			html.Max("50"),
			html.Step("1"),
			html.Class("w-full"),
			alpine.XModel("bgRadius"),
		),
	)
}

func indexEditorBgColor() gomponents.Node {
	return html.Div(
		alpine.XShow("bgType === 'solid'"),
		html.Label(html.For("bgColor"), component.SpanText("Background Color:")),
		html.Input(
			html.ID("bgColor"),
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
				html.Class("w-full"),
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
				html.Class("w-full"),
				alpine.XModel("bgGradientCutLine"),
			),
		),
	)
}
