package main

import (
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

func TestTemplateToHTML(t *testing.T) {
	data := ui.Home{
		BaseUI: ui.NewBaseUI(),
		Inputs: []ui.Input{
			{
				Label:       "邮箱",
				ID:          "email",
				Type:        "email",
				Name:        "credentials[email]",
				Value:       "",
				Placeholder: "电子邮箱",
				MaxLength:   "64",
				Required:    true,
			},
			{
				Label:       "密码",
				ID:          "password",
				Type:        "password",
				Name:        "credentials[password]",
				Placeholder: "密码",
				MaxLength:   "64",
				Required:    true,
			},
		},
		PwResetLink: "/password-reset",
	}

	tmpl, err := ParseDirectory("./views")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Tempalete name %s", tmpl.Name())

	for _, v := range tmpl.Templates() {
		t.Logf("Template name %s", v.Name())
	}

}
