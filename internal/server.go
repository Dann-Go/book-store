package internal

import (
	"context"
	"fmt"
	delivery "github.com/Dann-Go/book-store/internal/book/delivery/http"
	"github.com/Dann-Go/book-store/internal/book/repository/postegres"
	"github.com/Dann-Go/book-store/internal/book/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	server *http.Server
}

type DbConfig struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string
}

func Inject() *gin.Engine  {
	cfg := DbConfig{
		Host: os.Getenv("host"),
		Port: os.Getenv("dbport"),
		Username: os.Getenv("username"),
		Password: os.Getenv("password"),
		DBName: os.Getenv("dbname"),
		SSLMode: os.Getenv("sslmode"),
	}
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf(err.Error())
	}


	router := gin.Default()
	bookRepo := postegres.NewPostgresqlRepository(db)
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