package main

import (
	"github.com/Dann-Go/book-store/internal"
	"log"
	"os"
)

func main()  {

	server := new(internal.Server)
	if err := server.Run(os.Getenv("servport")); err != nil {
		log.Fatalf("error while running server %s", err.Error())
	}
}
