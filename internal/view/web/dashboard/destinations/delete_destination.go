package destinations

import (
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

func (h *handlers) deleteDestinationHandler(c echo.Context) error {
	ctx := c.Request().Context()

	destinationID, err := uuid.Parse(c.Param("destinationID"))
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	err = h.servs.DestinationsService.DeleteDestination(ctx, destinationID)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	return respondhtmx.Refresh(c)
}

func deleteDestinationButton(reqCtx reqctx.Ctx, destinationID uuid.UUID) nodx.Node {
	t := func(key string) string {
		if val, ok := i18n.Translations[reqCtx.Language][key]; ok {
			return val
		}
		return key
	}

	return component.OptionsDropdownButton(
		htmx.HxDelete(pathutil.BuildPath(fmt.Sprintf("/dashboard/destinations/%s", destinationID))),
		htmx.HxConfirm(t("Are you sure?")),
		lucide.Trash(),
		component.SpanText(t("Delete destination")),
	)
}
