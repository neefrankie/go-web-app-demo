package views

import (
	"github.com/neefrankie/go-web-demo/using-pongo/models"
	"github.com/neefrankie/go-web-demo/using-pongo/widget"
)

func NewResetLetterForm(i models.Identity) widget.Form {
	return widget.Form{
		Disabled: false,
		Method:   widget.MethodPost,
		Action:   "",
		Fields: []widget.FormControl{
			{
				Label:       "邮箱",
				ID:          "email",
				Type:        widget.ControlTypeEmail,
				Name:        "email",
				Value:       i.Email,
				Placeholder: "admin@example.org",
				Required:    true,
			},
		},
		SubmitBtn: widget.PrimaryBlockBtn.
			SetName("发送邮件").
			SetDisabledText("正在发送..."),
		CancelBtn: widget.Link{},
		DeleteBtn: widget.Link{},
	}
}
