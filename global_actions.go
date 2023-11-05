package main

func actionLock() {
	_ = oxi.Lock()
	NavToLogin()
}

func actionStopApp() {
	_ = oxi.Close()
	app.Stop()
}
