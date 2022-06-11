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

    dp := make([][]int, 6)
    for row := range dp {
        dp[row] = make([]int, 1001)
    }
    dp[0][0] = 1

    for _, num := range nums {
        for i := 5; i >= 1; i-- {
            for j := 1000; j >= num; j-- {
                dp[i][j] += dp[i-1][j-num]
            }
        }
    }

    fmt.Printf("%d\n", dp[5][1000])
}
