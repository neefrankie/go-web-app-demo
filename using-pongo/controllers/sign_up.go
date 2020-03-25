package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/neefrankie/go-web-demo/using-pongo/models"
	"github.com/neefrankie/go-web-demo/using-pongo/views"
	"net/http"
)

func GetSignUp(c echo.Context) error {
	ctx := views.NewCtxBuilder().
		WithForm(views.NewSignUpForm(models.SignUp{})).
		Build()

	return c.Render(http.StatusOK, "signup.html", ctx)
}

func PostSignUp(c echo.Context) error {
	var s models.SignUp
	if err := c.Bind(&s); err != nil {
		return err
	}

	if ok := s.Sanitize().Validate(); !ok {
		ctx := views.NewCtxBuilder().
			WithForm(views.NewSignUpForm(s)).
			Build()

		return c.Render(http.StatusOK, "signup.html", ctx)
	}

	// TODO: Save signup data to db; retrieve the account.

	sess := createSession(c)
	sess.Values[loggedInKey] = s.Email
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, SiteMap.Home)
}
