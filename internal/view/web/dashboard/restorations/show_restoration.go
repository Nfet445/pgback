package restorations

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func showRestorationButton(
	reqCtx reqctx.Ctx,
	restoration dbgen.RestorationsServicePaginateRestorationsRow,
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
						nodx.Td(component.SpanText(restoration.ID.String())),
					),
					nodx.Tr(
						nodx.Th(component.SpanText(t("Status"))),
						nodx.Td(component.StatusBadge(restoration.Status)),
					),
					nodx.Tr(
						nodx.Th(component.SpanText(t("Backup"))),
						nodx.Td(component.SpanText(restoration.BackupName)),
					),
					nodx.Tr(
						nodx.Th(component.SpanText(t("Database"))),
						nodx.Td(component.SpanText(func() string {
							if restoration.DatabaseName.Valid {
								return restoration.DatabaseName.String
							}
							return t("Other database")
						}())),
					),
					nodx.If(
						restoration.Message.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText(t("Message"))),
							nodx.Td(
								nodx.Class("break-all"),
								component.SpanText(restoration.Message.String),
							),
						),
					),
					nodx.Tr(
						nodx.Th(component.SpanText(t("Started at"))),
						nodx.Td(component.SpanText(
							restoration.StartedAt.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
						)),
					),
					nodx.If(
						restoration.FinishedAt.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText(t("Finished at"))),
							nodx.Td(component.SpanText(
								restoration.FinishedAt.Time.Local().Format(timeutil.LayoutYYYYMMDDHHMMSSPretty),
							)),
						),
					),
					nodx.If(
						restoration.FinishedAt.Valid,
						nodx.Tr(
							nodx.Th(component.SpanText(t("Took"))),
							nodx.Td(component.SpanText(
								restoration.FinishedAt.Time.Sub(restoration.StartedAt).String(),
							)),
						),
					),
				),
			),
		},
	})

	button := nodx.Button(
		mo.OpenerAttr,
		nodx.Class("btn btn-square btn-sm btn-ghost"),
		lucide.Eye(),
	)

	return nodx.Div(
		nodx.Class("inline-block tooltip tooltip-right"),
		nodx.Data("tip", t("Show details")),
		mo.HTML,
		button,
	)
}
