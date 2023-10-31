package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TODO: Add buttons: About (0), Settings (S)

func GetSearchFlex() *tview.Flex {
	var tags []string
	tags = append(tags, "------------")
	tags = append(tags, "Internet")
	tags = append(tags, "Banking")
	tags = append(tags, "Crypto")
	tags = append(tags, "Phones")
	searchFlex := tview.NewFlex()
	searchInput = tview.NewInputField().SetLabel("Find (Ctrl+F): ").SetDoneFunc(startSearch)
	tagsFilter = tview.NewDropDown()
	tagsFilter.SetLabel("Filter by tag (Ctrl+T): ")
	tagsFilter.SetOptions(tags, tagSelected)
	tagsFilter.SetCurrentOption(0)
	searchFlex.AddItem(tview.NewBox(), 0, 2, false).
		AddItem(tagsFilter, 0, 47, false).
		AddItem(tview.NewBox(), 0, 2, false).
		AddItem(searchInput, 0, 47, false).
		AddItem(tview.NewBox(), 0, 2, false)
	return searchFlex
}

func tagSelected(option string, ind int) {

}

func startSearch(key tcell.Key) {

}

func GetButtonsFlex() *tview.Flex {
	buttonsFlex := tview.NewFlex()

	buttonAddItem = tview.NewButton("Add Item (1)").SetSelectedFunc(addItemButtonPressed)
	buttonAddField = tview.NewButton("Add Field (2)").SetSelectedFunc(addFieldButtonPressed)
	buttonLock = tview.NewButton("Lock (L)").SetSelectedFunc(actionLock)
	buttonQuit = tview.NewButton("Quit (Ctrl+Q)").SetSelectedFunc(actionStopApp)

	buttonsFlex.AddItem(buttonAddItem, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonAddField, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonLock, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonQuit, 15, 0, false)

	return buttonsFlex
}
