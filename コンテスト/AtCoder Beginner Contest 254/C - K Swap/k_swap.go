package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
  	fmt.Scanf("%d %d", &n, &k)

  	nums := make([][]int, k)
	rest := n % k
	for i := 0; i < k; i++ {
		length := n / k
		if rest > 0 {
			length++
			rest--
		}

		nums[i] = make([]int, length)
	}

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i%k][i/k])
	}

	for i := 0; i < k; i++ {
		sort.Ints(nums[i])
	}

	for i := 0; i < n-1; i++ {
		if nums[i%k][i/k] > nums[(i+1)%k][(i+1)/k] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
