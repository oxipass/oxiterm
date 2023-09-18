package main

type Config struct {
	DbFile     string `json:"db_file"`
	ShowMenu   bool   `json:"show_menu"`
	ConfigFile string
	Loaded     bool
}
