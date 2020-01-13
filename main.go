package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/ftchinese/backyard/controllers"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

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
		Year:    time.Now().Year(),
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
	e.Use(middleware.Recover())

	e.GET("/", controllers.GetHome)
	e.GET("/login", controllers.GetLogIn)
	e.POST("/login", controllers.PostLogIn)

	e.GET("/audio", controllers.AudioPage)

	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)

	e.GET("/users/:id", func(context echo.Context) error {
		id := context.Param("id")
		return context.String(http.StatusOK, id)
	})
	e.POST("/users", func(context echo.Context) error {
		u := new(User)
		if err := context.Bind(u); err != nil {
			return err
		}
		return context.JSON(http.StatusCreated, u)
	})

	// /show?team=x-men&member=wolverine
	e.GET("/show", func(context echo.Context) error {
		team := context.QueryParam("team")
		member := context.QueryParam("member")
		return context.String(http.StatusOK, "team: "+team+", member: "+member)
	})

	// application/x-www-form-urlencoded
	// /save
	// "name=Joe Smith", "email=joe@labstack.com"
	e.POST("/save", func(context echo.Context) error {
		name := context.FormValue("name")
		email := context.FormValue("email")
		return context.String(http.StatusOK, "name: "+name+", email:"+email)
	})

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, context echo.Context) (b bool, e error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	e.Logger.Fatal(e.Start(":1323"))
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
