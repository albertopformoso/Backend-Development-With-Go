package main

func Add(a, b int) int {
	return a + b
}

func AddAccumulative(numbers ...int) (ans int) {
	for _, v := range numbers {
		ans += v
	}

	return
}
