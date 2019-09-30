package models

type DataList struct {
	ID      string
	Options []string
}

type TextInput struct {
	Label       string
	ID          string
	Type        string
	Name        string
	Value       string
	Placeholder string
	Required    bool
	ReadOnly    bool
	MinLength   string
	MaxLength   string
	Desc        string
	ErrMsg      string
	DataList    DataList
}
