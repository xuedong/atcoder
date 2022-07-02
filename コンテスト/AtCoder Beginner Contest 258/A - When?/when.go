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

    if n < 10 {
        fmt.Printf("21:0%d\n", n)
    } else if n < 60 {
        fmt.Printf("21:%d\n", n)
    } else if n < 70 {
        fmt.Printf("22:0%d\n", n-60)
    } else {
        fmt.Printf("22:%d\n", n-60)
    }
}
