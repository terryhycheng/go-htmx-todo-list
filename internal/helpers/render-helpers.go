package helpers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render()
func Render(c echo.Context, statusCode int, t templ.Component) error {
	c.Response().Writer.WriteHeader(statusCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(c.Request().Context(), c.Response().Writer)
}