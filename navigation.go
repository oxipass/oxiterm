package main

func NavToScreen(screenName string) {
	switch screenName {
	case cScreenMain:
		NavToMain(cViewItems)
	case cScreenMenu:
		NavToMenu()
	case cScreenRegistration:
		NavToRegistration()
	case cScreenLogin:
		NavToLogin()
	}
}

func NavToRegistration() {
	app.SetRoot(GetRegisterScreen(), true)
}

func NavToMain(activeView string) {
	app.SetRoot(GetMainScreen(activeView), true)
}

func NavToLogin() {
	app.SetRoot(GetLoginScreen(), true)
}

func NavToAddField() {
	app.SetRoot(GetAddFieldScreen(), true)
}

func NavToAddItem() {
	app.SetRoot(GetAddItemScreen(), true)
}

func NavToError(errorMessage string, returnScreen string) {
	app.SetRoot(GetErrorView(errorMessage, returnScreen), true)
}

func NavToAbout(returnScreen string) {
	app.SetRoot(GetAboutView(returnScreen), true)
}

func NavToMenu() {
	app.SetRoot(GetMainMenu(), true)
}
