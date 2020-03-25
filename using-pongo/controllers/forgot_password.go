package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/neefrankie/go-web-demo/using-pongo/models"
	"github.com/neefrankie/go-web-demo/using-pongo/views"
	"net/http"
)

// GetForgotPassword show a form to collection user's email
func GetForgotPassword(c echo.Context) error {
	ctx := views.NewCtxBuilder().
		WithForm(views.NewResetLetterForm(models.Identity{})).
		Build()
	return c.Render(http.StatusOK, "password_reset_email.html", ctx)
}

// PostForgotPassword handles sending email to help reset password.
func PostForgotPassword(c echo.Context) error {
	var i models.Identity
	if err := c.Bind(&i); err != nil {
		return err
	}

	i.Sanitize()

	if ok := i.Validate(); !ok {
		ctx := views.NewCtxBuilder().
			WithForm(views.NewResetLetterForm(i)).
			Build()

		return c.Render(http.StatusOK, "password_reset_email.html", ctx)
	}

	ctx := views.NewCtxBuilder().Set("done", true).Build()

	return c.Render(http.StatusOK, "password_reset_email.html", ctx)
}

func VerifyPasswordToken(c echo.Context) error {
	return nil
}

func GetResetPassword(c echo.Context) error {
	return nil
}

func PostResetPassword(c echo.Context) error {
	return nil
}
