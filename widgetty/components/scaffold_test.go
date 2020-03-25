package components

import (
	"html/template"
	"testing"
)

func TestScaffold_Build(t *testing.T) {
	tmpl, err := template.ParseGlob("../templates/*.html")
	if err != nil {
		t.Error(err)
	}

	scaffold := NewScaffold(tmpl, "base.html").
		WithIcon("http://interactive.ftchinese.com/favicons").
		WithTitle("Go built-in template engine example")

	result, err := scaffold.Build()
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}
