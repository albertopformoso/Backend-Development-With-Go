package storage

import (
	"go-api/internal/model"
	"testing"
)

func TestCreate(t *testing.T) {
	data := []struct {
		name      string
		person    *model.Person
		wantedErr error
	}{
		{"Empty person", nil, model.ErrPersonCanNotBeNil},
		{"Non empty 1", &model.Person{Name: "Albert"}, nil},
		{"Non empty 2", &model.Person{Name: "Matt"}, nil},
	}

	m := NewMemory()
	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			gotErr := m.Create(v.person)
			if gotErr != v.wantedErr {
				t.Errorf("expected: %v, obtained: %v", v.wantedErr, gotErr)
			}
		})
	}

	if m.currentID != len(data)-1 {
		t.Errorf("expected %d registries, obtained %d registries", len(data)-1, m.currentID)
	}
}

func TestCreate_empty_person(t *testing.T) {
	m := NewMemory()
	err := m.Create(nil)
	if err == nil {
		t.Errorf("an error was expected, nil was obtained")
	}

	if err != model.ErrPersonCanNotBeNil {
		t.Errorf("expected error: %v, obtained error: %v", model.ErrPersonCanNotBeNil, err)
	}
}

func TestCreate_count_elements(t *testing.T) {
	m := NewMemory()
	p := model.Person{Name: "Albert"}
	err := m.Create(&p)

	if err != nil {
		t.Errorf("Error not expected, obtained: %v", err)
	}

	if m.currentID != 1 {
		t.Errorf("1 Element expected, obtained: %d", m.currentID)
	}
}
