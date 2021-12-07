package internal

import (
	"fmt"
	delivery "github.com/Dann-Go/book-store/internal/book/delivery/http"
	"github.com/Dann-Go/book-store/internal/book/repository/postgres"
	"github.com/Dann-Go/book-store/internal/book/usecase"
	"github.com/Dann-Go/book-store/pkg/middleware"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Server struct {
	server *manners.GracefulServer
}

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func initLogger() {
	logger := log.New()
	logger.Out = os.Stdout
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func Inject() *gin.Engine {
	cfg := DbConfig{
		Host:     os.Getenv("Host"),
		Port:     os.Getenv("DBport"),
		Username: os.Getenv("Username"),
		Password: os.Getenv("Password"),
		DBName:   os.Getenv("DBname"),
		SSLMode:  os.Getenv("SSLmode"),
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

	query, err := ioutil.ReadFile("internal\\book\\repository\\postgres\\migrations\\migration.sql")
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err := db.Exec(string(query)); err != nil {
		log.Fatal(err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.Logger())
	bookRepo := postgres.NewPostgresqlRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo, 30)
	v := validator.New()

	new(delivery.BookHandler).NewBookHandler(router.RouterGroup.Group("/books"), bookUsecase, v)
	return router

}
func (s *Server) Run(port string) error {
	initLogger()
	router := Inject()
	s.server = manners.NewWithServer(&http.Server{
		Addr:           ":" + port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
	})
	log.Println("Server is running")
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() {
	log.Println("Shutting down")

	s.server.Close()
}
