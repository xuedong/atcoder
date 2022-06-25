package main

import (
    "bufio"
    "fmt"
    "os"
    "math"
)

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscan(in, &n)

    moves := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &moves[i])
    }

    sum := 0
    var i int
    for i = n-1; i >= 0; i-- {
        sum += moves[i]
        if sum >= 4 {
            break
        }
    }

    fmt.Println(i+1)
}
