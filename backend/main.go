package main

import (
	"github.com/BurntSushi/toml"
	"github.com/SerhiiCho/reciper/backend/apiserver"
	"github.com/SerhiiCho/reciper/backend/utils"
)

func main() {
	config := apiserver.NewConfig()
	_, configErr := toml.DecodeFile("config/server.toml", config)
	utils.FatalIfError("toml decode error in main@main", configErr)

	if err := apiserver.Start(config); err != nil {
		utils.FatalIfError("error starting api server in main@main", err)
	}
}
