package main

import (
	"fmt"
)

func main() {
	numNumber := 11
	denoNumber := 2

	quotient, remainder := numNumber/denoNumber, numNumber%denoNumber

	fmt.Println("quotient result: ", quotient)
	fmt.Println("remainder result: ", remainder)

}
