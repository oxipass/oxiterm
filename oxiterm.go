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
var oxi *oxilib.StorageSingleton // Oxi Library active instance

// TODO: URGENT! Use oxilib multilingual localization

// TODO: Add interactive and non-interactive mode with flags
// TODO: Add the flag for configuration file
// TODO: Add the flag for storage folder (can maintain and test many databases)
// TODO: Add the flag to backup the data
// TODO: Add the flag to import the data (backup, json, NS Wallet, LastPass, 1Password)
// TODO: Add the flag to export the data
// TODO: Add the flag to generate human readable PDF file for printing
// TODO:

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
	oxi = oxilib.GetInstance()
	log.Println("OxiPass is initiating... ")
	err = oxi.Open(dbFile)
	if err != nil {
		log.Fatal("Initiation error: " + err.Error())
		return
	}

	isNew := oxi.IsNew()

	log.Println("App started in interactive mode")
	if isNew {
		log.Println("OxiPass first time start, performing secure database initialization")
		app.SetRoot(GetRegisterScreen(), true)
	} else {
		NavToLogin()
	}
	app.EnableMouse(true)
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Println("App stopped, bye")
}
