package main

import (
	"github.com/Dann-Go/book-store/internal"
	"log"
)

func main()  {

	server := new(internal.Server)
	if err := server.Run("8000"); err != nil {
		log.Fatalf("error while running server %s", err.Error())
	}
}
