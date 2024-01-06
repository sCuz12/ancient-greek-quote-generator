package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/sCuz12/ancient-greek-quote-api/models"
	"github.com/sCuz12/ancient-greek-quote-api/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Repository struct {
	DB *gorm.DB
}

var DB *gorm.DB

func main() {
	fmt.Println("hELLO WORLD");

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}

	//creates new router object
	router := chi.NewRouter()


	//---MIDLEWARES--/
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))


	//connect to database
	connectDatabase()

		//get counter service
		counterService := services.ConstructCounterService(DB)

		v1Router := chi.NewRouter()
		v1Router.Get("/healthz",handlerReadiness)
		v1Router.Get("/err",handlerErr)
	
		router.Mount("/v1",v1Router)
		v1Router.Get("/greek_quote", GetAllGreekQuotes)
	
		v1Router.Get("/random_greek_quote", func(w http.ResponseWriter, r *http.Request) {
			getRandomQuote(counterService, w, r)
		})
	

	//connect router to http server
	srv := &http.Server {
		Handler : router,
		Addr : ":" + portString,
	}

	log.Printf("Server Starting on port %v" , portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:",portString)
	
	
}

func connectDatabase() {
	//dbConnString :=  os.Getenv("SUPABASE_URL")
	dns := "host=db.bopvsnzxqntkwfxodesi.supabase.co user=postgres password=Malakopedo_123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	database , dbErr := gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	  }), &gorm.Config{})

	if dbErr != nil {
		panic("failed to connect database")
	}

	//database.AutoMigrate(&models.Quote{})
	 // Auto migrate your models
	database.AutoMigrate(&models.Counter{})
	DB = database

}