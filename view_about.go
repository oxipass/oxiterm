package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var aboutView *tview.Modal

func GetAboutView(returnScreen string) (form *tview.Modal) {
	if aboutView != nil {
		return aboutView
	}
	aboutView = tview.NewModal().SetText(cAppName + "\n" +
		oxi.T("version") + cVersion + "\n" +
		cHomepage).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				NavToScreen(returnScreen)
			}
		})

	aboutView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlQ:
			// Exit the application
			actionStopApp()
			return nil
		case tcell.KeyEscape:
			NavToScreen(returnScreen)
			return nil
		}
		return event
	})

	return aboutView
}
