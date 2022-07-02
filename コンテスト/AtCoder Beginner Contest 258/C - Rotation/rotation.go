package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, q int
    fmt.Fscan(in, &n, &q)

    var str string
    fmt.Fscan(in, &str)

    len := len(str)

    rotation := 0
    for i := 1; i <= q; i++ {
        var query, num int
        fmt.Fscan(in, &query, &num)

        if query == 1 {
            rotation = (rotation + num) % len
        }

        if query == 2 {
            curr := (num - 1 - rotation) % len
            if curr < 0 {
                curr += len
            }
            fmt.Println(str[curr:curr+1])
        }
    }
}
