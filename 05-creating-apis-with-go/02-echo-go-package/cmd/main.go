package main

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"go-api/internal/authorization"
	"go-api/internal/handler"
	"go-api/internal/storage"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := authorization.LoadFiles("cmd/certificates/app.rsa", "cmd/certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("couldn't load the certificates: %v\n", err)
	}

	store := storage.NewMemory()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	apiVersion := e.Group("/api/v1")

	handler.RoutePerson(apiVersion, &store)
	handler.RouteLogin(apiVersion, &store)

	err = e.Start(":" + port)
	if err != nil {
		log.Printf("Error on server: %v\n", err)
	}
}
