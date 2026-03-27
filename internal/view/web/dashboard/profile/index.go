package profile

import (
	"net/http"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/util/echoutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/layout"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
)

func (h *handlers) indexPageHandler(c echo.Context) error {
	ctx := c.Request().Context()
	reqCtx := reqctx.GetCtx(c)

	sessions, err := h.servs.AuthService.GetUserSessions(ctx, reqCtx.User.ID)
	if err != nil {
		logger.Error("failed to get user sessions", logger.KV{"err": err})
		return c.String(http.StatusInternalServerError, "failed to get user sessions")
	}

	return echoutil.RenderNodx(
		c, http.StatusOK, indexPage(reqCtx, sessions),
	)
}

func indexPage(reqCtx reqctx.Ctx, sessions []dbgen.Session) nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	content := []nodx.Node{
		component.H1Text(t("Profile")),

		nodx.Div(
			nodx.Class("mt-4 grid grid-cols-2 gap-4"),
			nodx.Div(updateUserForm(reqCtx, reqCtx.User)),
			nodx.Div(closeAllSessionsForm(reqCtx, sessions)),
		),
	}

	return layout.Dashboard(reqCtx, layout.DashboardParams{
		Title: t("Profile"),
		Body:  content,
	})
}
