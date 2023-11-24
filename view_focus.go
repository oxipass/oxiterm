package main

import "github.com/rivo/tview"

func mainSetFocus(prim tview.Primitive) {
	app.SetRoot(wrapperFlex, true).SetFocus(prim)
}
