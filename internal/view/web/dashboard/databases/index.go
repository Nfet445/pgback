package databases

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	reqCtx := reqctx.GetCtx(c)
	return echoutil.RenderNodx(c, http.StatusOK, indexPage(reqCtx))
}

func indexPage(reqCtx reqctx.Ctx) nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	content := []nodx.Node{
		nodx.Div(
			nodx.Class("flex justify-between items-start"),
			component.H1Text(t("Databases")),
			nodx.Div(
				nodx.Class("flex space-x-2"),
				connectPostgresButton(reqCtx),
				createDatabaseButton(reqCtx),
			),
		),
		component.CardBox(component.CardBoxParams{
			Class: "mt-4",
			Children: []nodx.Node{
				nodx.Div(
					nodx.Class("overflow-x-auto"),
					nodx.Table(
						nodx.Class("table text-nowrap"),
						nodx.Thead(
							nodx.Tr(
								nodx.Th(nodx.Class("w-1")),
								nodx.Th(component.SpanText(t("Name"))),
								nodx.Th(component.SpanText(t("Version"))),
								nodx.Th(component.SpanText(t("Connection string"))),
								nodx.Th(component.SpanText(t("Created at"))),
							),
						),
						nodx.Tbody(
							component.SkeletonTr(8),
							htmx.HxGet(pathutil.BuildPath("/dashboard/databases/list?page=1")),
							htmx.HxTrigger("load"),
						),
					),
				),
			},
		}),
	}

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: t("Databases"),
		Body:  content,
	})
}
