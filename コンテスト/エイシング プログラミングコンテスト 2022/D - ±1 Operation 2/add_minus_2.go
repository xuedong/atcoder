package main

import (
	"fmt"
	"sort"
)

func binarySearch(target int64, nums []int64) int64 {
	var left, right int64
	left, right = 0, int64(len(nums)) - 1

	for left <= right {
		mid := left + (right - left) / 2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main()  {
	var n, q int64
	fmt.Scanf("%v %v", &n, &q)

	nums := make([]int64, n)
	prefixes := make([]int64, n)
	var i int64
	for i = 0; i < n; i++ {
		fmt.Scanf("%v", &nums[i])
		if i == 0 {
			prefixes[i] = nums[i]
		} else {
			prefixes[i] = prefixes[i-1] + nums[i]
		}
	}

	for i = 0; i < n; i++ {
		nums[]
	}
	sort.Int64s(nums)

	var j int64
	for j = 1; j <= q; j++ {
		var target int64
		fmt.Scanf("%v", &target)
		id := binarySearch(target, nums)

		var ans int64
		ans = id * target
    	ans -= prefixes[id]
    	ans += prefixes[n] - prefixes[id]
    	ans -= (n - id) * target
		fmt.Printf("%v\n", ans)
	}
}
