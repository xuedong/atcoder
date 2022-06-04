package main

import (
	"fmt"
  	"strconv"
)

func main() {
	var n int
  	fmt.Scanf("%d", &n)

  	s := strconv.Itoa(n)
  	fmt.Println(s[len(s)-2:])
}
