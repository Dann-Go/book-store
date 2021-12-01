package internal

import (
	"context"
	delivery "github.com/Dann-Go/book-store/internal/book/delivery/http"
	"github.com/Dann-Go/book-store/internal/book/repository/postegres"
	"github.com/Dann-Go/book-store/internal/book/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func Inject() *gin.Engine  {
	router := gin.Default()
	bookRepo := postegres.NewPostgresqlRepository()
	bookUsecase := usecase.NewBookUsecase(bookRepo ,30)
	new(delivery.BookHandler).NewBookHandler(router.RouterGroup.Group("/books"), bookUsecase)
	return router

}
func (s *Server) Run (port string) error {
	router := Inject()

	s.server = &http.Server{
		Addr: ":" + port,
		Handler: router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("Server is running")
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown (ctx context.Context) error{
	return s.server.Shutdown(ctx)
}