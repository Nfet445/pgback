package component

import (
	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

type ChangeLanguageButtonParams struct {
	Position    dropdownPosition
	AlignsToEnd bool
	Size        size
	Language    string
}

func ChangeLanguageButton(params ChangeLanguageButtonParams) nodx.Node {
	isEN := params.Language == i18n.LangEN || params.Language == ""
	isRU := params.Language == i18n.LangRU

	return nodx.Div(
		alpine.XData("alpineChangeLanguageButton()"),
		alpine.XCloak(),

		nodx.ClassMap{
			"dropdown":        true,
			"dropdown-end":    params.AlignsToEnd,
			"dropdown-right":  params.Position == DropdownPositionRight,
			"dropdown-left":   params.Position == DropdownPositionLeft,
			"dropdown-top":    params.Position == DropdownPositionTop,
			"dropdown-bottom": params.Position == DropdownPositionBottom,
		},
		nodx.Div(
			nodx.Tabindex("0"),
			nodx.Role("button"),
			nodx.ClassMap{
				"btn btn-neutral": true,
				"btn-sm":          params.Size == SizeSm,
				"btn-lg":          params.Size == SizeLg,
			},

			nodx.Div(
				nodx.Class("flex space-x-1"),
				lucide.Globe(nodx.Class("size-4")),
				SpanText("Language"),
				lucide.ChevronDown(),
			),
		),
		nodx.Ul(
			nodx.Tabindex("0"),
			nodx.ClassMap{
				"dropdown-content":                   true,
				"bg-base-100":                        true,
				"rounded-btn shadow-md":              true,
				"z-[1] w-[150px] p-2 space-y-2 my-2": true,
			},
			nodx.Li(
				nodx.Button(
					htmx.HxPost(pathutil.BuildPath("/api/language/en")),
					htmx.HxSwap("none"),
					htmx.HxOn("after-request", "window.location.reload()"),
					nodx.ClassMap{
						"btn btn-block": true,
						"btn-primary":   isEN,
						"btn-neutral":   !isEN,
						"btn-sm":        params.Size == SizeSm,
						"btn-lg":        params.Size == SizeLg,
					},
					nodx.Type("button"),
					SpanText("English"),
				),
			),
			nodx.Li(
				nodx.Button(
					htmx.HxPost(pathutil.BuildPath("/api/language/ru")),
					htmx.HxSwap("none"),
					htmx.HxOn("after-request", "window.location.reload()"),
					nodx.ClassMap{
						"btn btn-block": true,
						"btn-primary":   isRU,
						"btn-neutral":   !isRU,
						"btn-sm":        params.Size == SizeSm,
						"btn-lg":        params.Size == SizeLg,
					},
					nodx.Type("button"),
					SpanText("Russian"),
				),
			),
		),
	)
}
