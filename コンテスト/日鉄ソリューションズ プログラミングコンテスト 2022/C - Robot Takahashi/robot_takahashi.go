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

    var s string
    fmt.Fscan(in, &s)
    runes := []rune(s)

    weights := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &weights[i])
    }

    best := 0
    max := 0
    for i := 0; i < n; i++ {
        threshold := weights[i]
        if threshold > max {
            max = threshold
        }

        curr := 0
        for j := 0; j < n; j++ {
            if weights[j] >= threshold && runes[j] == '1' {
                curr++
            } else if weights[j] < threshold && runes[j] == '0' {
                curr++
            }
        }

        if curr > best {
            best = curr
        }
    }

    curr := 0
    for j := 0; j < n; j++ {
        if runes[j] == '0' {
            curr++
        }
    }
    if curr > best {
        best = curr
    }

    fmt.Println(best)
}
