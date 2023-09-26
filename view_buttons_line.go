package main

import "github.com/rivo/tview"

func GetButtonsFlex() *tview.Flex {
	buttonsFlex := tview.NewFlex()

	buttonAddItem = tview.NewButton("Add Item (1)").SetSelectedFunc(addButtonPressed)
	buttonAddField = tview.NewButton("Add Field (2)").SetSelectedFunc(addButtonPressed)
	buttonExit = tview.NewButton("Exit (X)").SetSelectedFunc(actionExit)

	buttonsFlex.AddItem(buttonAddItem, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonAddField, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonExit, 15, 0, false)

	return buttonsFlex
}
