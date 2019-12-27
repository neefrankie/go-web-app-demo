package ui

type AudioArticle struct {
	BaseUI
	Heading string
	Form    Form
}

func BuildAudioUI() AudioArticle {
	return AudioArticle{
		BaseUI:  NewBaseUI(),
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
