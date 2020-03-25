package components

import "testing"

func TestButton_Build(t *testing.T) {
	btn := NewButton().WithAttr(
		NewAttributes().
			Set("class", "btn btn-primary btn-block"),
	).
		WithContent("Login").
		WithType(BtnTypeButton)

	t.Logf("%s", btn.Build())
}
