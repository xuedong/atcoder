package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, x int
    fmt.Fscan(in, &n, &x)

    movies := make([]int64, n)
    gameplays := make([]int64, n)
    for i := 0; i <= n-1; i++ {
        fmt.Fscan(in, &movies[i], &gameplays[i])
    }

    ans := movies[0] + int64(x) * gameplays[0]
    prefix := movies[0] + gameplays[0]
    min := gameplays[0]
    for i := 1; i < n; i++ {
        prefix += movies[i] + gameplays[i]
        if gameplays[i] < min {
            min = gameplays[i]
        }
        curr := prefix + int64(x-i-1) * min
        if curr < ans {
            ans = curr
        }
    }

    fmt.Println(ans)
}
