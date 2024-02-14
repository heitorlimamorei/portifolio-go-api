package main

import (
	"github.com/heitorlimamorei/portifolio-go-api/config"
	"github.com/heitorlimamorei/portifolio-go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.InitDb()

	if err != nil {
		logger.ErrorF("config initialization error %v", err)
		panic(err)
	}

	router.Initialize()
}
