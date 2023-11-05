package main

import (
	"github.com/oxipass/oxicrypt"
	"github.com/oxipass/oxilib"
	"log"
)

func CheckPassword() error {
	oxi := oxilib.GetInstance()
	return oxi.Unlock(actualPassword)
}

func SetNewPassword(newPassword string) error {
	oxiInstance := oxilib.GetInstance()

	err := oxiInstance.SetNewPassword(newPassword, oxicrypt.AES256Text)
	if err != nil {
		log.Println(err.Error()) // Critical error, initial password is not set
		return err
	}
	templatesErr := oxiInstance.AddDefaultItemsTemplates()

	if templatesErr != nil {
		log.Println(templatesErr.Error()) // Error happened but the app can work without templates
	}
	return nil
}
