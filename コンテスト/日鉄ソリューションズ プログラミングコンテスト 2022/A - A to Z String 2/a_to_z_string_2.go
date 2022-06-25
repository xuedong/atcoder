package main

import (
    "bufio"
    "fmt"
    "os"
    "math"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var n, x float64
    fmt.Fscan(in, &n, &x)

    res := string(64 + int(math.Ceil(x / n)))
    fmt.Println(res)
}
