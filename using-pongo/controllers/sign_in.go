package controllers

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	admin "github.com/neefrankie/go-web-demo/using-pongo/models"
	"github.com/neefrankie/go-web-demo/using-pongo/views"
	"net/http"
)

func GetLogin(c echo.Context) error {
	ctx := views.NewCtxBuilder().
		WithForm(views.NewLoginForm(admin.Login{})).
		Build()

	return c.Render(http.StatusOK, "login.html", ctx)
}

func createSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get(sessionKey, c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: false,
	}

	return sess
}

func PostLogin(c echo.Context) error {
	var l admin.Login
	if err := c.Bind(&l); err != nil {
		return err
	}

	if ok := l.Sanitize().Validate(); !ok {
		ctx := views.NewCtxBuilder().
			WithForm(views.NewLoginForm(l)).
			Build()

		return c.Render(http.StatusOK, "login.html", ctx)
	}

	sess := createSession(c)
	sess.Values[loggedInKey] = l.Email
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, SiteMap.Home)
}

func LogOut(c echo.Context) error {
	sess, _ := session.Get(sessionKey, c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1, // A zero or negative number will expire the cookie immediately. If both Expires and Max-Age are set, Max-Age has precedence.
		HttpOnly: false,
	}
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, SiteMap.Login)
}
