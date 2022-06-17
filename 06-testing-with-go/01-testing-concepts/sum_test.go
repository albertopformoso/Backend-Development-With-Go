package main

import "testing"

func TestAdd(t *testing.T) {
	want := 5
	t.Logf("want value: %d \n", want)
	got := Add(2, 3)
	t.Logf("go value: %d \n", got)
	if got != want {
		t.Logf("Error: expected %d, obtained %d", want, got)
		t.Fail()
	}

	t.Log("Add test executed successfully")
}

func TestAddAccumulative(t *testing.T) {
	want := 6
	got := AddAccumulative(1, 2, 3)
	if got != want {
		t.Errorf("Error: expected %d, obtained %d", want, got)
	}
}
