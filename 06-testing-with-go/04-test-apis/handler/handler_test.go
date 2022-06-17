package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	Get(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code: %d, obtained: %d", http.StatusOK, w.Code)
	}

	// t.Log(w.Body.String())
	got := Person{}
	err := json.NewDecoder(w.Body).Decode(&got)
	if err != nil {
		t.Errorf("can't process the json: %v", err)
	}

	want := Person{
		Name: "Jhoana",
		Age: 31,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected: %v, obtained: %v", want, got)
	}
}