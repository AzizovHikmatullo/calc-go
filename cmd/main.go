package main

import (
	"github.com/AzizovHikmatullo/calc-go/internal/application"
	"log"
	"net/http"
)

func main() {
	app := application.New()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/calculate", application.CalculateHandler)

	log.Println("Starting server on port", app.Config.Addr)
	if err := http.ListenAndServe(":"+app.Config.Addr, mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
