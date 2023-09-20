package main

import (
	"fmt"
	"github.com/oxipass/oxilib"
	"github.com/rivo/tview"
	"log"
	"os"
	"os/user"
)

var app = tview.NewApplication()
var dbFile string
var oxiInstance *oxilib.StorageSingleton

func main() {
	fmt.Println(cAppName, cVersion)
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	appFolder := usr.HomeDir + "/" + cAppFolder
	dbFile = appFolder + "/" + cDbFile
	if _, err := os.Stat(appFolder); os.IsNotExist(err) {
		err := os.MkdirAll(appFolder, os.ModePerm)
		if err != nil {
			log.Fatal("Initiation error: " + err.Error())
			return
		}
	}
	oxiInstance = oxilib.GetInstance()
	log.Println("OxiPass is initiating... ")
	err = oxiInstance.Open(dbFile)
	if err != nil {
		log.Fatal("Initiation error: " + err.Error())
		return
	}

	isNew := oxiInstance.IsNew()

	log.Println("App started in interactive mode")
	if isNew {
		log.Println("OxiPass first time start, performing secure database initialization")
		app.SetRoot(GetRegisterScreen(), true)
	} else {
		app.SetRoot(GetLoginScreen(), true)
	}
	app.EnableMouse(true)
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Println("App stopped")
}

func StopApp() {

	_ = oxiInstance.Close()
	app.Stop()
}
