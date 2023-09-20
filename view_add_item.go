package main

import (
	"github.com/oxipass/oxilib/models"
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
			var newItem models.UpdateItemForm
			newItem.Name = itemName
			newItem.Icon = "brands/golang"
			_, err := oxiInstance.AddNewItem(newItem)
			if err != nil {
				app.SetRoot(GetErrorView("Item adding error: "+err.Error(), addItemForm), true)
				return
			}
			app.SetRoot(GetMainScreen(), true)
		}).
		AddButton("Back", func() {
			app.SetRoot(GetMainScreen(), true)
		})
	return addItemForm
}

func ItemNameChangedChanged(iName string) {
	itemName = iName
}

func ItemIconChangedChanged(iIcon string) {
	itemIcon = iIcon
}
