package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var errorView *tview.Modal

func GetErrorView(errorMessage string, screenToReturn tview.Primitive) tview.Primitive {

	errorView = tview.NewModal().SetText(errorMessage).
		SetBackgroundColor(tcell.ColorRed).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				if screenToReturn != nil {
					app.SetRoot(screenToReturn, true)
				}
			}
		})

	return tview.Primitive(errorView)
}
