package ui

import (
	"time"
)

// BaseUI describes the bare bone UI.
// It only has a top banner and footer.
type BaseUI struct {
	Year    int
	Version string
}

func NewBaseUI() BaseUI {
	return BaseUI{
		Year:    time.Now().Year(),
		Version: "0.0.1",
	}
}
