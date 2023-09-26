package main

import "github.com/rivo/tview"

func GetItemsList() (*tview.List, error) {
	itmsList := tview.NewList().ShowSecondaryText(false)

	items, err := oxiInstance.ReadAllItems(false, false)
	if err != nil {
		return nil, err
	}

	for _, objItem := range items {
		itmsList.AddItem(objItem.Name, "", 0, ItemSelected)
	}

	return itmsList, nil
}
