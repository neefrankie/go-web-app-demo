package controllers

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gitlab.com/ftchinese/backyard/models"
	"gitlab.com/ftchinese/backyard/ui"
	"net/http"
)

func GetLogIn(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["foo"] = "bar"
	sess.Save(c.Request(), c.Response())

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

	return c.Redirect(http.StatusFound, "/")
}
