package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, m, x, t, d int
    fmt.Fscan(in, &n, &m, &x, &t, &d)

    if m >= x {
        fmt.Printf("%d\n", t)
    } else {
        ans := t - d * (x - m)
        fmt.Printf("%d\n", ans)
    }
}
