package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

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

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return fmt.Errorf("template %s not found", name)
	}
	return tmpl.Execute(w, data)
}

func NewTemplate(dirPath string) *Template {

	templates := make(map[string]*template.Template, 0)

	paths, err := GetAllFilePathsInDirectory(dirPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, p := range paths {
		templates[p] = template.Must(template.ParseFiles(p))
	}

	return &Template{
		templates: templates,
	}
}