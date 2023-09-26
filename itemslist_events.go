package main

import "github.com/gdamore/tcell/v2"

func processItemsEvents(event *tcell.EventKey) *tcell.EventKey {
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
	case 'f':
		app.SetRoot(wrapperFlex, true).SetFocus(fieldsList)
		return nil
	case 'F':
		app.SetRoot(wrapperFlex, true).SetFocus(fieldsList)
		return nil
	case 'x':
		app.Stop()
		return nil
	}
	return event
}
