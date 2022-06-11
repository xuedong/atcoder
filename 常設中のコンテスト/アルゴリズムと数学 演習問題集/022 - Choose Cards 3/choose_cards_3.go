package main

import (
    "fmt"
)

func main()  {
    var n int
    fmt.Scanf("%d", &n)

    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &nums[i])
    }

    counts := make([]int64, 100000)
    for _, num := range nums {
        counts[num]++
    }

    var ans int64
    ans = 0
    for i := 1; i <= 49999; i++ {
        ans += counts[i] * counts[100000 - i]
    }
    ans += counts[50000] * (counts[50000] - 1) / 2

    fmt.Printf("%d\n", ans)
}
