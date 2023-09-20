package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/oxipass/oxilib"
	"github.com/rivo/tview"
)

// TODO: In gui Ctrl+Q is not working as exit, find a way to exit

var itemsList *tview.List
var fieldsList *tview.List
var buttonAddItem *tview.Button
var buttonAddField *tview.Button

func GetMainScreen() *tview.Flex {
	wrapperFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	mainFlex := tview.NewFlex()
	buttonsFlex := tview.NewFlex()

	itemsList = tview.NewList().ShowSecondaryText(false)

	oxilib.GetInstance()

	items, err := oxiInstance.ReadAllItems(false, false)
	if err != nil {
		app.SetRoot(GetErrorView("Items reading error: "+err.Error(), addItemForm), true)
	}

	for _, objItem := range items {
		itemsList.AddItem(objItem.Name, "", 0, ItemSelected)
	}

	itemsList.SetBorder(true).SetTitle(" Items (Ctrl+A) ").SetBorderPadding(1, 1, 1, 1)

	fieldsList = tview.NewList().ShowSecondaryText(true)
	fieldsList.AddItem("🗄 User", "alexanrb", 0, FieldSelected)
	fieldsList.AddItem("🗄 Password", "Bsyu87z1", 0, FieldSelected)
	fieldsList.AddItem("🗄 PIN", "1799", 0, FieldSelected)
	fieldsList.SetBorder(true).SetTitle(" Fields (Ctrl+S) ").SetBorderPadding(1, 1, 1, 1)

	buttonAddItem = tview.NewButton("Add Item (1)").SetSelectedFunc(addButtonPressed)
	buttonAddField = tview.NewButton("Add Field (2)").SetSelectedFunc(addButtonPressed)

	mainFlex.AddItem(itemsList, 0, 1, true).
		AddItem(fieldsList, 0, 2, false)

	buttonsFlex.AddItem(buttonAddItem, 15, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(buttonAddField, 15, 0, false)

	wrapperFlex.AddItem(mainFlex, 0, 1, true).
		AddItem(buttonsFlex, 1, 0, false)

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
		switch event.Rune() {
		case '1':
			app.SetRoot(wrapperFlex, true).SetFocus(buttonAddItem)
			return nil
		case '2':
			app.SetRoot(wrapperFlex, true).SetFocus(buttonAddField)
			return nil
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
			if buttonAddItem.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(buttonAddField)
				return nil
			}
		case tcell.KeyRight:
			if buttonAddItem.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(buttonAddField)
				return nil
			}
		case tcell.KeyLeft:
			if buttonAddField.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(buttonAddItem)
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
				app.SetRoot(wrapperFlex, true).SetFocus(buttonAddItem)
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
