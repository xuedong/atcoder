package main

import (
    "fmt"
)

func gcd(a, b int64) int64  {
    for b > 0 {
        a, b = b, a % b
    }

    return a
}

func main()  {
    var n int
    fmt.Scanf("%d", &n)

    nums := make([]int64, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%v", &nums[i])
    }

    ans := nums[0]
    for i := 1; i < n; i++ {
        ans = gcd(ans, nums[i])
    }

    fmt.Printf("%v\n", ans)
}
