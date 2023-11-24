package main

import "log"

func GetItemsTemplates() []string {
	var templNames []string
	templates, err := oxi.GetTemplatesItems()
	if err != nil {
		log.Println(err.Error())
	}
	for _, templItem := range templates {
		templNames = append(templNames, templItem.Name)
	}
	templNames = append(templNames, "Empty (without template)")
	return templNames
}
