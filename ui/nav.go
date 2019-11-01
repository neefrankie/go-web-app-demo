package ui

type Anchor struct {
	Text string
	Href string
}

type NavItem struct {
	Anchor
	Active bool
}

var InteractiveNav = []NavItem{
	{
		Anchor: Anchor{
			Text: "Create Audio Article",
			Href: "",
		},
	},
	{
		Anchor: Anchor{
			Text: "Create Text-only Article",
			Href: "",
		},
	},
	{
		Anchor: Anchor{
			Text: "Create Speed Reading",
			Href: "",
		},
	},
}
