package main

import (
	"github.com/augustosang/gopportunities/configs"
	"github.com/augustosang/gopportunities/router"
)

var (
	logger *configs.Logger
)

func main() {
	logger = configs.GetLogger("main")

	// Initialize configs
	err := configs.Init()
	if err != nil {
		logger.Errf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
