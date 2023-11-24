package main

import (
	"github.com/oxipass/oxilib/models"
	"github.com/rivo/tview"
)

// TODO: Implement coming back by pressing Esc
// TODO: Check that item name is not empty
// TODO: Assign tags button
// TODO: Unassign tags button
// TODO: Show assigned items tags

var addItemForm *tview.Form
var itemName string
var itemIcon string
var itemTemplate string

func GetAddItemScreen() (form *tview.Form) {
	templatesName := GetItemsTemplates()
	addItemForm = tview.NewForm().
		AddInputField("Item name", "", 16, nil, ItemNameChangedChanged).
		AddInputField("Item icon", "", 16, nil, ItemIconChangedChanged).
		AddDropDown("Item template", templatesName, 0, func(option string, ind int) {

		}).
		AddButton("Save Item", func() {
			var newItem models.UpdateItemForm
			newItem.Name = itemName
			newItem.Icon = "brands/golang"
			_, err := oxi.AddNewItem(newItem)
			if err != nil {
				NavToError("Item adding error: "+err.Error(), cScreenAddItem)
				return
			}
			wrapperFlex = nil // Forcing screen reload
			NavToMain(cViewItems)
		}).
		AddButton("Back", func() {
			NavToMain(cScreenMain)
		})
	return addItemForm
}

func ItemNameChangedChanged(iName string) {
	itemName = iName
}

func ItemIconChangedChanged(iIcon string) {
	itemIcon = iIcon
}
