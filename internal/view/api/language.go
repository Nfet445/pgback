package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) setLanguageHandler(c echo.Context) error {
	lang := c.Param("lang")
	
	cookie := &http.Cookie{
		Name:     "lang",
		Value:    lang,
		Path:     "/",
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   365 * 24 * 60 * 60, // 1 year
	}
	c.SetCookie(cookie)
	
	return c.NoContent(http.StatusOK)
}
