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

func TestDirectory(t *testing.T) {
	paths, err := GetAllFilePathsInDirectory("./")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Current paths: %v", paths)
}
