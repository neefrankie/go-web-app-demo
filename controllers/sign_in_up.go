package controllers

import (
	"github.com/labstack/echo"
	"gitlab.com/ftchinese/backyard/ui"
	"net/http"
)

func GetLogIn(c echo.Context) error {
	data := ui.BuildLoginUI(ui.LoginFormState{})

	return c.Render(http.StatusOK, "login.html", data)
}

func PostLogIn(c echo.Context) error {

	return nil
}
