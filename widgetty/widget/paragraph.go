package widget

type Para struct {
	BaseTag
}

func NewPara() *Para {
	return &Para{}
}

func (p *Para) WithAttr(attr Attributes) *Para {
	p.Attr = attr
	return p
}

func (p *Para) WithContent(s string) *Para {
	p.Content = s
	return p
}

func (p *Para) Build() string {
	tag := Tag{
		Name:        "p",
		BaseTag:     p.BaseTag,
		SelfClosing: false,
	}
	return tag.Build()
}
