package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		logger: logger,
	}

	srv := &http.Server{
		Addr:         ":9090",
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Start server")
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
