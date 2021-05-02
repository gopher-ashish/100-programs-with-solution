package main

import (
	"fmt"
)

func main() {
	array := []int{98, 23, 65, 34, 12, 1, 3, 2, 5, 7, 9, 8, 6, 4, 0, 10, 11, 13, 15, 61}
	fmt.Println("Before Sorting: ")
	fmt.Println("\t ", array)

	sortedArrya := mergeSortUsingGoRoutine(array)
	fmt.Println("After Sorting: ")
	fmt.Println("\t ", sortedArrya)
}

func mergeSortUsingGoRoutine(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	var left []int

	done := make(chan bool)
	mid := len(array) / 2
	go func() {
		left = mergeSortUsingGoRoutine(array[:mid])
		done <- true
	}()
	right := mergeSortUsingGoRoutine(array[mid:])
	<-done
	return mergeData(left, right)
}

func mergeData(left, right []int) []int {
	sortedArray := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(right) == 0 {
			return append(sortedArray, left...)
		} else if len(left) == 0 {
			return append(sortedArray, right...)
		} else if left[0] < right[0] {
			sortedArray = append(sortedArray, left[0])
			left = left[1:]
		} else {
			sortedArray = append(sortedArray, right[0])
			right = right[1:]
		}
	}

	return sortedArray
}
