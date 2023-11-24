package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetSearchFlex() *tview.Flex {
	tags := GetTags()
	tags = append([]string{"------"}, tags...)
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
