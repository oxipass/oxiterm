package main

// TODO: finalize main screen navigation

func NavToMain(activeView string) {
	app.SetRoot(GetMainScreen(activeView), true)
}

func NavToLogin() {
	app.SetRoot(GetLoginScreen(), true)
}

func NavToNext() bool {
	if buttonAddItem.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonAddField)
		return true
	} else if buttonAddField.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonLock)
		return true
	} else if buttonLock.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonQuit)
		return true
	} else if buttonQuit.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(itemsList)
		return true
	}
	return false
}

func NavToPrevious() bool {
	if buttonAddField.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonAddItem)
		return true
	} else if buttonQuit.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonLock)
		return true
	} else if buttonLock.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonAddField)
		return true
	}
	return false
}
