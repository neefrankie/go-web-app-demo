package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetHome(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.Render(http.StatusOK, "home.html", nil)
}
