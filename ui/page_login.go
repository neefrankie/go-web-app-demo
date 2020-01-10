package ui

import (
	"github.com/flosch/pongo2"
	"gitlab.com/ftchinese/backyard/models"
)

func BuildLoginUI(formData models.Login) pongo2.Context {
	return pongo2.Context{
		"form": Form{
			Disabled: false,
			Action:   "",
			Inputs:   BuildLoginInputs(formData),
			SubmitBtn: SubmitButton{
				DisableWith: "Logging in...",
				Text:        "Login",
			},
			CancelBtn: Anchor{},
			DeleteBtn: Anchor{},
		},
	}
}

func BuildLoginInputs(formData models.Login) []Input {
	return []Input{
		{
			Label:       "User Name",
			ID:          "userName",
			Type:        InputTypeText,
			Name:        "userName",
			Value:       formData.UserName,
			Placeholder: "User name",
			MaxLength:   64,
			Required:    true,
			ErrMsg:      formData.Errors["UserName"],
		},
		{
			Label:       "Password",
			ID:          "password",
			Type:        InputTypePassword,
			Name:        "password",
			Placeholder: "Password",
			MaxLength:   64,
			Required:    true,
			ErrMsg:      formData.Errors["Password"],
		},
	}
}
