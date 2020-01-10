package main

import (
	"errors"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/ftchinese/backyard/controllers"
	"gitlab.com/ftchinese/backyard/ui"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

type Renderer struct {
	Debug bool
}

func (r Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var ctx = pongo2.Context{}

	if data != nil {
		var ok bool
		ctx, ok = data.(pongo2.Context)
		if !ok {
			return errors.New("no pongo2.Context data was passed")
		}
	}

	var t *pongo2.Template
	var err error

	if r.Debug {
		t, err = pongo2.FromFile(name)
	} else {
		t, err = pongo2.FromCache(name)
	}

	if err != nil {
		return err
	}

	ctx["env"] = ui.NewBaseUI()

	return t.ExecuteWriter(ctx, w)
}

func main() {
	if err := pongo2.DefaultLoader.SetBaseDir("templates"); err != nil {
		os.Exit(1)
	}

	e := echo.New()
	e.Renderer = Renderer{Debug: true}

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
