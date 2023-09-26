package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TODO: In gui Ctrl+Q is not working as exit, find a way to exit
// TODO: Implement Lock command
// TODO: Implement adding new field

func GetMainScreen() *tview.Flex {
	var err error
	wrapperFlex = tview.NewFlex().SetDirection(tview.FlexRow)
	mainFlex := tview.NewFlex()
	buttonsFlex := GetButtonsFlex()

	itemsList, err = GetItemsList()
	if err != nil {
		app.SetRoot(GetErrorView("Items reading error: "+err.Error(), addItemForm), true)
		return nil
	}
	itemsList.SetBorder(true).SetTitle(" Items (Ctrl+I) ").SetBorderPadding(1, 1, 1, 1)

	fieldsList = tview.NewList().ShowSecondaryText(true)
	fieldsList.AddItem("🗄 User", "alexanrb", 0, FieldSelected)
	fieldsList.AddItem("🗄 Password", "Bsyu87z1", 0, FieldSelected)
	fieldsList.AddItem("🗄 PIN", "1799", 0, FieldSelected)
	fieldsList.SetBorder(true).SetTitle(" Fields (Ctrl+S) ").SetBorderPadding(1, 1, 1, 1)

	mainFlex.AddItem(itemsList, 0, 1, true).
		AddItem(fieldsList, 0, 2, false)

	wrapperFlex.AddItem(mainFlex, 0, 1, true).
		AddItem(buttonsFlex, 1, 0, false)

	itemsList.SetInputCapture(processItemsEvents)

	fieldsList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyLeft:
			app.SetRoot(wrapperFlex, true).SetFocus(itemsList)
			return nil
		case tcell.KeyRight:
			return nil
		}
		switch event.Rune() {
		case 'I':
			app.SetRoot(wrapperFlex, true).SetFocus(itemsList)
			return nil
		case 'i':
			app.SetRoot(wrapperFlex, true).SetFocus(itemsList)
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
			} else if buttonAddField.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(buttonExit)
				return nil
			}
		case tcell.KeyRight:
			if buttonAddItem.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(buttonAddField)
				return nil
			} else if buttonAddField.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(buttonExit)
				return nil
			}
		case tcell.KeyLeft:
			if buttonAddField.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(buttonAddItem)
				return nil
			} else if buttonExit.HasFocus() {
				app.SetRoot(wrapperFlex, true).SetFocus(buttonAddField)
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

func actionExit() {
	app.Stop()
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
