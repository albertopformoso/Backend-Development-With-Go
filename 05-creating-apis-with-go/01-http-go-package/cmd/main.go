package main

import (
	"go-api/internal/authorization"
	"go-api/internal/handler"
	"go-api/internal/storage"
	"log"
	"net/http"
)

func main() {
	err := authorization.LoadFiles("cmd/certificates/app.rsa", "cmd/certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("couldn't load the certificates: %v\n", err)
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error on server: %v\n", err)
	}
}
