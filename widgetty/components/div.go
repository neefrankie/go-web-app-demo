package components

import "strings"

type Div struct {
	BaseTag
	Children []Widget
}

func NewDiv() *Div {
	return &Div{
		Children: []Widget{},
	}
}
func (d *Div) WithAttr(attr Attributes) *Div {
	d.Attr = attr
	return d
}

func (d *Div) WithChildren(w []Widget) *Div {
	d.Children = w

	return d
}

func (d *Div) Build() string {
	var buf strings.Builder
	for _, v := range d.Children {
		buf.WriteString(v.Build())
	}

	tag := Tag{
		Name: "div",
		BaseTag: BaseTag{
			Attr:    d.Attr,
			Content: buf.String(),
		},
		SelfClosing: false,
	}

	return tag.Build()
}
