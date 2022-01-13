package main

import (
	"github.com/Dann-Go/book-store/internal"
	log "github.com/sirupsen/logrus"
	"os"
)

// @title           Book-store
// @version         1.0
// @description     This is s simple app that is simulating a bookstore

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	server := new(internal.Server)
	if err := server.Run(os.Getenv("SERVPORT")); err != nil {
		log.Fatalf("error while running server %s", err.Error())

	}

}
