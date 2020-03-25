package widget

import (
	"html/template"
	"os"
	"testing"
)

type Inventory struct {
	Material string
	Count    uint
}

func TestTemplate(t *testing.T) {
	sweaters := Inventory{
		Material: "wool",
		Count:    17,
	}

	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")

	if err != nil {
		t.Error(err)
	}

	err = tmpl.Execute(os.Stdout, sweaters)

	if err != nil {
		t.Error(err)
	}
}

const defineTest = `
defineTest starts

{{define "body"}}
body
{{end}}

defineTest ends
`

func TestDefine(t *testing.T) {
	tmpl, err := template.New("base").Parse(defineTest)

	if err != nil {
		t.Error(err)
	}

	// The output is body, base
	t.Logf("%+v\n", tmpl.DefinedTemplates())

	// The output is base.
	t.Logf("%s\n", tmpl.Name())

	// The output does not include the `define` part.
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		t.Error(err)
	}

	// The output only includes `define "body"` part.
	err = tmpl.ExecuteTemplate(os.Stdout, "body", nil)
	if err != nil {
		t.Error(err)
	}
}

const blockTest = `
**** container starts ****
{{block "body" .}}
This is body
{{end}}
**** container ends ****
`

func TestBlock(t *testing.T) {
	tmpl, err := template.New("base").Parse(blockTest)

	if err != nil {
		t.Error(err)
	}

	// The output is body, base
	t.Logf("%+v\n", tmpl.DefinedTemplates())

	// The output is base.
	t.Logf("%s\n", tmpl.Name())

	// The output includes the `block` part.
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		t.Error(err)
	}

	// The output only includes `block "body"` part.
	err = tmpl.ExecuteTemplate(os.Stdout, "body", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestOverride(t *testing.T) {

	tmpl, err := template.ParseFiles("center.html", "index.html", "layout.html")

	if err != nil {
		t.Error(err)
	}

	// The output is body, base
	t.Logf("%+v\n", tmpl.DefinedTemplates())

	// The output is base.
	t.Logf("%s\n", tmpl.Name())

	// layout.html
	if err = tmpl.ExecuteTemplate(os.Stdout, "layout.html", nil); err != nil {
		t.Error(err)
	}

	if err = tmpl.ExecuteTemplate(os.Stdout, "index.html", nil); err != nil {
		t.Error(err)
	}

}
