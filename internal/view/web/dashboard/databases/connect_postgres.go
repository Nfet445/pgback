package databases

import (
	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func connectPostgresButton(reqCtx reqctx.Ctx) nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	htmxAttributes := func(url string) nodx.Node {
		return nodx.Group(
			htmx.HxPost(pathutil.BuildPath(url)),
			htmx.HxInclude("#connect-postgres-form"),
			htmx.HxDisabledELT(".connect-postgres-btn"),
			htmx.HxIndicator("#connect-postgres-loading"),
			htmx.HxValidate("true"),
		)
	}

	mo := component.Modal(component.ModalParams{
		Size:  component.SizeMd,
		Title: t("Connect PostgreSQL"),
		Content: []nodx.Node{
			nodx.FormEl(
				nodx.Id("connect-postgres-form"),
				nodx.Class("space-y-3"),

				component.InputControl(component.InputControlParams{
					Name:        "host",
					Label:       t("Host"),
					Placeholder: "localhost",
					Required:    true,
					Type:        component.InputTypeText,
				}),

				component.InputControl(component.InputControlParams{
					Name:        "port",
					Label:       t("Port"),
					Placeholder: "5432",
					Required:    false,
					Type:        component.InputTypeText,
					HelpText:    "Default: 5432",
				}),

				component.InputControl(component.InputControlParams{
					Name:        "user",
					Label:       t("Username"),
					Placeholder: "postgres",
					Required:    true,
					Type:        component.InputTypeText,
				}),

				component.InputControl(component.InputControlParams{
					Name:        "password",
					Label:       t("Password"),
					Placeholder: "",
					Required:    true,
					Type:        component.InputTypePassword,
				}),

				component.SelectControl(component.SelectControlParams{
					Name:     "ssl_mode",
					Label:    t("SSL mode"),
					Required: false,
					HelpText: "Default: prefer",
					Children: []nodx.Node{
						nodx.Option(nodx.Value("prefer"), nodx.Text("prefer")),
						nodx.Option(nodx.Value("disable"), nodx.Text("disable")),
						nodx.Option(nodx.Value("require"), nodx.Text("require")),
						nodx.Option(nodx.Value("verify-ca"), nodx.Text("verify-ca")),
						nodx.Option(nodx.Value("verify-full"), nodx.Text("verify-full")),
					},
				}),

				component.SelectControl(component.SelectControlParams{
					Name:        "version",
					Label:       t("Version"),
					Placeholder: t("Select version"),
					Required:    true,
					Children: []nodx.Node{
						nodx.Option(nodx.Value("13"), nodx.Text("13")),
						nodx.Option(nodx.Value("14"), nodx.Text("14")),
						nodx.Option(nodx.Value("15"), nodx.Text("15")),
						nodx.Option(nodx.Value("16"), nodx.Text("16")),
						nodx.Option(nodx.Value("17"), nodx.Text("17")),
						nodx.Option(nodx.Value("18"), nodx.Text("18")),
					},
				}),
			),

			nodx.Div(
				nodx.Class("flex justify-between items-center pt-4"),
				nodx.Div(),
				nodx.Div(
					nodx.Class("flex justify-end items-center space-x-2"),
					component.HxLoadingMd("connect-postgres-loading"),
					nodx.Button(
						htmxAttributes("/dashboard/databases/connect-postgres"),
						nodx.Class("connect-postgres-btn btn btn-primary"),
						nodx.Type("button"),
						component.SpanText(t("Connect and import")),
						lucide.Plug(),
					),
				),
			),
		},
	})

	button := nodx.Button(
		mo.OpenerAttr,
		nodx.Class("btn btn-secondary"),
		component.SpanText(t("Connect PostgreSQL")),
		lucide.Database(),
	)

	return nodx.Div(
		nodx.Class("inline-block"),
		mo.HTML,
		button,
	)
}
