package components

import "strings"

type BaseTag struct {
	Attr    Attributes
	Content string
}

type Tag struct {
	BaseTag
	Name        string
	SelfClosing bool
}

func (t *Tag) Build() string {
	var buf strings.Builder

	buf.WriteByte('<')
	buf.WriteString(t.Name)
	if t.Attr != nil {
		buf.WriteByte(' ')
		buf.WriteString(t.Attr.Encode())
	}

	if t.SelfClosing {
		buf.WriteString(" />")
		return buf.String()
	}

	buf.WriteByte('>')

	buf.WriteString(t.Content)
	buf.WriteString("</")
	buf.WriteString(t.Name)
	buf.WriteByte('>')

	return buf.String()
}
