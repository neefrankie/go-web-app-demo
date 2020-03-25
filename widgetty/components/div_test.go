package components

import "testing"

func TestDiv_Build(t *testing.T) {
	div := NewDiv().
		WithAttr(
			NewAttributes().
				Set("class", "mt-4"),
		).
		WithChildren(
			[]Widget{
				NewPara().WithContent("Paragraph 1"),
				NewPara().WithContent("Paragraph 2"),
				NewPara().WithContent("Paragraph 3"),
				NewPara().WithContent("Paragraph 4"),
				NewButton().
					WithContent("Submit").
					WithType(BtnTypeButton),
			},
		)

	t.Log(div.Build())
}
