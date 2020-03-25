package views

import (
	"github.com/flosch/pongo2"
	"github.com/neefrankie/go-web-demo/using-pongo/widget"
)

// CtxBuilder builds template data.
type CtxBuilder struct {
	ctx pongo2.Context
}

func NewCtxBuilder() *CtxBuilder {
	return &CtxBuilder{
		ctx: make(pongo2.Context, 0),
	}
}

func (b *CtxBuilder) Set(name string, value interface{}) *CtxBuilder {
	b.ctx[name] = value
	return b
}

func (b *CtxBuilder) WithFlash(f *widget.Flash) *CtxBuilder {
	return b.Set("flash", f)
}

func (b *CtxBuilder) WithForm(form widget.Form) *CtxBuilder {
	return b.Set("form", form)
}

func (b *CtxBuilder) Build() pongo2.Context {
	return b.ctx
}
