package main

func actionLock() {
	_ = oxiInstance.Lock()
	NavToLogin()
}

func actionStopApp() {
	_ = oxiInstance.Close()
	app.Stop()
}
