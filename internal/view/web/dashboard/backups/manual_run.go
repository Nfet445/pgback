package backups

import (
	"context"
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/i18n"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	nodx "github.com/nodxdev/nodxgo"
	htmx "github.com/nodxdev/nodxgo-htmx"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func (h *handlers) manualRunHandler(c echo.Context) error {
	backupID, err := uuid.Parse(c.Param("backupID"))
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	go func() {
		_ = h.servs.ExecutionsService.RunExecution(context.Background(), backupID)
	}()

	return respondhtmx.ToastSuccess(c, "Backup started, check the backup executions for more details")
}

func manualRunbutton(reqCtx reqctx.Ctx, backupID uuid.UUID) nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	return component.OptionsDropdownButton(
		htmx.HxPost(pathutil.BuildPath(fmt.Sprintf("/dashboard/backups/%s/run", backupID))),
		htmx.HxDisabledELT("this"),
		lucide.Zap(),
		component.SpanText(t("Run backup now")),
	)
}
