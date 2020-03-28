package main

import "testing"

func TestNewTemplate(t *testing.T) {
	tpl := NewTemplate("templates")

	t.Logf("%v", tpl)
}