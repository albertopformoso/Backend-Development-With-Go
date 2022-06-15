package handler

import (
	"go-api/internal/middleware"
	"net/http"
)

// RoutePerson .
func RoutePerson(mux *http.ServeMux, sotrage Storage) {
	h := newPerson(sotrage)

	mux.HandleFunc("/api/v1/persons/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/api/v1/persons/update", h.update)
	mux.HandleFunc("/api/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/api/v1/persons/get-by-id", h.getByID)
	mux.HandleFunc("/api/v1/persons/get-all", h.getAll)
}

// RouteLogin .
func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)

	mux.HandleFunc("/api/v1/login", h.login)
}
