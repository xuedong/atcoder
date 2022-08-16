package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	character := string(rune(n))
	fmt.Printf(character)
}
