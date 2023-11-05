package main

import "github.com/gdamore/tcell/v2"

func processItemsEvents(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyRight:
		mainSetFocus(fieldsList)
		return nil
	case tcell.KeyLeft:
		return nil
	case tcell.KeyUp:
	case tcell.KeyDown:
		ItemSelected()
	case tcell.KeyCtrlF:
		mainSetFocus(searchInput)
		return nil
	case tcell.KeyCtrlT:
		mainSetFocus(tagsFilter)
		return nil
	}
	switch event.Rune() {
	case '1':
		mainSetFocus(buttonAddItem)
		return nil
	case '2':
		mainSetFocus(buttonAddField)
		return nil
	case 'f':
		mainSetFocus(fieldsList)
		return nil
	case 'F':
		mainSetFocus(fieldsList)
		return nil
	}
	return event
}
