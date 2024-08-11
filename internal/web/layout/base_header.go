package layout

import (
	"github.com/eduardolat/generate-logo-online/internal/web/alpine"
	"github.com/eduardolat/generate-logo-online/internal/web/component"
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func baseHeader() gomponents.Node {
	return html.Header(
		components.Classes{
			"normal-width h-[60px] bg-base-300":           true,
			"flex justify-between items-center space-x-2": true,
		},
		html.Div(
			html.Class("flex justify-start items-center space-x-2"),
			html.Img(
				html.Class("h-full max-h-[40px] w-auto rounded-btn"),
				html.Src("/images/logo.svg"),
			),
			html.Span(
				html.Class("text-lg font-bold desk:text-2xl"),
				gomponents.Text("Generate Logo Online"),
			),
		),
		html.Nav(
			html.Class("hidden desk:inline"),
			html.Ul(
				html.Class("text-xl flex justify-end items-center space-x-2"),
				html.Li(
					switchThemeButton(),
				),
				html.Li(
					html.A(
						html.Class("btn btn-neutral space-x-1"),
						html.Href("https://x.com/eduardoolat"),
						html.Target("_blank"),
						lucide.Twitter(),
						component.SpanText("X (Twitter)"),
					),
				),
				html.Li(
					html.A(
						html.Class("btn btn-neutral space-x-1"),
						html.Href("https://github.com/eduardolat/generate-logo-online"),
						html.Target("_blank"),
						lucide.Github(),
						component.SpanText("GitHub"),
					),
				),
			),
		),
	)
}

func switchThemeButton() gomponents.Node {
	return html.Div(
		alpine.XData(`{
			theme: "",
			
			getCurrentTheme() {
				const el = document.querySelector("html");
				const theme = el.getAttribute("data-theme");
				if (theme) {
					this.theme = theme;
					return
				}
				this.theme = "system";
			},

			init() {
				setTimeout(() => {
					this.getCurrentTheme();
				}, 200)
			}
		}`),
		alpine.XCloak(),
		alpine.XOn("click", "getCurrentTheme()"),
		alpine.XOn("click.outside", "getCurrentTheme()"),

		html.Class("dropdown dropdown-end"),
		html.Div(
			html.TabIndex("0"),
			html.Role("button"),
			html.Class("btn btn-neutral space-x-1"),

			html.Div(
				html.Class("inline-block size-4"),
				lucide.Laptop(alpine.XShow(`theme === "system"`)),
				lucide.Sun(alpine.XShow(`theme === "light"`)),
				lucide.Moon(alpine.XShow(`theme === "dark"`)),
			),

			component.SpanText("Theme"),
			lucide.ChevronDown(),
		),
		html.Ul(
			html.TabIndex("0"),
			components.Classes{
				"dropdown-content bg-base-200 rounded z-[1] w-[150px]": true,
				"p-2 space-y-2 mt-2 shadow-sm":                         true,
			},
			html.Li(
				html.Button(
					html.Data("set-theme", ""),
					html.Class("btn btn-neutral btn-block"),
					html.Type("button"),
					lucide.Laptop(html.Class("mr-1")),
					component.SpanText("System"),
				),
			),
			html.Li(
				html.Button(
					html.Data("set-theme", "light"),
					html.Class("btn btn-neutral btn-block"),
					html.Type("button"),
					lucide.Sun(html.Class("mr-1")),
					component.SpanText("Light"),
				),
			),
			html.Li(
				html.Button(
					html.Data("set-theme", "dark"),
					html.Class("btn btn-neutral btn-block"),
					html.Type("button"),
					lucide.Moon(html.Class("mr-1")),
					component.SpanText("Dark"),
				),
			),
		),
	)
}
