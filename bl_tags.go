package main

import "github.com/oxipass/oxilib"

func GetTags() (tags []string) {
	oxiInstance := oxilib.GetInstance()
	tagsModels, err := oxiInstance.GetTags()
	if err != nil {
		return tags
	}
	for _, tagM := range tagsModels {
		tags = append(tags, tagM.Name)
	}
	return tags
}
