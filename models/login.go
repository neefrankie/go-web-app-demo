package models

import "strings"

type Login struct {
	UserName string `form:"userName"`
	Password string `form:"password"`
	Errors   map[string]string
}

func (l *Login) Validate() bool {
	l.UserName = strings.TrimSpace(l.UserName)
	l.Password = strings.TrimSpace(l.Password)

	l.Errors = make(map[string]string)

	if l.UserName == "" {
		l.Errors["UserName"] = "User name cannot be empty"
	}

	if l.Password == "" {
		l.Errors["Password"] = "Password cannot be empty"
	}

	if len(l.Password) < 6 {
		l.Errors["Password"] = "Password too short"
	}

	return len(l.Errors) == 0
}
