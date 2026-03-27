package layout

import (
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func dashboardAside(lang string) nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[lang][key]; ok {
			return val
		}
		return key
	}

	return nodx.Aside(
		nodx.Id("dashboard-aside"),
		nodx.ClassMap{
			"flex-none h-[100dvh] bg-base-300 shadow-sm p-4": true,
			"overflow-y-auto overflow-x-hidden":              true,
		},

		nodx.Div(
			nodx.Class("space-y-4"),

			dashboardAsideItem(
				lucide.LayoutDashboard,
				t("Summary"),
				pathutil.BuildPath("/dashboard"),
				true,
			),

			dashboardAsideItem(
				lucide.Database,
				t("Databases"),
				pathutil.BuildPath("/dashboard/databases"),
				false,
			),

			dashboardAsideItem(
				lucide.HardDrive,
				t("Destinations"),
				pathutil.BuildPath("/dashboard/destinations"),
				false,
			),

			dashboardAsideItem(
				lucide.DatabaseBackup,
				t("Backup tasks"),
				pathutil.BuildPath("/dashboard/backups"),
				false,
			),

			dashboardAsideItem(
				lucide.List,
				t("Executions"),
				pathutil.BuildPath("/dashboard/executions"),
				false,
			),

			dashboardAsideItem(
				lucide.ArchiveRestore,
				t("Restorations"),
				pathutil.BuildPath("/dashboard/restorations"),
				false,
			),

			dashboardAsideItem(
				lucide.Webhook,
				t("Webhooks"),
				pathutil.BuildPath("/dashboard/webhooks"),
				false,
			),

			dashboardAsideItem(
				lucide.User,
				t("Profile"),
				pathutil.BuildPath("/dashboard/profile"),
				false,
			),

			dashboardAsideItem(
				lucide.Info,
				t("About"),
				pathutil.BuildPath("/dashboard/about"),
				false,
			),
		),
	)
}

func dashboardAsideItem(
	icon func(children ...nodx.Node) nodx.Node,
	text, link string, strict bool,
) nodx.Node {
	return nodx.A(
		alpine.XData(fmt.Sprintf("alpineDashboardAsideItem('%s', %t)", link, strict)),
		nodx.Class("block flex flex-col items-center justify-center group"),

		nodx.Href(link),
		htmx.HxBoost("true"),
		htmx.HxTarget("#dashboard-main"),
		htmx.HxSwap("transition:true show:unset"),

		nodx.Button(
			alpine.XBind("class", `{'btn-active': is_active}`),
			nodx.Class("btn btn-ghost btn-neutral btn-square group-hover:btn-active"),
			icon(nodx.Class("size-6")),
		),
		nodx.SpanEl(
			nodx.Class("text-xs"),
			nodx.Text(text),
		),
	)
}
