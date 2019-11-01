package ui

import (
	"time"
)

type BaseUI struct {
	SiteName    string
	SiteURL     string
	Sidebar     []NavItem
	Footer      []FooterColumn
	CurrentYear int
	Version     string
}

func NewUIBase(sidebar []NavItem) BaseUI {
	return BaseUI{
		SiteName:    "My Great Site",
		SiteURL:     "https://github.com/neefrankie",
		Sidebar:     sidebar,
		Footer:      Footer,
		CurrentYear: time.Now().Year(),
		Version:     "0.0.1",
	}
}

type Home struct {
	BaseUI
	Inputs      []Input
	PwResetLink string
}

func BuildHomeUI() Home {
	return Home{
		BaseUI: NewUIBase([]NavItem{}),
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

type AudioArticle struct {
	BaseUI
	Heading string
	Form    Form
}

func BuildAudioUI() AudioArticle {
	return AudioArticle{
		BaseUI:  NewUIBase(InteractiveNav),
		Heading: "New Audio Article",
		Form: Form{
			Disabled: false,
			Action:   "",
			Inputs: []Input{
				{
					Label:       "Title",
					ID:          "title",
					Type:        InputTypeText,
					Name:        "title",
					Value:       "",
					Placeholder: "The title of this article",
					Required:    true,
					Desc:        "Required",
				},
				{
					Label:       "Standfirst",
					ID:          "standfirst",
					Type:        InputTypeText,
					Name:        "standfirst",
					Value:       "",
					Placeholder: "Lead-in",
					Required:    true,
					Desc:        "Required",
				},
				{
					Label:       "Audio URL",
					ID:          "audioUrl",
					Type:        InputTypeURL,
					Name:        "audioUrl",
					Value:       "",
					Placeholder: "MP3 URL",
					Required:    true,
					Desc:        "Required",
				},
				{
					Label:       "Author",
					ID:          "author",
					Type:        InputTypeText,
					Name:        "author",
					Value:       "",
					Placeholder: "Who wrote this article?",
					Desc:        "Optional",
				},
				{
					Label: "Article Body",
					ID:    "body",
					Type:  InputTypeTextArea,
					Name:  "body",
					Value: "",
				},
			},
			SubmitBtn: SubmitButton{
				DisableWith: "Creating...",
				Text:        "Create",
			},
			CancelBtn: Anchor{
				Text: "Cancel",
				Href: "",
			},
		},
	}
}
