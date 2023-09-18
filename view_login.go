package main

import (
	"github.com/oxipass/oxilib"
	"github.com/rivo/tview"
)

var loginForm *tview.Form
var actualPassword string

func GetLoginScreen() (form *tview.Form) {

	loginForm = tview.NewForm().
		AddPasswordField("Master Password", "", 16, '*', PasswordChanged).
		AddButton("Login", func() {
			if CheckPassword() {
				app.SetRoot(GetMainScreen(), true)
			} else {

				app.SetRoot(GetErrorView("Password is wrong, try again", GetLoginScreen()), true)
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

func CheckPassword() bool {
	oxi := oxilib.GetInstance()

	//err := oxi.Open(dbFile)
	//if err != nil {
	//	return false
	//}

	err := oxi.Unlock(actualPassword)

	if err != nil {
		return false
	}
	return true
}
