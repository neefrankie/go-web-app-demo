package controllers

import (
	"github.com/labstack/echo"
	"gitlab.com/ftchinese/backyard/models"
	"gitlab.com/ftchinese/backyard/ui"
	"net/http"
)

func GetLogIn(c echo.Context) error {
	data := ui.BuildLoginUI(models.Login{})

	return c.Render(http.StatusOK, "login.html", data)
}

func PostLogIn(c echo.Context) error {
	l := models.Login{}

	if err := c.Bind(&l); err != nil {
		return err
	}

	if !l.Validate() {
		data := ui.BuildLoginUI(l)

		return c.Render(http.StatusOK, "login.html", data)
	}

	return c.JSON(http.StatusOK, l)
}
