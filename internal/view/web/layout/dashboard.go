package layout

import (
	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
)

type DashboardParams struct {
	Title string
	Body  []nodx.Node
}

func Dashboard(reqCtx reqctx.Ctx, params DashboardParams) nodx.Node {
	title := "PG Back Web"
	pageTitle := params.Title
	if pageTitle != "" {
		t := func(key string) string {
			if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
				return val
			}
			return key
		}
		pageTitle = t(pageTitle)
		title = pageTitle + " - " + title
	}

	if reqCtx.IsHTMXBoosted {
		body := append(params.Body, nodx.TitleEl(nodx.Text(title)))
		return component.RenderableGroup(body)
	}

	body := nodx.Group(
		nodx.ClassMap{
			"w-screen h-screen bg-base-200":      true,
			"flex justify-start overflow-hidden": true,
		},
		dashboardAside(reqCtx.Language),
		nodx.Div(
			nodx.Class("flex-grow overflow-y-auto"),
			dashboardHeader(reqCtx.Language),
			nodx.Main(
				nodx.Id("dashboard-main"),
				nodx.Class("p-4"),
				nodx.Group(params.Body...),
			),
		),
	)

	return commonHtmlDoc(title, body)
}
