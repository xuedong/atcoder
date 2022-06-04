package main

import (
	"fmt"
)

func main() {
  	var n int
  	fmt.Scanf("%d", &n)

	var res [][]int
	for i := 0; i < n; i++ {
      	tmp := make([]int, i+1)
		tmp[0] = 1
      	fmt.Printf("1 ")

		res = append(res, tmp)
		for j := 1; j <= i; j++ {
			if i == j {
				res[i][j] = 1
              	fmt.Printf("%d ", res[i][j])
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
              	fmt.Printf("%d ", res[i][j])
			}
		}
      	fmt.Println()
	}
}
