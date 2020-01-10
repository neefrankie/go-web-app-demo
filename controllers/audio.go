package controllers

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/ftchinese/backyard/ui"
	"net/http"
)

func AudioPage(c echo.Context) error {
	return c.Render(http.StatusOK, "interactive.html", ui.BuildAudioUI())
}
