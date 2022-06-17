package handler

import (
	"errors"
	"go-api/internal/model"
)

type StorageMockError struct{}

func (sme *StorageMockError) Create(person *model.Person) error {
	return errors.New("mock error")
}

func (sme *StorageMockError) Update(ID int, person *model.Person) error {
	return errors.New("mock error")
}

func (sme *StorageMockError) Delete(ID int) error {
	return errors.New("mock error")
}

func (sme *StorageMockError) GetByID(ID int) (model.Person, error) {
	return model.Person{}, errors.New("mock error")
}

func (sme *StorageMockError) GetAll() (model.Persons, error) {
	return model.Persons{}, errors.New("mock error")
}

type StorageMockOK struct{}

func (sme *StorageMockOK) Create(person *model.Person) error {
	return nil
}

func (sme *StorageMockOK) Update(ID int, person *model.Person) error {
	return nil
}

func (sme *StorageMockOK) Delete(ID int) error {
	return nil
}

func (sme *StorageMockOK) GetByID(ID int) (model.Person, error) {
	return model.Person{}, nil
}

func (sme *StorageMockOK) GetAll() (model.Persons, error) {
	return model.Persons{}, nil
}
