package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscan(in, &n)

    ancestors := make([]int, n-1)
    for i := range ancestors {
        fmt.Fscan(in, &ancestors[i])
    }

    ans := 1
    curr := ancestors[n-2]
    for curr != 1 {
        curr = ancestors[curr-2]
        ans++
    }

    fmt.Printf("%d\n", ans)
}
