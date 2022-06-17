package handler

import (
	"go-api/internal/middleware"

	"github.com/labstack/echo"
)

// RoutePerson .
func RoutePerson(e *echo.Group, sotrage Storage) {
	h := newPerson(sotrage)

	person := e.Group("/persons")
	person.Use(middleware.Authentication)

	person.POST("/", h.create)
	person.PUT("/:id", h.update)
	person.DELETE("/:id", h.delete)
	person.GET("/:id", h.getByID)
	person.GET("/", h.getAll)
}

// RouteLogin .
func RouteLogin(e *echo.Group, storage Storage) {
	h := newLogin(storage)

	e.POST("/login", h.login)
}
