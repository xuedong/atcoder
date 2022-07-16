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

    fmt.Println(2*n)

    count := 0
    for n > 4 {
        n -= 4
        count++
    }

    if n > 0 {
        fmt.Print(n)
    }
    for i := 1; i <= count; i++ {
        fmt.Print(4)
    }
}
