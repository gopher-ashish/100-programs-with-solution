package main

import "fmt"

func swap(a, b int) (int, int) {
	b = a + b
	a = b - a
	b = b - a
	return a, b
}

func main() {
	a, b := 10, 20
	fmt.Printf("Before a = %d, b = %d \n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After a = %d, b = %d", a, b)
}
