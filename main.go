package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

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


	v1Router := chi.NewRouter()
	v1Router.Get("/healthz",handlerReadiness)
	v1Router.Get("/err",handlerErr)

	router.Mount("/v1",v1Router)

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