package main

import (
	rice "github.com/GeertJohan/go.rice"
	"io"
)

type RiceTemplateLoader struct {
	box *rice.Box
}

func NewRiceTemplateLoader(box *rice.Box) *RiceTemplateLoader {

	return &RiceTemplateLoader{box: box}
}

func (loader RiceTemplateLoader) Abs(base, name string) string {
	return name
}

func (loader RiceTemplateLoader) Get(path string) (io.Reader, error) {
	return loader.box.Open(path)
}
