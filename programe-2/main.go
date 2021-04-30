package main

import "fmt"

/* their are two ways to find the ASCII value of characters
1. consider character as rune and save in array
2. using printf function
*/

func usingRune(str string) (result []int) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	return
}

func usingPrintFstatement(str string) {
	for _, r := range str {
		fmt.Printf("%d ", r)
	}

}

func main() {
	str := "HELLOW WORLD"
	fmt.Println("------------------- OPTION-1 -----------------------------------")

	array := usingRune(str)
	fmt.Println(array, "\n")

	fmt.Println("------------------- OPTION-1 -----------------------------------")
	usingPrintFstatement(str)
	fmt.Println("")

}
