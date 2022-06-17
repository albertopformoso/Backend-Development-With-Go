package main

import (
	"testing"
	"reflect"
)

func TestDog(t *testing.T) {
	want := &Dog{
		Name: "Firulais",
		Age: 1,
		Kind: Kind{
			Name: "criollo",
		},
	}
	got := &Dog{
		Name: "Firulais",
		Age: 1,
		Kind: Kind{
			Name: "criollo",
		},
	}

	// t.Logf("want %p, got %p", want, got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, obtained: %v", want, got)
	}
}