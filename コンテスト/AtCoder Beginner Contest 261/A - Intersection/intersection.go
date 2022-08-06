package main

import (
    "bufio"
    "fmt"
    "os"
)

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var l1, r1, l2, r2 int
    fmt.Fscan(in, &l1, &r1, &l2, &r2)

    if l2 >= r1 {
        fmt.Println(0)
    } else if l1 >= r2 {
        fmt.Println(0)
    } else {
        ans := min(r1, r2) - max(l1, l2)
        fmt.Printf("%d\n", ans)
    }
}
