package main

import (
	"github.com/rivo/tview"
)

var registerForm *tview.Form
var newPassword string
var confirmPassword string

func GetRegisterScreen() (form *tview.Form) {
	//if registerForm != nil {
	//	return registerForm
	//}
	registerForm = tview.NewForm().
		AddPasswordField("New Password", "", 16, '*', NewPasswordChanged).
		AddPasswordField("Confirm Password", "", 16, '*', ConfirmPasswordChanged).
		AddButton("Create storage", func() {
			if CheckNewPassword() {
				err := SetNewPassword(newPassword)
				if err != nil {
					NavToError(err.Error(), cScreenRegistration)
				} else {
					NavToMain(cViewDefault)
				}

			}
		}).
		AddButton("Quit (Ctrl+Q)", func() {
			actionStopApp()
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
		NavToError("Password cannot be empty", cScreenRegistration)
		return false
	}
	if newPassword != confirmPassword {
		NavToError("Password confirmation failed, try again", cScreenRegistration)
		// show error, passwords are different
		return false
	}
	return true
}
