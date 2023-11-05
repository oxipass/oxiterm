package main

func addItemButtonPressed() {
	app.SetRoot(GetAddItemScreen(), true)
}

func addFieldButtonPressed() {
	app.SetRoot(GetAddFieldScreen(), true)
}

func aboutButtonPressed() {
	app.SetRoot(GetAboutView(cScreenMain), true)
}
