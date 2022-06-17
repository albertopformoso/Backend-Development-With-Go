package handler

import (
	"bytes"
	"encoding/json"
	"go-api/internal/model"
	"go-api/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestPerson_Create_integration(t *testing.T) {
	t.Cleanup(func() {
		cleanData(t)
	})

	data := []byte(`{"name": "Integration Test", "age": 99}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))

	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	store := storage.NewPsql()
	p := newPerson(&store)
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

type responsePerson struct {
	MessageType string         `json:"message_type"`
	Message     string         `json:"message"`
	Data        []model.Person `json:"data"`
}

func TestPerson_GetAll_integration(t *testing.T) {
	t.Cleanup(func() {
		cleanData(t)
	})
	insertData(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	r.Header.Set("Content-Type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	store := storage.NewPsql()
	p := newPerson(&store)
	err := p.getAll(ctx)
	if err != nil {
		t.Errorf("error not expected, obtained: %v", err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Status code expected: %v, obtained: %v", http.StatusOK, w.Code)
	}

	resp := responsePerson{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("can't unmarshal of the response: %v", err)
	}

	lenPersonsWant := 2
	lenPersonsGot := len(resp.Data)
	if lenPersonsGot != lenPersonsWant {
		t.Errorf("expected %d persons, obtained %d persons", lenPersonsWant, lenPersonsGot)
	}

	wantName := "Albert"
	if resp.Data[0].Name != wantName {
		t.Errorf("expected name: %q, obtained: %q", wantName, resp.Data[0].Name)
	}
}

func cleanData(t *testing.T) {
	store := storage.NewPsql()
	defer store.CloseDB()

	err := store.DeleteAll()
	if err != nil {
		t.Fatalf("can't clean the table: %v", err)
	}
}

func insertData(t *testing.T) {
	store := storage.NewPsql()
	defer store.CloseDB()

	err := store.Create(&model.Person{Name: "Albert", Age: 99})
	if err != nil {
		t.Fatalf("no se pudo registrar la persona %v", err)
	}

	err = store.Create(&model.Person{Name: "Matt", Age: 26})
	if err != nil {
		t.Fatalf("no se pudo registrar la persona %v", err)
	}

}
