package main

import (
	"github.com/flosch/pongo2"
	"gitlab.com/ftchinese/backyard/models"
	"gitlab.com/ftchinese/backyard/ui"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
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

func TestViews(t *testing.T) {
	tmpl, err := ParseDirectory("./views")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", tmpl.DefinedTemplates())

	for _, v := range tmpl.Templates() {
		t.Logf("%+v\n", v.Name())

		err := v.Execute(os.Stdout, ui.NewBaseUI())

		if err != nil {
			t.Error(err)
		}
	}
}

func TestPongo(t *testing.T) {
	pongo2.DefaultLoader.SetBaseDir("templates")
	var tplExample = pongo2.Must(pongo2.FromFile("login.html"))

	err := tplExample.ExecuteWriter(pongo2.Context{
		"form": ui.Form{
			Disabled: false,
			Action:   "",
			Inputs:   ui.BuildLoginInputs(models.Login{}),
			SubmitBtn: ui.SubmitButton{
				DisableWith: "Logging in...",
				Text:        "Login",
			},
			CancelBtn: ui.Anchor{},
			DeleteBtn: ui.Anchor{},
		},
	}, os.Stdout)

	if err != nil {
		t.Error(err)
	}
}
