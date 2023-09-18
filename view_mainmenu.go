package main

import "github.com/rivo/tview"

var mainMenu *tview.List

func GetMainMenu() (list *tview.List) {
	if mainMenu != nil {
		return mainMenu
	}

	mainMenu = tview.NewList().
		AddItem("Login", "Enter your master password to login", 'l', func() {
			app.SetRoot(GetLoginScreen(), true)
		}).
		AddItem("About", "About this app", 'a', func() {
			app.SetRoot(GetAboutView(), true)

		}).
		AddItem("Quit", "Leave the app", 'q', StopApp)
	return mainMenu
}
