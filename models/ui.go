package models

type Home struct {
	SiteName    string
	SiteURL     string
	Footer      []FooterColumn
	CurrentYear int
	Version     string
}
