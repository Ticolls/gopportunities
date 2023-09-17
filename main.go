package main

import (
	"github.com/Ticolls/gopportunities/config"
	"github.com/Ticolls/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {

	logger := *config.GetLogger("main")

	// Initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
