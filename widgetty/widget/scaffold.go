package widget

import (
	"html/template"
)

type ScaffoldData struct {
	IconURL string
	Title   string
	NavBar  string
	Body    string
	Footer  string
}

type Scaffold struct {
	Setting
	iconURL string
	title   string
	navBar  Widget
	body    Widget
	footer  Widget
}

func NewScaffold(tmpl *template.Template, name string) *Scaffold {
	return &Scaffold{}
}

func (s *Scaffold) WithIcon(i string) *Scaffold {
	s.iconURL = i
	return s
}

func (s *Scaffold) WithNavBar(n Widget) *Scaffold {
	s.navBar = n
	return s
}

func (s *Scaffold) WithBody(b Widget) *Scaffold {
	s.body = b
	return s
}

func (s *Scaffold) WithFooter(f Widget) *Scaffold {
	s.footer = f
	return s
}

func (s *Scaffold) WithTitle(t string) *Scaffold {
	s.title = t
	return s
}

//func (s *Scaffold) Build() (string, error) {
//	var sb strings.Builder
//
//	data := ScaffoldData{
//		IconURL: s.iconURL,
//		Title:   s.title,
//	}
//
//	if s.navBar != nil {
//		navBar, err := s.navBar.Build()
//		if err != nil {
//			return "", err
//		}
//		data.NavBar = navBar
//	}
//
//	if s.body != nil {
//		body, err := s.body.Build()
//		if err != nil {
//			return "", err
//		}
//		data.Body = body
//	}
//
//	if s.footer != nil {
//		footer, err := s.footer.Build()
//		if err != nil {
//			return "", err
//		}
//		data.Footer = footer
//	}
//
//	err := s.tmpl.ExecuteTemplate(&sb, s.name, data)
//
//	if err != nil {
//		return "", err
//	}
//
//	return sb.String(), nil
//}
