package main

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"path"
)

var Debugging bool

var templateNames = map[string][]string{
	"home.html":  {"layouts/base.html", "layouts/two-cols.html", "home.html"},
	"login.html": {"layouts/base.html", "layouts/center.html", "login.html"},
}

type TemplateSet struct {
	Debug         bool
	parser        Parser
	templateFiles map[string][]string
	templateCache map[string]*template.Template
}

func NewSet(debug bool) *TemplateSet {
	var parser Parser
	if debug {
		parser = NewFileParser("templates")
	} else {
		parser = NewRiceParser(rice.MustFindBox("templates"))
	}

	return &TemplateSet{
		Debug:         debug,
		parser:        parser,
		templateFiles: templateNames,
		templateCache: map[string]*template.Template{},
	}
}

func (set *TemplateSet) FromCache(name string) (*template.Template, error) {

	tmpl, ok := set.templateCache[name]
	if !ok {
		tmpl, err := set.FromRaw(name)
		if err != nil {
			return nil, err
		}

		set.templateCache[name] = tmpl
		return tmpl, nil
	}

	return tmpl, nil
}

func (set *TemplateSet) FromRaw(name string) (*template.Template, error) {
	files := set.templateFiles[name]
	return set.parser.Parse(name, files)
}

func (set *TemplateSet) GetTemplate(name string) (*template.Template, error) {
	if set.Debug {
		return set.FromRaw(name)
	}

	return set.FromCache(name)
}

func (set *TemplateSet) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t, err := set.GetTemplate(name)
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, name, data)
}

type Parser interface {
	Parse(name string, files []string) (*template.Template, error)
}

type FileParser struct {
	baseDir string
}

func NewFileParser(baseDir string) FileParser {
	return FileParser{baseDir: baseDir}
}

func (parser FileParser) Parse(name string, files []string) (*template.Template, error) {
	var paths []string
	for _, v := range files {
		paths = append(paths, path.Join(parser.baseDir, v))
	}

	return template.New(name).ParseFiles(paths...)
}

type RiceParser struct {
	box *rice.Box
}

func NewRiceParser(box *rice.Box) RiceParser {
	return RiceParser{box: box}
}

func (parser RiceParser) Parse(name string, files []string) (*template.Template, error) {

	var tplStr string
	for _, v := range files {
		s, err := parser.box.String(v)
		if err != nil {
			return nil, err
		}
		tplStr = tplStr + s
	}

	tmpl, err := template.New(name).Parse(tplStr)

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
