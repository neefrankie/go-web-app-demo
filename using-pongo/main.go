package main

import (
	"flag"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neefrankie/go-web-demo/using-pongo/controllers"
)

var (
	isProd bool
	build  string
	config Config
)

func init() {
	flag.BoolVar(&isProd, "production", false, "Indicate productions environment if present")

	flag.Parse()

	config = Config{
		Debug: !isProd,
	}
}

func main() {
	e := echo.New()

	e.Renderer = MustNewRenderer(config)
	e.HTTPErrorHandler = errorHandler
	if !isProd {
		e.Static("/css", "client/node_modules/bootstrap/dist/css")
		e.Static("/js", "client/node_modules/bootstrap.native/dist")
	}

	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())

	e.GET("/",
		controllers.Home,
		controllers.RequireLoggedIn)
	e.GET("/login",
		controllers.GetLogin,
		controllers.RedirectIfLoggedIn)
	e.POST("/login", controllers.PostLogin)

	e.GET("/signup", controllers.GetSignUp)
	e.POST("/signup", controllers.PostSignUp)

	e.GET("/forgot-password", controllers.GetForgotPassword)
	e.POST("/forgot-password", controllers.PostForgotPassword)

	pwResetGroup := e.Group("/forgot-password")
	pwResetGroup.GET("/letter", controllers.GetForgotPassword)
	pwResetGroup.POST("/letter", controllers.PostForgotPassword)
	pwResetGroup.GET("/token/:token", controllers.VerifyPasswordToken)

	e.Logger.Fatal(e.Start(":1323"))
}
