package main

import (
    "bufio"
    "fmt"
    "os"
)

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, l, r int
    fmt.Fscan(in, &n, &l, &r)

    a := make([]int, n)
    for i := range a {
        fmt.Fscan(in, &a[i])
    }

    dpLeft := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dpLeft[i] = min(dpLeft[i-1] + a[i-1], l * i)
    }

    dpRight := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dpRight[i] = min(dpRight[i-1] + a[n-i], r * i)
    }

    ans := l * n
    for i := range a {
        ans = min(ans, dpLeft[i] + dpRight[n-i])
    }

    fmt.Printf("%d\n", ans)
}
