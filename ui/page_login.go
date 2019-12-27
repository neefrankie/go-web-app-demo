package ui

import "gitlab.com/ftchinese/backyard/models"

type LoginUI struct {
	BaseUI
	Heading     string
	Form        Form
	PwResetLink string
}

func BuildLoginUI(state LoginFormState) LoginUI {
	return LoginUI{
		BaseUI:  NewBaseUI(),
		Heading: "CMS Login",
		Form: Form{
			Disabled: false,
			Action:   "",
			Inputs:   buildLoginInputs(state),
			SubmitBtn: SubmitButton{
				DisableWith: "Logging in...",
				Text:        "Login",
			},
		},
		PwResetLink: "/password-reset",
	}
}

type LoginFormState struct {
	FormData models.Login
	Invalid  models.Login
}

func buildLoginInputs(state LoginFormState) []Input {
	return []Input{
		{
			Label:       "User Name",
			ID:          "userName",
			Type:        InputTypeText,
			Name:        "login[userName]",
			Value:       state.FormData.UserName,
			Placeholder: "User name",
			MaxLength:   64,
			Required:    true,
			ErrMsg:      state.Invalid.UserName,
		},
		{
			Label:       "Password",
			ID:          "password",
			Type:        InputTypePassword,
			Name:        "login[password]",
			Placeholder: "Password",
			MaxLength:   64,
			Required:    true,
			ErrMsg:      state.Invalid.Password,
		},
	}
}
