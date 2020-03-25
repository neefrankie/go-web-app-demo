package widget

import "strings"

// Attributes represents all the attributes of an HTML tag.
// It maps a string key to a list of values.
type Attributes map[string][]string

func NewAttributes() Attributes {
	return Attributes{}
}

// Set sets the name to value. It replaces any existing values.
// To set a boolean attribute, use empty string as value.
func (a Attributes) Set(name, value string) Attributes {
	value = strings.TrimSpace(value)

	if value == "" {
		a[name] = []string{}
		return a
	}

	a[name] = []string{value}
	return a
}

// Add adds the value to name. It appends to any existing
// values associated with the name
func (a Attributes) Add(name, value string) Attributes {
	if value == "" {
		return a
	}
	a[name] = append(a[name], value)

	return a
}

// Get gets the the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string.
func (a Attributes) Get(key string) string {
	if a == nil {
		return ""
	}
	vs := a[key]
	if len(vs) == 0 {
		return ""
	}

	return vs[0]
}

// Del deletes the values associated with name
func (a Attributes) Del(name string) Attributes {
	delete(a, name)

	return a
}

// Encode turns the values into tag attribute string.
// class="btn btn-primary" aria-label="hello"
func (a Attributes) Encode() string {
	if a == nil {
		return ""
	}

	var buf strings.Builder
	for n, v := range a {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(n)
		if len(v) == 0 {
			continue
		}
		buf.WriteByte('=')
		buf.WriteByte('"')
		buf.WriteString(strings.Join(v, " "))
		buf.WriteByte('"')
	}

	return buf.String()
}
