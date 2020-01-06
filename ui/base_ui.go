package ui

import (
	"time"
)

// BaseUI describes the bare bone UI.
// It only has a top banner and footer.
type BaseUI struct {
	SiteName    string
	SiteURL     string
	Footer      []FooterColumn
	CurrentYear int
	Version     string
}

func NewBaseUI() BaseUI {
	return BaseUI{
		SiteName:    "My Great Site",
		SiteURL:     "https://github.com/neefrankie",
		CurrentYear: time.Now().Year(),
		Version:     "0.0.1",
	}
}

// ContentBaseUI describes the layout of content UI.
// It has a sidebar in addition to top banner
// and footer.
type ContentBaseUI struct {
	BaseUI
	Sidebar []NavItem
}

func NewContentBaseUI(sidebar []NavItem) ContentBaseUI {
	return ContentBaseUI{
		BaseUI:  NewBaseUI(),
		Sidebar: sidebar,
	}
}

type Home struct {
	BaseUI
	Inputs      []Input
	PwResetLink string
}

func BuildHomeUI() Home {
	return Home{
		BaseUI: NewBaseUI(),
		Inputs: []Input{
			{
				Label:       "邮箱",
				ID:          "email",
				Type:        InputTypeEmail,
				Name:        "credentials[email]",
				Value:       "",
				Placeholder: "电子邮箱",
				MaxLength:   64,
				Required:    true,
			},
			{
				Label:       "密码",
				ID:          "password",
				Type:        InputTypePassword,
				Name:        "credentials[password]",
				Placeholder: "密码",
				MaxLength:   64,
				Required:    true,
			},
		},
		PwResetLink: "/password-reset",
	}
}
