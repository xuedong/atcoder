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

    fmt.Println(powInt(2, n))
}
