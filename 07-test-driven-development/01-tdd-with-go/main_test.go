package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(5, 2)
	want := 7
	if got != want {
		t.Errorf("expected: %d, obtained: %d", want, got)
	}
}

func TestFactorial(t *testing.T) {
	got := Factorial(5)
	want := 120
	if got != want {
		t.Errorf("expected: %d, obtained: %d", want, got)
	}
}