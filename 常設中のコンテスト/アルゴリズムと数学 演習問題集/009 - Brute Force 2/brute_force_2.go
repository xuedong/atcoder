package main

import (
    "fmt"
)

func dp(nums []int, s int) bool {
    if len(nums) == 0 {
        return false
    }

    if nums[0] == s {
        return true
    } else if nums[0] > s {
        return dp(nums[1:], s)
    } else {
        return dp(nums[1:], s) || dp(nums[1:], s-nums[0])
    }
}

func main()  {
    var n, s int
    fmt.Scanf("%d %d", &n, &s)

    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &nums[i])
    }

    flag := dp(nums, s)
    if flag {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
