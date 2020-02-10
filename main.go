package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gitlab.com/ftchinese/backyard/controllers"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

var (
	isProd  bool
	version string
	build   string
	config  Config
)

func init() {
	flag.BoolVar(&isProd, "production", false, "Indicate productions environment if present")
	var v = flag.Bool("v", false, "print current version")

	flag.Parse()

	if *v {
		fmt.Printf("%s\nBuild at %s\n", version, build)
		os.Exit(0)
	}

	config = Config{
		Debug:   !isProd,
		Version: version,
	}

	viper.SetConfigName("api")
	viper.AddConfigPath("$HOME/config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := echo.New()
	r, err := NewRenderer(config)
	if err != nil {
		panic(err)
	}

	e.Renderer = r

	if !isProd {
		e.Static("/", "build/dev")
	}

	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(mustGetSessionKey()))))
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())

	e.GET("/", controllers.GetHome)
	e.GET("/login", controllers.GetLogIn)
	e.POST("/login", controllers.PostLogIn)

	e.GET("/audio", controllers.AudioPage)

	e.Logger.Fatal(e.Start(":1323"))
}

func mustGetSessionKey() string {
	k := viper.GetString("web_app.superyard.echo_session")

	if k == "" {
		log.Fatal("Echo session key not found")
	}

	return k
}

// Recursively get all file paths in directory, including sub-directories.
func GetAllFilePathsInDirectory(dirPath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

// Recursively parse all files in directory, including sub-directories.
func ParseDirectory(dirPath string) (*template.Template, error) {
	paths, err := GetAllFilePathsInDirectory(dirPath)
	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}
