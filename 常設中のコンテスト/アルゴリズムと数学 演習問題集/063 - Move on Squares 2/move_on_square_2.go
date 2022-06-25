package main

import (
  	"bufio"
  	"fmt"
  	"os"
)

func main() {
  	in := bufio.NewReader(os.Stdin)

	var n int
  	fmt.Fscan(in, &n)

  	if n % 2 == 0 {
      	fmt.Println("Yes")
    } else {
      	fmt.Println("No")
    }
}
