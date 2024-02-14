package main

import (
	"fmt"

	"github.com/heitorlimamorei/portifolio-go-api/config"
	"github.com/heitorlimamorei/portifolio-go-api/router"
)

func main() {
	err := config.Init()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	router.Initialize()
}
