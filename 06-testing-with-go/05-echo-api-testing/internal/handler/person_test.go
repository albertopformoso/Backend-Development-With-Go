package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestPerson_Create_wrong_struct(t *testing.T) {
	data := []byte(`{"name": 123, "age": 18}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))

	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockOK{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("error not expected, obtained %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("Status code expected %d, obtained %d", http.StatusBadRequest, w.Code)
	}

	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("can't unmarshal the body: %v", err)
	}

	wantMessage := "The person doesn't have the correct structure"
	if resp.Message != wantMessage {
		t.Errorf("the response isn't the one expected, obtained: %q, expected: %q", resp.Message, wantMessage)
	}
}

func TestPerson_Create_wrong_storage(t *testing.T) {
	data := []byte(`{"name": "Albert", "age": 18}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))

	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockError{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("error not expected, obtained %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Status code expected %d, obtained %d", http.StatusInternalServerError, w.Code)
	}

	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("can't unmarshal the body: %v", err)
	}

	wantMessage := "A problem occured when creating the person"
	if resp.Message != wantMessage {
		t.Errorf("the response isn't the one expected, obtained: %q, expected: %q", resp.Message, wantMessage)
	}
}

func TestPerson_Create(t *testing.T) {
	data := []byte(`{"name": "Albert", "age": 18}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))

	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockOK{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("error not expected, obtained %v", err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Status code expected %d, obtained %d", http.StatusCreated, w.Code)
	}

	resp := response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("can't unmarshal the body: %v", err)
	}

	wantMessage := "Person created successfully!"
	if resp.Message != wantMessage {
		t.Errorf("the response isn't the one expected, obtained: %q, expected: %q", resp.Message, wantMessage)
	}
}
