package components

import (
	"html/template"
	"strings"
)

type Paragraph struct {
	Setting
	text string
}

func NewParagraph(tmpl *template.Template, name string) *Paragraph {
	return &Paragraph{Setting: NewSetting(tmpl, name)}
}

func (p *Paragraph) WithText(t string) *Paragraph {
	p.text = t
	return p
}

func (p *Paragraph) Build() (string, error) {
	var sb strings.Builder

	err := p.tmpl.ExecuteTemplate(&sb, p.name, p.text)

	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
