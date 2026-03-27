package databases

import (
	"github.com/eduardolat/pgbackweb/internal/service/databases"
	"github.com/eduardolat/pgbackweb/internal/util/pathutil"
	"github.com/eduardolat/pgbackweb/internal/validate"
	"github.com/eduardolat/pgbackweb/internal/view/web/respondhtmx"
	"github.com/labstack/echo/v4"
)

type connectPostgresDTO struct {
	Host     string `form:"host" validate:"required"`
	Port     string `form:"port"`
	User     string `form:"user" validate:"required"`
	Password string `form:"password" validate:"required"`
	SSLMode  string `form:"ssl_mode"`
	Version  string `form:"version" validate:"required"`
}

func (h *handlers) connectPostgresHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var formData connectPostgresDTO
	if err := c.Bind(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}
	if err := validate.Struct(&formData); err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	importedCount, err := h.servs.DatabasesService.ImportPostgresDatabases(
		ctx, databases.ImportPostgresParams{
			Host:     formData.Host,
			Port:     formData.Port,
			User:     formData.User,
			Password: formData.Password,
			SSLMode:  formData.SSLMode,
			Version:  formData.Version,
		},
	)
	if err != nil {
		return respondhtmx.ToastError(c, err.Error())
	}

	_ = importedCount
	return respondhtmx.Redirect(c, pathutil.BuildPath("/dashboard/databases"))
}
