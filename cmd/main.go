package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"olx-clone-server/internal/config"
	"olx-clone-server/internal/handler"
	"olx-clone-server/internal/middleware"
	"olx-clone-server/internal/service"
	"os"
	"time"
)

func main() {
	c := config.Config{}
	if err := config.LoadConfig(os.Getenv("configPath"), "config", &c); err != nil {
		panic(err)
	}
	db, err := service.NewDB(&c)
	if err != nil {
		panic (err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/postings", handler.AllPostings(db)).Methods("GET")
	r.HandleFunc("/postings", handler.CreatePosting(db)).Methods("POST")
	r.HandleFunc("/auth/sign-in", handler.SignIn(db)).Methods("POST")
	r.HandleFunc("/auth/sign-up", handler.SignUp(db)).Methods("POST")

	r.Use(middleware.LoggerMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))


	// For dev only - Set up CORS so our client can consume the API
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	srv := &http.Server{
		Handler: corsWrapper.Handler(r),
		Addr: "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	fmt.Println("Listening at:", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}