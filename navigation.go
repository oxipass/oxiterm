package main

// TODO: finalize main screen navigation

func NavToScreen(screenName string) {
	switch screenName {
	case cScreenMain:
		NavToMain(cViewItems)
	case cScreenMenu:
		NavToMenu()
	case cScreenLogin:
		NavToLogin()
	}
}

func NavToMain(activeView string) {
	app.SetRoot(GetMainScreen(activeView), true)
}

func NavToLogin() {
	app.SetRoot(GetLoginScreen(), true)
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