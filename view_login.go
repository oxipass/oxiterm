package main

import (
	"github.com/oxipass/oxilib"
	"github.com/rivo/tview"
)

var loginForm *tview.Form
var actualPassword string

// FIXME: Hangs on the login with empty password

func GetLoginScreen() (form *tview.Form) {

	loginForm = tview.NewForm().
		AddPasswordField("Master Password", "", 16, '*', PasswordChanged).
		AddButton("Login", func() {
			if err := CheckPassword(); err == nil {
				app.SetRoot(GetMainScreen(), true)
			} else {
				app.SetRoot(GetErrorView(err.Error(), GetLoginScreen()), true)
			}
		}).
		AddButton("Main Menu", func() {
			app.SetRoot(GetMainMenu(), true)
		})
	return loginForm
}

func PasswordChanged(password string) {
	actualPassword = password
}

func CheckPassword() error {
	oxi := oxilib.GetInstance()
	return oxi.Unlock(actualPassword)
}
