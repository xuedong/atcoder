package main

import (
    "bufio"
    "fmt"
    "os"
    "math"
)

func main()  {
    in := bufio.NewReader(os.Stdin)

    var a, b, d float64
    fmt.Fscan(in, &a, &b, &d)

    d = math.Pi * d / 180
    x, y := a * math.Cos(d) - b * math.Sin(d), b * math.Cos(d) + a * math.Sin(d)
    fmt.Printf("%f %f\n", x, y)
}
