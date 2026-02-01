package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/k1941996/student-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student api"))
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server %s", err.Error())

	}
	fmt.Println("Server started")

	// connect to db
	// setup router
	// setup server
}
