package components

import (
	"html/template"
	"testing"
)

func TestParagraph_Build(t *testing.T) {
	tmpl, err := template.ParseGlob("../templates/*.html")
	if err != nil {
		t.Error(err)
	}

	para := NewParagraph(tmpl, "paragraph.html").
		WithText("This is test paragraph")

	result, err := para.Build()
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}
