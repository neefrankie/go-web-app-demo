package main

import (
	"errors"
	rice "github.com/GeertJohan/go.rice"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
)

// Renderer is used to render pong2 templates.
type Renderer struct {
	templateSet *pongo2.TemplateSet // Load templates from filesystem or rice.
	config      Config
}

// NewRenderer creates a new instance of Renderer based on runtime configuration.
func NewRenderer(config Config) (Renderer, error) {
	// In debug mode, we use pongo's default local file system loader.
	if config.Debug {
		log.Info("Development environment using local file system loader")

		loader := pongo2.MustNewLocalFileSystemLoader("templates")
		set := pongo2.NewSet("local", loader)
		set.Debug = true

		return Renderer{
			config:      config,
			templateSet: set,
		}, nil
	}

	log.Info("Production environment using rice template loader")
	box, err := rice.FindBox("templates")
	if err != nil {
		return Renderer{}, err
	}
	loader := NewRiceTemplateLoader(box)
	set := pongo2.NewSet("rice", loader)
	set.Debug = false

	return Renderer{
		config:      config,
		templateSet: set,
	}, nil
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

	if r.config.Debug {
		// In development the file is loaded from local
		// file system.
		t, err = r.templateSet.FromFile(name)
	} else {
		// In production the file is loaded from rice.
		t, err = r.templateSet.FromCache(name)
	}

	if err != nil {
		return err
	}

	ctx["env"] = r.config

	return t.ExecuteWriter(ctx, w)
}

type Config struct {
	Debug   bool
	Version string
	Year    int
}
