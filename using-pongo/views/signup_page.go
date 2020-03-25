package views

import (
	"github.com/neefrankie/go-web-demo/using-pongo/models"
	"github.com/neefrankie/go-web-demo/using-pongo/widget"
)

func NewSignUpForm(value models.SignUp) widget.Form {
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
				Value:       value.Email,
				Placeholder: "admin@example.org",
				Required:    true,
			},
			{
				Label:    "密码",
				ID:       "password",
				Type:     widget.ControlTypePassword,
				Name:     "password",
				Value:    "",
				Required: true,
			},
			{
				Label:    "确认密码",
				ID:       "confirmPassword",
				Type:     widget.ControlTypePassword,
				Name:     "confirmPassword",
				Value:    "",
				Required: true,
			},
		},
		SubmitBtn: widget.PrimaryBlockBtn.
			SetName("创建账号").
			SetDisabledText("正在创建..."),
		CancelBtn: widget.Link{},
		DeleteBtn: widget.Link{},
	}
}
