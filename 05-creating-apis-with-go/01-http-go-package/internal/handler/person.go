package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"go-api/internal/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "The person doesn't have the correct structure", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "A problem occured when creating the person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Person created successfully!", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "The id should be a integer positive number", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "The person doesn't have the correct structure", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		errMsg := "A problem occured searching for the person with ID " + strconv.Itoa(ID)
		response := newResponse(Error, errMsg, nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Person updated successfully!", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "The id should be a integer positive number", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		errMsg := "A problem occured searching for the person with ID " + strconv.Itoa(ID)
		response := newResponse(Error, errMsg, nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Something went wrong when trying to deleting the person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Person deleted successfully!", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "The id should be a integer positive number", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		errMsg := "A problem occured searching for the person with ID " + strconv.Itoa(ID)
		response := newResponse(Error, errMsg, nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "A problem occured trying to delete the person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Person deleted successfully!", data)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "A problem occured trying to get all registries", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, response)
}
