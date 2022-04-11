package main

import "fmt"

func swap(a, b int) (int, int) {
	var temp int

	temp = a
	a = b
	b = temp

	return a, b
}

func main() {
	a, b := 10, 20
	fmt.Printf("Before a = %d, b = %d \n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After a = %d, b = %d ", a, b)
}
