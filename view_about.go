package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var aboutView *tview.Modal

func GetAboutView() (form *tview.Modal) {
	if aboutView != nil {
		return aboutView
	}
	aboutView = tview.NewModal().SetText(cAppName + "\n" + "Version 1.0").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(GetMainMenu(), true)
			}
		})

	aboutView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlQ:
			// Exit the application
			app.Stop()
			return nil
		}
		return event
	})

	return aboutView
}
