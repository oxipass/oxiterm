package main

// Navigation on the main screen

func NavToNext() bool {
	if itemsList.HasFocus() {
		mainSetFocus(fieldsList)
		return true
	} else if fieldsList.HasFocus() {
		mainSetFocus(tagsFilter)
		return true
	} else if tagsFilter.HasFocus() {
		mainSetFocus(searchInput)
		return true
	} else if searchInput.HasFocus() {
		mainSetFocus(buttonLock)
		return true
	} else if buttonLock.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonAddItem)
		return true
	} else if buttonAddItem.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonAddField)
		return true
	} else if buttonAddField.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonSettings)
		return true
	} else if buttonSettings.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonAbout)
		return true
	} else if buttonAbout.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(buttonQuit)
		return true
	} else if buttonQuit.HasFocus() {
		app.SetRoot(wrapperFlex, true).SetFocus(itemsList)
		return true
	}
	return false
}

func NavToPrevious() bool {
	if itemsList.HasFocus() {
		mainSetFocus(buttonQuit)
		return true
	} else if fieldsList.HasFocus() {
		mainSetFocus(itemsList)
		return true
	} else if tagsFilter.HasFocus() {
		mainSetFocus(fieldsList)
		return true
	} else if searchInput.HasFocus() {
		mainSetFocus(tagsFilter)
		return true
	} else if buttonLock.HasFocus() {
		mainSetFocus(searchInput)
		return true
	} else if buttonAddItem.HasFocus() {
		mainSetFocus(buttonLock)
		return true
	} else if buttonAddField.HasFocus() {
		mainSetFocus(buttonAddItem)
		return true
	} else if buttonSettings.HasFocus() {
		mainSetFocus(buttonAddField)
		return true
	} else if buttonAbout.HasFocus() {
		mainSetFocus(buttonSettings)
		return true
	} else if buttonQuit.HasFocus() {
		mainSetFocus(buttonAbout)
		return true
	}
	return false
}
