package main

import (
	"github.com/VladimirSh98/Gophermart.git/internal/app/logger"
	"github.com/VladimirSh98/Gophermart.git/internal/app/service"
	"log"
)

func main() {
	initLogger, err := logger.Initialize()
	defer initLogger.Sync()
	if err != nil {

		log.Fatalf("Logger configuration failed: %v", err)
	}
	err = service.Run()
	if err != nil {
		log.Fatalf("Service run failed: %v", err)
	}
}
