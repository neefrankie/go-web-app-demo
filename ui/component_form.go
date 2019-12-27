package ui

type SubmitButton struct {
	DisableWith string
	Text        string
}

type Form struct {
	Disabled  bool
	Action    string
	Inputs    []Input
	SubmitBtn SubmitButton
	CancelBtn Anchor
	DeleteBtn Anchor
}
