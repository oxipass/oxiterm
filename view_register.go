package main

import (
	"github.com/oxipass/oxicrypt"
	"github.com/oxipass/oxilib"
	"github.com/rivo/tview"
	"log"
)

var registerForm *tview.Form
var newPassword string
var confirmPassword string

func GetRegisterScreen() (form *tview.Form) {
	if registerForm != nil {
		return registerForm
	}
	registerForm = tview.NewForm().
		AddPasswordField("New Password", "", 16, '*', NewPasswordChanged).
		AddPasswordField("Confirm Password", "", 16, '*', ConfirmPasswordChanged).
		AddButton("Create storage", func() {
			if CheckNewPassword() {
				oxiInstance := oxilib.GetInstance()

				err := oxiInstance.SetNewPassword(newPassword, oxicrypt.AES256Text)
				if err != nil {
					app.SetRoot(GetRegisterScreen(), true)
				}
				templatesErr := oxiInstance.AddDefaultItemsTemplates()
				if templatesErr != nil {
					log.Println(err.Error()) // Error happened but the app can work without templates
				}
				NavToMain(cViewDefault)

			} else {
				// E,pty password fields, set focus to new password
			}
		}).
		AddButton("Exit", func() {
			app.Stop()
		})
	return registerForm
}

func NewPasswordChanged(password string) {
	newPassword = password
}

func ConfirmPasswordChanged(password string) {
	confirmPassword = password
}

func CheckNewPassword() bool {
	if newPassword == "" {
		// show error, empty password
		return false
	}
	if newPassword != confirmPassword {
		// show error, passwords are different
		return false
	}
	return true
}
