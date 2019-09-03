package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
)

const (
	master = `Names: {{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	overlay = `{{define "list"}} {{join . ","}}{{end}} `
)
func TestTemplate(t *testing.T)  {
	var (
		funcs = template.FuncMap{
			"join": strings.Join,
		}
		guadians = []string{
			"Gamora",
			"Goot",
			"Nebula",
			"Rocket",
			"Star-Load",
		}
	)

	masterTempl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}

	overlayTmpl, err := template.Must(masterTempl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}

	if err := masterTempl.Execute(os.Stdout, guadians); err != nil {
		log.Fatal(err)
	}

	if err := overlayTmpl.Execute(os.Stdout, guadians); err != nil {
		log.Fatal(err)
	}
}

func TestGlobTemplate(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	pattern := filepath.Join(dir, "./views/**/*.html")

	tmpl := template.Must(template.ParseGlob(pattern))

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template excecution: %s", err)
	}
}

func parseTemplateDir() (*template.Template, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	pattern := filepath.Join(dir, "./views/**/*.html")

	return template.ParseGlob(pattern)
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
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

	return template.ParseFiles(paths...)
}

type TemplateDate struct {
	SiteName string
	SiteURL string
}

func TestTemplateToHTML(t *testing.T)  {
	data := TemplateDate{
		SiteName: "Theory and Practice",
		SiteURL:  "http://www.ftchinese.com",
	}

	tmpl, err := ParseTemplateDir("views")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Tempalate name %s", tmpl.Name())

	for _, v := range tmpl.Templates()  {
		t.Logf("Template name %s", v.Name())
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "index.html", &data)
	if err != nil {
		t.Error(err)
	}
}

func TestGlob(t *testing.T)  {
	filenames, err := filepath.Glob("./views/**/*.html")
	if err != nil {
		t.Error(err)
	}

	t.Logf("File names: %+v", filenames)
}