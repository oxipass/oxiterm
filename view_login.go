package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var loginForm *tview.Form
var actualPassword string

// FIXME: Hangs on the login with empty password

func GetLoginScreen() *tview.Flex {
	loginVerticalFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	loginHorizontalFlex := tview.NewFlex().SetDirection(tview.FlexColumn)

	masterPasswordLabel := "Master Password"
	mpLabelLength := len(masterPasswordLabel)
	loginForm = tview.NewForm().
		//AddTextView("", cAppName+" "+cVersion, 56, 1, true, false).
		AddPasswordField(masterPasswordLabel, "", 58-mpLabelLength, '*', PasswordChanged).
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
	loginForm.SetBorder(true)
	loginForm.SetTitle(" " + cAppName + " " + cVersion + " ")
	footer := tview.NewTextView().SetText("Copyleft (c) OxiSoft 2022-2023").
		SetTextAlign(1)
	loginHorizontalFlex.AddItem(tview.NewBox(), 0, 50, false).
		AddItem(loginForm, 60, 0, true).
		AddItem(tview.NewBox(), 0, 50, false)
	loginVerticalFlex.AddItem(tview.NewBox(), 0, 50, false).
		AddItem(loginHorizontalFlex, 8, 0, true).
		AddItem(tview.NewBox(), 0, 50, false).
		AddItem(footer, 0, 50, false)
	return loginVerticalFlex
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
