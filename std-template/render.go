package main

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

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
