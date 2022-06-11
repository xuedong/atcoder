package main

import (
	"fmt"
	"sort"
)

func main()  {
	var n, q int
	fmt.Scanf("%v %v", &n, &q)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	sort.Ints(nums)

	prefixes := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		prefixes[i] = prefixes[i-1] + int64(nums[i-1])
	}

	for j := 1; j <= q; j++ {
		var target int
		fmt.Scanf("%v", &target)

		left, right := 0, len(nums) - 1

		for left <= right {
			mid := left + (right - left) / 2

			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		var ans int64
		ans = int64(left) * int64(target)
    	ans -= prefixes[right+1]
    	ans += prefixes[n] - prefixes[left]
    	ans -= int64((n - left)) * int64(target)
		fmt.Printf("%v\n", ans)
	}
}
