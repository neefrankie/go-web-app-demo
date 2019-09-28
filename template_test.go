package main

import (
	"gitlab.com/ftchinese/backyard/models"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
	"time"
)

const (
	master  = `Names: {{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	overlay = `{{define "list"}} {{join . ","}}{{end}} `
)

func TestFilePathGlob(t *testing.T) {
	filenames, err := filepath.Glob("./views/**/*.html")
	if err != nil {
		t.Error(err)
	}

	t.Logf("File names: %+v", filenames)
}

func TestTemplate(t *testing.T) {
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

func TestDirectory(t *testing.T) {
	paths, err := GetAllFilePathsInDirectory("./")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Current paths: %v", paths)
}

type TemplateData struct {
	SiteName string
	SiteURL  string
}

func TestTemplateToHTML(t *testing.T) {
	data := models.Home{
		SiteName:    "My Great Site",
		SiteURL:     "https://github.com/neefrankie",
		Footer:      models.Footer,
		CurrentYear: time.Now().Year(),
		Version:     "0.0.1",
	}

	tmpl, err := ParseDirectory("./views")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Tempalete name %s", tmpl.Name())

	for _, v := range tmpl.Templates() {
		t.Logf("Template name %s", v.Name())
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "index.html", &data)
	if err != nil {
		t.Error(err)
	}
}
