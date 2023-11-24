package main

import (
	"github.com/rivo/tview"
)

// TODO: Add buttons: About (0), Settings (S)

func GetButtonsFlex() *tview.Flex {
	buttonsFlex := tview.NewFlex()

	buttonLock = tview.NewButton("Lock (L)").SetSelectedFunc(actionLock)
	buttonAddItem = tview.NewButton("Add Item (1)").SetSelectedFunc(addItemButtonPressed)
	buttonAddField = tview.NewButton("Add Field (2)").SetSelectedFunc(addFieldButtonPressed)
	buttonSettings = tview.NewButton("Settings (S)").SetSelectedFunc(addFieldButtonPressed)
	buttonAbout = tview.NewButton("About (A)").SetSelectedFunc(aboutButtonPressed)
	buttonQuit = tview.NewButton("Quit (Ctrl+Q)").SetSelectedFunc(actionStopApp)

	buttonsFlex.AddItem(buttonLock, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonAddItem, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonAddField, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonSettings, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonAbout, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonQuit, 15, 0, false)

	return buttonsFlex
}
