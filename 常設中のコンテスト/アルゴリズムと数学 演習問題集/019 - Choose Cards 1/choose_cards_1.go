package main

import (
    "fmt"
)

func main()  {
    var n int
    fmt.Scanf("%d", &n)

    counts := make([]int, 3)
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &nums[i])
        switch nums[i] {
        case 1:
            counts[0]++
        case 2:
            counts[1]++
        case 3:
            counts[2]++
        default:
            continue
        }
    }

    sum := (counts[0] * (counts[0]-1) + counts[1] * (counts[1]-1) + counts[2] * (counts[2]-1)) / 2
    fmt.Printf("%d\n", sum)
}
