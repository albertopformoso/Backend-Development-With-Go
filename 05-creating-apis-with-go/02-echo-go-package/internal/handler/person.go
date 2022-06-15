package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go-api/internal/model"

	"github.com/labstack/echo"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {

	data := model.Person{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "The person doesn't have the correct structure", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "A problem occured when creating the person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Person created successfully!", nil)
	return c.JSON(http.StatusCreated, response)
}

func (p *person) update(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	fmt.Printf("Update ID: %v", ID)
	if err != nil {
		response := newResponse(Error, "The id should be a integer positive number", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err = c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "The person doesn't have the correct structure", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		errMsg := "A problem occured searching for the person with ID " + strconv.Itoa(ID)
		response := newResponse(Error, errMsg, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Person updated successfully!", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) delete(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "The id should be a integer positive number", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		errMsg := "A problem occured searching for the person with ID " + strconv.Itoa(ID)
		response := newResponse(Error, errMsg, nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "Something went wrong when trying to deleting the person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Person deleted successfully!", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getByID(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "The id should be a integer positive number", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		errMsg := "A problem occured searching for the person with ID " + strconv.Itoa(ID)
		response := newResponse(Error, errMsg, nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "A problem occured trying to delete the person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Person deleted successfully!", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getAll(c echo.Context) error {

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "A problem occured trying to get all registries", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}
