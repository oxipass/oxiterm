package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TODO: In gui Ctrl+Q is not working as exit, find a way to exit

var itemsList *tview.List
var fieldsList *tview.List
var buttonAdd *tview.Button

func GetMainScreen() *tview.Flex {
	wrapperFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	mainFlex := tview.NewFlex()

	itemsList = tview.NewList().ShowSecondaryText(false)

	itemsList.AddItem("📁 Games", "", 0, ItemSelected)
	itemsList.AddItem("🗄 Paypal", "", 0, ItemSelected)
	itemsList.AddItem("🗄 Gmail", "", 0, ItemSelected)
	itemsList.AddItem("🗄 Bitcoin", "", 0, ItemSelected)
	itemsList.AddItem("🗄 Bitcoin", "", 0, ItemSelected)

	itemsList.SetBorder(true).SetTitle(" Items (Ctrl+A) ").SetBorderPadding(1, 1, 1, 1)

	fieldsList = tview.NewList().ShowSecondaryText(false)
	fieldsList.AddItem("🗄 User", "", 0, FieldSelected)
	fieldsList.AddItem("🗄 Password", "", 0, FieldSelected)
	fieldsList.AddItem("🗄 PIN", "", 0, FieldSelected)
	fieldsList.SetBorder(true).SetTitle(" Fields (Ctrl+S) ").SetBorderPadding(1, 1, 1, 1)

	buttonAdd = tview.NewButton("Add Item").SetSelectedFunc(addButtonPressed)

	mainFlex.AddItem(itemsList, 0, 1, true).
		AddItem(fieldsList, 0, 2, false)

	wrapperFlex.AddItem(mainFlex, 0, 1, true).
		AddItem(buttonAdd, 1, 0, false)

	itemsList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRight:
			app.SetRoot(wrapperFlex, true).SetFocus(fieldsList)
			return nil
		case tcell.KeyLeft:
			return nil
		case tcell.KeyUp:
		case tcell.KeyDown:
			ItemSelected()
		}
		return event
	})

	fieldsList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyLeft:
			app.SetRoot(wrapperFlex, true).SetFocus(itemsList)
			return nil
		case tcell.KeyRight:
			return nil
		}
		return event
	})

	wrapperFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			if buttonAdd.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(itemsList)
				return nil
			}
		case tcell.KeyCtrlF:
			app.SetRoot(wrapperFlex, true).SetFocus(fieldsList)
			return nil
		}
		return event
	})

	mainFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {
		case tcell.KeyTab:
			if itemsList.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(fieldsList)
			} else if fieldsList.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(buttonAdd)
			}
			return nil

		//case tcell.KeyCtrlI:
		//	app.SetRoot(flex, true).SetFocus(itemsList)
		case tcell.KeyCtrlF:
			app.SetRoot(wrapperFlex, true).SetFocus(fieldsList)
			return nil
		case tcell.KeyCtrlQ:
			// Exit the application
			app.Stop()
			return nil
		}
		return event
	})
	return wrapperFlex
}

func addButtonPressed() {
	app.SetRoot(GetAddItemScreen(), true)
}

func ItemSelected() {
	switch itemsList.GetCurrentItem() {
	case 0:
		fieldsList.Clear()
		fieldsList.AddItem("🗄 Games User", "", 0, FieldSelected)
		fieldsList.AddItem("🗄 Games Password", "", 0, FieldSelected)
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
