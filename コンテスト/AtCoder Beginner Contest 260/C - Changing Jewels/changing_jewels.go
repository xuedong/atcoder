package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, x, y int
    fmt.Fscan(in, &n, &x, &y)

    if n == 1 {
        fmt.Println(0)
        return
    }

    red := make([]int, n+1)
    blue := make([]int, n+1)
    blue[1] = 1
    blue[2] = y
    red[2] = x*y

    for i := 3; i <= n; i++ {
        blue[i] = red[i-1] + y * blue[i-1]
        red[i] = red[i-1] + x * blue[i]
    }

    fmt.Println(red[n])
}
