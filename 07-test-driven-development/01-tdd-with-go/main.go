package main

func Add(a, b int) int {
	return b + a
}

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}