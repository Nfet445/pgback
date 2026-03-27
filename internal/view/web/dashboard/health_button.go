package dashboard

import (
	"fmt"
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
)

func healthButtonHandler(servs *service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		reqCtx := reqctx.GetCtx(c)
		lang := reqCtx.Language

		databasesQty, err := servs.DatabasesService.GetDatabasesQty(ctx)
		if err != nil {
			return respondhtmx.ToastError(c, err.Error())
		}
		destinationsQty, err := servs.DestinationsService.GetDestinationsQty(ctx)
		if err != nil {
			return respondhtmx.ToastError(c, err.Error())
		}

		return echoutil.RenderNodx(c, http.StatusOK, healthButton(
			lang, databasesQty, destinationsQty,
		))
	}
}

func healthButton(
	lang string,
	databasesQty dbgen.DatabasesServiceGetDatabasesQtyRow,
	destinationsQty dbgen.DestinationsServiceGetDestinationsQtyRow,
) nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[lang][key]; ok {
			return val
		}
		return key
	}

	areDatabasesHealthy := databasesQty.Unhealthy == 0
	areDestinationsHealthy := destinationsQty.Unhealthy == 0
	isHealthy := areDatabasesHealthy && areDestinationsHealthy

	pingColor := component.ColorSuccess
	if !isHealthy {
		pingColor = component.ColorError
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: t("Health status"),
		Content: []nodx.Node{
			component.PText(`
				The health check for both databases and destinations runs automatically
				every 10 minutes, when PG Back Web starts, and when you click the
				"Test connection" button on each resource. You can see additional
				information and error messages by clicking the health check button
				for each resource.
			`),
			nodx.Table(
				nodx.Class("table mt-2"),
				nodx.Thead(
					nodx.Tr(
						nodx.Th(component.SpanText(t("Resource"))),
						nodx.Th(component.SpanText(t("Total"))),
						nodx.Th(component.SpanText(t("Healthy"))),
						nodx.Th(component.SpanText(t("Unhealthy"))),
					),
				),
				nodx.Tbody(
					nodx.Tr(
						nodx.Td(component.SpanText(t("Databases"))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", databasesQty.All))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", databasesQty.Healthy))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", databasesQty.Unhealthy))),
					),
					nodx.Tr(
						nodx.Td(component.SpanText(t("Destinations"))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", destinationsQty.All))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", destinationsQty.Healthy))),
						nodx.Td(component.SpanText(fmt.Sprintf("%d", destinationsQty.Unhealthy))),
					),
				),
			),
		},
	})

	return nodx.Div(
		nodx.Class("inline-block"),
		mo.HTML,
		nodx.Button(
			mo.OpenerAttr,
			nodx.Class("btn btn-ghost btn-neutral"),
			component.SpanText(t("Health status")),
			component.Ping(pingColor),
		),
	)
}
