package main

import (
	"github.com/ShunyaNagashige/ca-tech-dojo-golang/config"
	"github.com/ShunyaNagashige/ca-tech-dojo-golang/model"
	"github.com/ShunyaNagashige/ca-tech-dojo-golang/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	model.Open()
}
