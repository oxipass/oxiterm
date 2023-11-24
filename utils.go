package main

import "os"

const cDefaultFolder = string(os.PathSeparator) + "OxiPass"
const cDefaultConfig = "config.json"
const cDefaultDbFile = "oxipass.db"

var config *Config

func GetConfig() (conf Config, err error) {
	if config != nil {
		return *config, nil
	}
	userConfigFolder, err := os.UserConfigDir()
	if err != nil {
		return conf, err
	}

	oxipassFullPath := userConfigFolder + cDefaultFolder

	config = new(Config)
	config.Loaded = false

	if _, err := os.Stat(oxipassFullPath); !os.IsNotExist(err) {
		// path/to/whatever exists
	}

	return *config, nil
}

func GetDefaultConfig() (conf *Config) {

	return conf
}
