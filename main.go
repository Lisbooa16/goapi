package main

import (
	"github.com/Lisbooa16/goapi/config"
	"github.com/Lisbooa16/goapi/router"
)

var (
	logger config.Logger
)

func main() {
	// Initializer configs
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initializer error: %v", err)
		return
	}

	// Initializer router
	router.Initializer()
}
