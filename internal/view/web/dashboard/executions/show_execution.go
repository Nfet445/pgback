package executions

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) downloadExecutionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	executionID, err := uuid.Parse(c.Param("executionID"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	isLocal, link, err := h.servs.ExecutionsService.GetExecutionDownloadLinkOrPath(
		ctx, executionID,
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if isLocal {
		return c.Attachment(link, filepath.Base(link))
	}

	return c.Redirect(http.StatusFound, link)
}

func showExecutionButton(
	reqCtx reqctx.Ctx,
	execution dbgen.ExecutionsServicePaginateExecutionsRow,
) nodx.Node {
	t := func(key string) string { if val, ok := i18n.Translations[reqCtx.Language][key]; ok { return val }; return key }

	mo := component.Modal(component.ModalParams{
		Title: t("Details"),
		Size:  component.SizeMd,
		Content: []nodx.Node{
			nodx.Div(
				nodx.Class("overflow-x-auto"),
				nodx.Table(
					nodx.Class("table [&_th]:text-nowrap"),
					nodx.Tr(
						nodx.Th(component.SpanText(t("ID"))),
						nodx.Td(component.SpanText(execution.ID.String())),
					),
					nodx.Tr(
						nodx.Th(component.SpanText(t("Status"))),
						nodx.Td(component.StatusBadge(execution.Status)),
					),
					nodx.Tr(
						nodx.Th(component.SpanText(t("Database"))),
						nodx.Td(component.SpanText(execution.DatabaseName)),
					),
					nodx.Tr(
						nodx.Th(component.SpanText(t("Destination"))),
						nodx.Td(component.PrettyDestinationName(
							execution.BackupIsLocal, execution.DestinationName,
						)),
					),
					nodx.If(
						execution.Message.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText(t("Message"))),
							nodx.Td(
								nodx.Class("break-all"),
								component.SpanText(execution.Message.String),
							),
						),
					),
					nodx.Tr(
						nodx.Th(component.SpanText(t("Started at"))),
						nodx.Td(component.SpanText(
							execution.StartedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
						)),
					),
					nodx.If(
						execution.FinishedAt.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText(t("Finished at"))),
							nodx.Td(component.SpanText(
								execution.FinishedAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
							)),
						),
					),
					nodx.If(
						execution.FinishedAt.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText(t("Took"))),
							nodx.Td(component.SpanText(
								execution.FinishedAt.Time.Sub(execution.StartedAt).String(),
							)),
						),
					),
					nodx.If(
						execution.DeletedAt.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText(t("Deleted at"))),
							nodx.Td(component.SpanText(
								execution.DeletedAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
							)),
						),
					),
					nodx.If(
						execution.FileSize.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText(t("File size"))),
							nodx.Td(component.PrettyFileSize(execution.FileSize)),
						),
					),
				),
				nodx.If(
					execution.Status == "success",
						nodx.Div(
							nodx.Class("flex justify-end items-center space-x-2"),
							deleteExecutionButton(reqCtx, execution.ID),
							nodx.A(
								nodx.Href(pathutil.BuildPath(fmt.Sprintf("/dashboard/executions/%s/download", execution.ID))),
								nodx.Target("_blank"),
								nodx.Class("btn btn-primary"),
								component.SpanText(t("Download")),
								lucide.Download(),
							),
						),
				),
			),
		},
	})

	return nodx.Div(
		mo.HTML,
		component.OptionsDropdownButton(
			mo.OpenerAttr,
			lucide.Eye(),
			component.SpanText(t("Show details")),
		),
	)
}
