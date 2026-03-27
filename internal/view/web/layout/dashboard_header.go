package layout

import (
	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func dashboardHeader(lang string) nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[lang][key]; ok {
			return val
		}
		return key
	}

	return nodx.Header(
		nodx.ClassMap{
			"sticky top-0 z-50":                 true,
			"space-x-4 p-4 min-w-max":           true,
			"w-[full] bg-base-200 shadow-sm":    true,
			"flex items-center justify-between": true,
		},
		nodx.Div(
			nodx.Class("flex justify-start items-center space-x-2"),
			component.ChangeThemeButton(component.ChangeThemeButtonParams{
				Position: component.DropdownPositionBottom,
				Size:     component.SizeSm,
				Language: lang,
			}),
			component.ChangeLanguageButton(component.ChangeLanguageButtonParams{
				Position: component.DropdownPositionBottom,
				Size:     component.SizeSm,
				Language: lang,
			}),
		),
		nodx.Div(
			nodx.Class("flex justify-end items-center space-x-2"),
			nodx.Div(
				htmx.HxGet(pathutil.BuildPath("/dashboard/health-button?lang="+lang)),
				htmx.HxSwap("outerHTML"),
				htmx.HxTrigger("load once"),
			),
			nodx.Button(
				htmx.HxPost(pathutil.BuildPath("/auth/logout")),
				htmx.HxDisabledELT("this"),
				nodx.Class("btn btn-ghost btn-neutral"),
				nodx.SpanEl(nodx.Text(t("Log out"))),
				lucide.LogOut(),
			),
		),
	)
}
