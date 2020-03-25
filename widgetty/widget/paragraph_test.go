package widget

import (
	"testing"
)

func TestParagraph_Build(t *testing.T) {
	p := NewPara().
		WithContent("Hello world!")

	t.Log(p.Build())
}
