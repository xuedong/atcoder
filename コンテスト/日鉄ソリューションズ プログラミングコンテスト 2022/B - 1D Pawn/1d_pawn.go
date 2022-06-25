package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, k, q int
    fmt.Fscan(in, &n, &k, &q)

    initial := make([]int, k)
    for i := 0; i < k; i++ {
        fmt.Fscan(in, &initial[i])
    }

    for i := 0; i < q; i++ {
        var operation int
        fmt.Fscan(in, &operation)

        if initial[operation-1] == n {
            continue
        }

        if operation < k && initial[operation] == initial[operation-1] + 1 {
            continue
        }

        initial[operation-1]++
    }

    for i := 0; i < k; i++ {
        fmt.Printf("%d ", initial[i])
    }
}
