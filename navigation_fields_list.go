package main

import "github.com/gdamore/tcell/v2"

func processFieldsEvents(event *tcell.EventKey) *tcell.EventKey {
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
}
