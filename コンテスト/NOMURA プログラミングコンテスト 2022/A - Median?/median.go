package main

import (
	"fmt"
)

func main() {
	var a, b, c int
  	fmt.Scanf("%d %d %d", &a, &b, &c)

  	if b >= a && b <= c {
      	fmt.Println("Yes")
    } else if b <= a && b >= c {
      	fmt.Println("Yes")
    } else {
      	fmt.Println("No")
    }
}
