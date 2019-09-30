package models

import "time"

type UIBase struct {
	SiteName    string
	SiteURL     string
	Footer      []FooterColumn
	CurrentYear int
	Version     string
}

func NewUIBase() UIBase {
	return UIBase{
		SiteName:    "My Great Site",
		SiteURL:     "https://github.com/neefrankie",
		Footer:      Footer,
		CurrentYear: time.Now().Year(),
		Version:     "0.0.1",
	}
}

type Home struct {
	UIBase
	Inputs      []TextInput
	PwResetLink string
}
