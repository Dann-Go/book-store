package main

import (
	"BookStore/internal"
	"BookStore/internal/book/delivery/http"
	"BookStore/internal/book/repository"
	"BookStore/internal/book/usecase"
	"log"
)

func main()  {
	bookRepo := repository.NewPostgresqlRepository()
	bookUsecase := usecase.NewBookUsecase(bookRepo ,30)
	handlers := new(http.BookHandler)

	server := new(internal.Server)
	if err := server.Run("8000", handlers.InitRoutes(bookUsecase)); err != nil {
		log.Fatalf("error while running server %s", err.Error())
	}
}
