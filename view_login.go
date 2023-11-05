package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var loginForm *tview.Form
var actualPassword string

// FIXME: Hangs on the login with empty password

func GetLoginScreen() (form *tview.Form) {

	loginForm = tview.NewForm().
		AddPasswordField("Master Password", "", 16, '*', PasswordChanged).
		AddButton("Login", func() {
			processCheckPassword()
		}).
		AddButton("Main Menu", func() {
			NavToMenu()
		}).
		AddButton("Quit (Ctrl+Q)", func() {
			actionStopApp()
		})
	loginForm.SetInputCapture(processLoginEvents)
	return loginForm
}

func processCheckPassword() {
	if err := CheckPassword(); err == nil {
		NavToMain(cViewDefault)
	} else {
		NavToError(err.Error(), cScreenLogin)
	}
}

func processLoginEvents(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlQ:
		actionStopApp()
		return nil
	}
	return event
}
func PasswordChanged(password string) {
	actualPassword = password
}
