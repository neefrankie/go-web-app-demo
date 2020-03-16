package components

import "html/template"

type Widget interface {
	Build() (string, error)
}

type Setting struct {
	tmpl *template.Template
	name string
}

func NewSetting(t *template.Template, n string) Setting {
	return Setting{
		tmpl: t,
		name: n,
	}
}
