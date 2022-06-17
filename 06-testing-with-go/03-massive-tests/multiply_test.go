package main

import "testing"

func TestMultiply(t *testing.T) {
	table := []struct {
		name string
		a int
		b int
		want int
	}{
		{name:"2x1", a: 2, b: 1, want: 2},
		{"2x2", 2, 2, 4},
		{"2x3", 2, 3, 6},
		{"2x4", 2, 4, 8},
		{"2x5", 2, 5, 10},
		{"2x6", 2, 6, 12},
		{"2x7", 2, 7, 14},
		{"2x8", 2, 8, 16},
		{"2x9", 2, 9, 18},
		{"2x10", 2, 10, 20},
	}

	for _, v := range table {
		t.Run(v.name, func(t *testing.T){
			got := multiply(v.a, v.b)
			if got != v.want {
				t.Errorf("Obtained %d, expected %d", got, v.want)
			}
		})
	}
}