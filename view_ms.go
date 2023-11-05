package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TODO: Implement adding new field

func GetMainScreen(activeView string) *tview.Flex {
	if wrapperFlex != nil {
		return wrapperFlex
	}
	var err error
	wrapperFlex = tview.NewFlex().SetDirection(tview.FlexRow)
	searchFlex := GetSearchFlex()
	mainFlex := tview.NewFlex()
	buttonsFlex := GetButtonsFlex()

	// Getting items list from the database
	itemsList, err = GetItemsList()
	if err != nil {
		NavToError("Items reading error: "+err.Error(), cScreenLogin)
		return nil
	}
	itemsList.SetInputCapture(processItemsEvents)
	itemsList.SetBorder(true).SetTitle(" Items (Ctrl+I) ").SetBorderPadding(1, 1, 1, 1)

	// Getting fields list from the database
	fieldsList = tview.NewList().ShowSecondaryText(true)
	fieldsList.SetInputCapture(processFieldsEvents)
	fieldsList.SetBorder(true).SetTitle(" Fields (F) ").SetBorderPadding(1, 1, 1, 1)

	// Joining items and fields in one horizontal flex
	mainFlex.AddItem(itemsList, 0, 1, true).
		AddItem(fieldsList, 0, 2, false)

	// Top level flex
	wrapperFlex.AddItem(mainFlex, 0, 1, true).
		AddItem(searchFlex, 1, 99, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonsFlex, 1, 0, false)

	wrapperFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			if NavToNext() {
				return nil
			}
		case tcell.KeyRight:
			if NavToNext() {
				return nil
			}
		case tcell.KeyLeft:
			if NavToPrevious() {
				return nil
			}
		case tcell.KeyCtrlF:
			mainSetFocus(searchInput)
			return nil
		}
		return event
	})

	mainFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			NavToNext()
			//if itemsList.HasFocus() {
			//	mainSetFocus(fieldsList)
			//} else if fieldsList.HasFocus() {
			//	mainSetFocus(buttonLock)
			//}
			return nil
		case tcell.KeyCtrlF:
			mainSetFocus(searchInput)
			return nil
		case tcell.KeyCtrlQ:
			actionStopApp()
			return nil
		}

		return event
	})

	return wrapperFlex
}

func ItemSelected() {
	switch itemsList.GetCurrentItem() {
	case 0:
		fieldsList.Clear()
		fieldsList.AddItem("🗄   User", "   "+"alexanrb", 0, FieldSelected)
		fieldsList.AddItem("🗄   Password", "   "+"Bsyu87z1", 0, FieldSelected)
		fieldsList.AddItem("🗄   PIN", "   "+"1799", 0, FieldSelected)
	case 1:
		fieldsList.Clear()
		fieldsList.AddItem("🗄 Paypal User", "", 0, FieldSelected)
		fieldsList.AddItem("🗄 Paypal Password", "", 0, FieldSelected)
	default:
		fieldsList.Clear()
		fieldsList.AddItem("🗄 User", "", 0, FieldSelected)
		fieldsList.AddItem("🗄 Password", "", 0, FieldSelected)
		fieldsList.AddItem("🗄 PIN", "", 0, FieldSelected)
	}

	//log.Println("Item selected")
}

func FieldSelected() {
	//log.Println("Field selected")
}
