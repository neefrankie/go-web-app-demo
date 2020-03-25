package components

import "html/template"

type Widget interface {
	Build() string
}

type Setting struct {
	tmpl *template.Template
	name string
}
