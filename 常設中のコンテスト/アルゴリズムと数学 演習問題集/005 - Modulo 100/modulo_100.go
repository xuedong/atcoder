package main

import (
  	"fmt"
)

func main() {
	var n int
  	fmt.Scanf("%d", &n)

  	nums := make([]int, n)
  	ans := 0
  	for i := 0; i < n; i++ {
      	fmt.Scanf("%d", &nums[i])
      	ans += nums[i]
      	ans %= 100
    }

  	fmt.Printf("%d\n", ans)
}
