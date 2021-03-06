package internal

import (
	"context"
	_ "context"
	_ "github.com/Dann-Go/book-store/docs"
	delivery "github.com/Dann-Go/book-store/internal/book/delivery/http"
	"github.com/Dann-Go/book-store/internal/book/repository/elastic_search"
	_ "github.com/Dann-Go/book-store/internal/book/repository/mongodb/indexes"
	_ "github.com/Dann-Go/book-store/internal/book/repository/postgres"
	"github.com/Dann-Go/book-store/internal/book/usecase"
	"github.com/Dann-Go/book-store/pkg/middleware"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/olivere/elastic/v7"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Server struct {
	server *manners.GracefulServer
}

type DbPostgresConfig struct {
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

func envsCheck() {
	requiredEnvs := []string{"HOST", "DBPORT", "USERNAME", "PASSWORD",
		"DBNAME", "SSLMODE", "SERVPORT", "MODE", "MONGOURI"}
	var msg []string
	for _, el := range requiredEnvs {
		val, exists := os.LookupEnv(el)
		if !exists || len(val) == 0 {
			msg = append(msg, el)
		}
	}
	if len(msg) > 0 {
		log.Fatal(strings.Join(msg, ", "), " env(s) not set")
	}
}

func Inject() *gin.Engine {
	//cfg := DbPostgresConfig{
	//	Host:     os.Getenv("HOST"),
	//	Port:     os.Getenv("DBPORT"),
	//	Username: os.Getenv("USERNAME"),
	//	Password: os.Getenv("PASSWORD"),
	//	DBName:   os.Getenv("DBNAME"),
	//	SSLMode:  os.Getenv("SSLMODE"),
	//}
	//connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	//db, err := sqlx.Open("postgres", connection)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//
	//err = db.Ping()
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//
	//query, err := ioutil.ReadFile("internal/book/repository/postgres/migrations/migration.sql")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//if _, err := db.Exec(string(query)); err != nil {
	//	log.Fatal(err.Error())
	//}

	if os.Getenv("MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
		//query, err := ioutil.ReadFile("internal/book/repository/postgres/migrations/seeds.sql")
		//if err != nil {
		//	log.Fatal(err.Error())
		//}
		//if _, err := db.Exec(string(query)); err != nil {
		//	log.Fatal(err.Error())
		//}
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	//mongoURI := os.Getenv("MONGOURI")

	//db, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	//if err != nil {
	//	log.Error(err)
	//}
	//err = db.Connect(context.TODO())
	//if err != nil {
	//	log.Error(err)
	//}
	//err = db.Ping(context.TODO(), nil)
	//if err != nil {
	//	log.Error(err)
	//}
	//
	//indexes.CreateIndex(db, "book-store", "books", "title", true)
	//indexes.CreateIndex(db, "book-store", "books", "authors", true)
	//indexes.CreateIndex(db, "book-store", "books", "year", true)
	//indexes.CreateIndex(db, "book-store", "books", "id", true)

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false), elastic.SetHealthcheck(false))
	if err != nil {
		log.Error(err)
	}
	mappings, err := ioutil.ReadFile("./internal/book/repository/elastic_search/config/elastic_cfg.json")
	if err != nil {
		log.Error(err)
	}
	exist, err := client.IndexExists("books").Do(context.Background())
	if err != nil {
		log.Error(err)
	} else if exist {
		log.Println("Index already exists")
		_, err = client.DeleteIndex("books").Do(context.Background())

	}
	if err != nil {
		log.Error(err)
	}

	_, err = client.CreateIndex("books").Body(string(mappings)).Do(context.Background())
	if err != nil {
		log.Error(err)
	}
	router := gin.New()
	metrics := middleware.NewPrometheusMiddleware("book_store", middleware.Opts{})
	private := router.Group("/api/books")
	private.Use(metrics.Metrics())
	prometheus.MustRegister(middleware.BOOKS_RESERVED)
	private.Use(middleware.Token_auth)
	public := router.Group("/")

	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	//bookRepo := postgres.NewPostgresqlRepository(db)
	//bookRepo := mongodb.NewMongoRepository(db)
	bookRepo := elastic_search.NewElasticRepository(client)
	bookUsecase := usecase.NewBookUsecase(bookRepo)

	new(delivery.BookHandler).NewBookHandler(private, bookUsecase)
	public.GET("/metrics", gin.WrapH(promhttp.Handler()))

	public.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "alive"})
	})
	public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router

}
func (s *Server) Run(port string) error {
	initLogger()
	envsCheck()
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
