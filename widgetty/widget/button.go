package widget

type BtnType string

const (
	BtnTypeButton = "button"
	BtnTypeSubmit = "submit"
)

type Button struct {
	BaseTag
}

func NewButton() *Button {
	return &Button{}
}

func (b *Button) WithAttr(attr Attributes) *Button {
	b.Attr = attr
	return b
}

func (b *Button) WithContent(s string) *Button {
	b.Content = s
	return b
}

func (b *Button) WithType(t BtnType) *Button {
	if b.Attr == nil {
		b.Attr = Attributes{}
	}
	b.Attr.Set("type", string(t))
	return b
}

func (b *Button) Build() string {
	tag := Tag{
		Name:        "button",
		BaseTag:     b.BaseTag,
		SelfClosing: false,
	}

	return tag.Build()
}
