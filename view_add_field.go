package main

import (
	"github.com/rivo/tview"
	"log"
)

// TODO: Implement coming back by pressing Esc

var addFieldForm *tview.Form
var fieldTemplate string
var fieldValue string
var fieldsTemplNames []string

func GetAddFieldScreen() (form *tview.Form) {

	fieldsTemplNamesStr, err := oxi.GetTemplatesItems()
	if err != nil {
		log.Println(err.Error())
	}
	for _, templItem := range fieldsTemplNamesStr {
		fieldsTemplNames = append(fieldsTemplNames, templItem.Name)
	}
	fieldsTemplNames = append(fieldsTemplNames, "Empty (without template)")
	addFieldForm = tview.NewForm().
		AddDropDown("Field template", fieldsTemplNames, 0, func(_ string, ind int) {
			FieldTemplateChanged(ind)
		}).
		AddInputField("Field value", "", 16, nil, FieldValueChanged).
		AddButton("Save Field", func() {
			NavToMain(cViewFields)
		}).
		AddButton("Back", func() {
			NavToMain(cViewFields)
		})
	return addFieldForm
}

func FieldTemplateChanged(ind int) {
	fieldTemplate = fieldsTemplNames[ind]
}

func FieldValueChanged(fValue string) {
	fieldValue = fValue
}
