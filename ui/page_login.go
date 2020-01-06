package ui

import "gitlab.com/ftchinese/backyard/models"

type LoginUI struct {
	BaseUI
	Heading     string
	Form        Form
	PwResetLink string
}

func BuildLoginUI(formData models.Login) LoginUI {
	return LoginUI{
		BaseUI:  NewBaseUI(),
		Heading: "CMS Login",
		Form: Form{
			Disabled: false,
			Action:   "",
			Inputs:   buildLoginInputs(formData),
			SubmitBtn: SubmitButton{
				DisableWith: "Logging in...",
				Text:        "Login",
			},
		},
		PwResetLink: "/password-reset",
	}
}

func buildLoginInputs(formData models.Login) []Input {
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
