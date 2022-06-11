package main

import (
    "fmt"
)

func main()  {
    var n int
    fmt.Scanf("%d", &n)

    counts := make([]int, 4)
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &nums[i])
        switch nums[i] {
        case 100:
            counts[0]++
        case 200:
            counts[1]++
        case 300:
            counts[2]++
        case 400:
            counts[3]++
        default:
            continue
        }
    }

    fmt.Printf("%d\n", counts[0] * counts[3] + counts[1] * counts[2])
}
