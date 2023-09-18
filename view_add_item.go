package main

import (
	"github.com/rivo/tview"
)

var addItemForm *tview.Form
var itemName string
var itemIcon string

func GetAddItemScreen() (form *tview.Form) {

	addItemForm = tview.NewForm().
		AddInputField("Item name", "", 16, nil, ItemNameChangedChanged).
		AddInputField("Item icon", "", 16, nil, ItemIconChangedChanged).
		AddButton("Add", func() {

		}).
		AddButton("Back", func() {
			app.SetRoot(GetMainScreen(), true)
		})
	return loginForm
}

func ItemNameChangedChanged(iName string) {
	itemName = iName
}

func ItemIconChangedChanged(iIcon string) {
	itemIcon = iIcon
}
