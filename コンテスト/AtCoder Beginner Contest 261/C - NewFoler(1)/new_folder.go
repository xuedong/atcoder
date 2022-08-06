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

    counts := make(map[string]int)
    for i := 1; i <= n; i++ {
        var curr string
        fmt.Fscan(in, &curr)
        _, ok := counts[curr]

        if !ok {
            fmt.Printf("%v\n", curr)
            counts[curr] = 1
        } else {
            count := counts[curr]
            fmt.Printf("%v(%d)\n", curr, count)
            counts[curr]++
        }
    }
}
